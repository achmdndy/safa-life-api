package health

type Service struct {
	repo CheckerRepository
}

func NewService(repo CheckerRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetHealth() HealthStatus {
	services := map[string]string{
		"database": "ok",
		"redis":    "ok",
	}

	if err := s.repo.CheckDatabase(); err != nil {
		services["database"] = "down"
	}
	if err := s.repo.CheckRedis(); err != nil {
		services["redis"] = "down"
	}

	status := "ok"
	for _, v := range services {
		if v != "ok" {
			status = "degraded"
			break
		}
	}

	return HealthStatus{
		Status:   status,
		Services: services,
	}
}
