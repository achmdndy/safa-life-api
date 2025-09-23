package dto

// HealthResponse represents the health status response for presentation layer
type HealthResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}