package health

// HealthServiceInterface defines the contract for health service operations
type HealthServiceInterface interface {
	GetHealth() HealthStatus
}