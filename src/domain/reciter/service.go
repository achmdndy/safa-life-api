package reciter

import (
	"context"
	"fmt"
)

type reciterService struct {
	repo ReciterRepository
}

// NewReciterService creates a new instance of ReciterService.
func NewReciterService(repo ReciterRepository) ReciterService {
	return &reciterService{repo: repo}
}

func (s *reciterService) GetAllReciters(ctx context.Context) ([]Reciter, error) {
	return s.repo.GetAll(ctx)
}

func (s *reciterService) GetReciterByID(ctx context.Context, id string) (*Reciter, error) {
	if id == "" {
		return nil, fmt.Errorf("reciter ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *reciterService) CreateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error) {
	if reciter.ID == "" || reciter.Name == "" {
		return nil, fmt.Errorf("reciter ID and name cannot be empty")
	}
	return s.repo.Create(ctx, reciter)
}

func (s *reciterService) UpdateReciter(ctx context.Context, reciter *Reciter) (*Reciter, error) {
	if reciter.ID == "" {
		return nil, fmt.Errorf("reciter ID cannot be empty")
	}
	return s.repo.Update(ctx, reciter)
}

func (s *reciterService) DeleteReciter(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("reciter ID cannot be empty")
	}
	return s.repo.Delete(ctx, id)
}
