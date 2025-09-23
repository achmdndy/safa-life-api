package middlewares

import (
	"github.com/gin-gonic/gin"
)

// LoggingMiddleware returns a logging middleware
func LoggingMiddleware() gin.HandlerFunc {
	return gin.Logger()
}

// RecoveryMiddleware returns a recovery middleware
func RecoveryMiddleware() gin.HandlerFunc {
	return gin.Recovery()
}