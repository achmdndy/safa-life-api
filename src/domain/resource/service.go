package resource

import (
	"context"
	"fmt"
	"time"
)

type resourceService struct {
	repo ResourceRepository
}

// NewResourceService creates a new instance of ResourceService.
func NewResourceService(repo ResourceRepository) ResourceService {
	return &resourceService{repo: repo}
}

func (s *resourceService) GetAllResources(ctx context.Context) ([]Resource, error) {
	return s.repo.GetAll(ctx)
}

func (s *resourceService) GetResourceByID(ctx context.Context, id string) (*Resource, error) {
	if id == "" {
		return nil, fmt.Errorf("resource ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *resourceService) CreateResource(ctx context.Context, resource *Resource) (*Resource, error) {
	if resource.ID == "" || resource.Name == "" || resource.Type == "" || resource.DownloadURL == "" {
		return nil, fmt.Errorf("resource ID, name, type, and download URL cannot be empty")
	}
	return s.repo.Create(ctx, resource)
}

func (s *resourceService) UpdateResource(ctx context.Context, resource *Resource) (*Resource, error) {
	if resource.ID == "" {
		return nil, fmt.Errorf("resource ID cannot be empty")
	}
	return s.repo.Update(ctx, resource)
}

func (s *resourceService) DeleteResource(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("resource ID cannot be empty")
	}
	return s.repo.Delete(ctx, id)
}

func (s *resourceService) UpdateInstallStatus(ctx context.Context, id string, isInstalled bool) (*Resource, error) {
	resource, err := s.GetResourceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resource.IsInstalled = isInstalled
	if isInstalled {
		now := time.Now()
		resource.InstalledAt = &now
	} else {
		resource.InstalledAt = nil
	}

	return s.repo.Update(ctx, resource)
}