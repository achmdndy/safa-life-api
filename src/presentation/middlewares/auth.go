package middlewares

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	appcore "github.com/safalife/core-api/cli/core"
	"github.com/safalife/core-api/src/presentation/core"
)

// simple JWKS cache to avoid fetching keys on every request
var jwksCache = struct {
	mu        sync.RWMutex
	keys      map[string]interface{}
	fetchedAt time.Time
}{
	keys: make(map[string]interface{}),
}

type jwk struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Alg string   `json:"alg"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type jwks struct {
	Keys []jwk `json:"keys"`
}

func base64URLToBigInt(s string) (*big.Int, error) {
	// base64url without padding
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetBytes(b)
	return i, nil
}

func rsaPublicKeyFromJWK(nStr, eStr string) (*rsa.PublicKey, error) {
	n, err := base64URLToBigInt(nStr)
	if err != nil {
		return nil, fmt.Errorf("invalid RSA modulus: %w", err)
	}
	eBI, err := base64URLToBigInt(eStr)
	if err != nil {
		return nil, fmt.Errorf("invalid RSA exponent: %w", err)
	}
	if !eBI.IsInt64() {
		return nil, errors.New("RSA exponent too large")
	}
	e := int(eBI.Int64())
	return &rsa.PublicKey{N: n, E: e}, nil
}

func fetchJWKS(client *http.Client, url string) (map[string]interface{}, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// best effort read error body
		var m map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&m)
		return nil, fmt.Errorf("jwks fetch status %d", resp.StatusCode)
	}
	var set jwks
	if err := json.NewDecoder(resp.Body).Decode(&set); err != nil {
		return nil, fmt.Errorf("failed to parse jwks: %w", err)
	}
	out := make(map[string]interface{}, len(set.Keys))
	for _, k := range set.Keys {
		// only support RSA keys for now
		if strings.EqualFold(k.Kty, "RSA") && k.N != "" && k.E != "" {
			pk, err := rsaPublicKeyFromJWK(k.N, k.E)
			if err == nil {
				out[k.Kid] = pk
			}
		}
	}
	return out, nil
}

func getKeyForKid(client *http.Client, url, kid string) (interface{}, error) {
	// try cache first
	jwksCache.mu.RLock()
	key, ok := jwksCache.keys[kid]
	fresh := time.Since(jwksCache.fetchedAt) < 5*time.Minute
	jwksCache.mu.RUnlock()
	if ok && fresh {
		return key, nil
	}
	// refresh
	keys, err := fetchJWKS(client, url)
	if err != nil {
		return nil, err
	}
	jwksCache.mu.Lock()
	jwksCache.keys = keys
	jwksCache.fetchedAt = time.Now()
	jwksCache.mu.Unlock()
	jwksCache.mu.RLock()
	key, ok = jwksCache.keys[kid]
	jwksCache.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("kid %s not found in jwks", kid)
	}
	return key, nil
}

// AuthMiddleware verifies Authorization bearer token via JWKS endpoint.
// It fetches JWKS from cfg.BaseURL + cfg.VerifyPath and validates JWT signature.
// On success, it injects user info (user_id, user_email, user_role, session_id) into Gin context.
func AuthMiddleware(cfg appcore.AuthConfig) gin.HandlerFunc {
	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second
	if timeout == 0 {
		timeout = 3 * time.Second
	}
	client := &http.Client{Timeout: timeout}
	jwksURL := cfg.BaseURL + cfg.VerifyPath

	return func(c *gin.Context) {
		start := time.Now()
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			detail := &core.ErrorDetail{Reason: "Missing or invalid Authorization header"}
			core.Error(c, http.StatusUnauthorized, "Unauthorized", detail, start)
			c.Abort()
			return
		}
		tokenStr := strings.TrimSpace(authHeader[len("Bearer "):])

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			// only accept RSA methods
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			kid, _ := t.Header["kid"].(string)
			if kid == "" {
				return nil, errors.New("missing kid in token header")
			}
			return getKeyForKid(client, jwksURL, kid)
		})
		if err != nil || token == nil || !token.Valid {
			reason := "invalid token"
			if err != nil {
				reason = err.Error()
			}
			detail := &core.ErrorDetail{Reason: reason}
			core.Error(c, http.StatusUnauthorized, "Unauthorized", detail, start)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			detail := &core.ErrorDetail{Reason: "invalid claims"}
			core.Error(c, http.StatusUnauthorized, "Unauthorized", detail, start)
			c.Abort()
			return
		}

		// basic expiration check
		if expRaw, ok := claims["exp"]; ok {
			var exp int64
			switch v := expRaw.(type) {
			case float64:
				exp = int64(v)
			case json.Number:
				n, _ := v.Int64()
				exp = n
			}
			if exp > 0 && time.Unix(exp, 0).Before(time.Now().Add(-30*time.Second)) {
				detail := &core.ErrorDetail{Reason: "token expired"}
				core.Error(c, http.StatusUnauthorized, "Unauthorized", detail, start)
				c.Abort()
				return
			}
		}

		// extract useful fields
		var sub, email, role, sessionID string
		if v, ok := claims["sub"].(string); ok {
			sub = v
		}
		if v, ok := claims["email"].(string); ok {
			email = v
		}
		if v, ok := claims["role"].(string); ok {
			role = v
		}
		// support both sid and sessionId
		if v, ok := claims["sid"].(string); ok {
			sessionID = v
		}
		if sessionID == "" {
			if v, ok := claims["sessionId"].(string); ok {
				sessionID = v
			}
		}

		if sub == "" {
			detail := &core.ErrorDetail{Reason: "missing sub in token"}
			core.Error(c, http.StatusUnauthorized, "Unauthorized", detail, start)
			c.Abort()
			return
		}

		// inject into context
		c.Set("user_id", sub)
		if email != "" {
			c.Set("user_email", email)
		}
		if role != "" {
			c.Set("user_role", role)
		}
		if sessionID != "" {
			c.Set("session_id", sessionID)
		}

		c.Next()
	}
}

// GetUserID returns the authenticated user's ID from context
func GetUserID(c *gin.Context) string {
	if v, ok := c.Get("user_id"); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// GetUserEmail returns the authenticated user's email from context
func GetUserEmail(c *gin.Context) string {
	if v, ok := c.Get("user_email"); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// GetUserRole returns the authenticated user's role from context
func GetUserRole(c *gin.Context) string {
	if v, ok := c.Get("user_role"); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// GetSessionID returns the current session id from context
func GetSessionID(c *gin.Context) string {
	if v, ok := c.Get("session_id"); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
