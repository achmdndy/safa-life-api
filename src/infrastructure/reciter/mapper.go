package reciter

import "github.com/achmdndy/safa-life-api/src/domain/reciter"

// ToReciterDomain converts ReciterModel to domain.Reciter
func (m *ReciterModel) ToReciterDomain() reciter.Reciter {
	return reciter.Reciter{
		ID:        m.ID,
		Name:      m.Name,
		Style:     m.Style,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromReciterDomain converts domain.Reciter to ReciterModel
func FromReciterDomain(d *reciter.Reciter) *ReciterModel {
	return &ReciterModel{
		ID:        d.ID,
		Name:      d.Name,
		Style:     d.Style,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToReciterDomainSlice converts a slice of ReciterModel to a slice of domain.Reciter
func ToReciterDomainSlice(models []ReciterModel) []reciter.Reciter {
	domains := make([]reciter.Reciter, len(models))
	for i, m := range models {
		domains[i] = m.ToReciterDomain()
	}
	return domains
}