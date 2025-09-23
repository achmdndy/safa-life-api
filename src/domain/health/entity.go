package health

type HealthStatus struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}
