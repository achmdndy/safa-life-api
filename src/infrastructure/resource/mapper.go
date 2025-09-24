package resource

import "github.com/achmdndy/safa-life-api/src/domain/resource"

// ToResourceDomain converts ResourceModel to domain.Resource
func (m *ResourceModel) ToResourceDomain() resource.Resource {
	return resource.Resource{
		ID:           m.ID,
		Type:         m.Type,
		Name:         m.Name,
		Language:     m.Language,
		Version:      m.Version,
		LastUpdatedAt: m.LastUpdatedAt,
		DownloadURL:  m.DownloadURL,
		Size:         m.Size,
		IsInstalled:  m.IsInstalled,
		InstalledAt:  m.InstalledAt,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

// FromResourceDomain converts domain.Resource to ResourceModel
func FromResourceDomain(d *resource.Resource) *ResourceModel {
	return &ResourceModel{
		ID:           d.ID,
		Type:         d.Type,
		Name:         d.Name,
		Language:     d.Language,
		Version:      d.Version,
		LastUpdatedAt: d.LastUpdatedAt,
		DownloadURL:  d.DownloadURL,
		Size:         d.Size,
		IsInstalled:  d.IsInstalled,
		InstalledAt:  d.InstalledAt,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

// ToResourceDomainSlice converts a slice of ResourceModel to a slice of domain.Resource
func ToResourceDomainSlice(models []ResourceModel) []resource.Resource {
	domains := make([]resource.Resource, len(models))
	for i, m := range models {
		domains[i] = m.ToResourceDomain()
	}
	return domains
}