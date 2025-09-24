package story

import "github.com/achmdndy/safa-life-api/src/domain/story"

// ToStoryDomain converts StoryModel to domain.Story
func (m *StoryModel) ToStoryDomain() story.Story {
	return story.Story{
		ID:          m.ID,
		Title:       m.Title,
		Summary:     m.Summary,
		Characters:  m.Characters,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// FromStoryDomain converts domain.Story to StoryModel
func FromStoryDomain(d *story.Story) *StoryModel {
	return &StoryModel{
		ID:          d.ID,
		Title:       d.Title,
		Summary:     d.Summary,
		Characters:  d.Characters,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

// ToStoryDomainSlice converts a slice of StoryModel to a slice of domain.Story
func ToStoryDomainSlice(models []StoryModel) []story.Story {
	domains := make([]story.Story, len(models))
	for i, m := range models {
		domains[i] = m.ToStoryDomain()
	}
	return domains
}

// ToStoryAyahDomain converts StoryAyahModel to domain.StoryAyah
func (m *StoryAyahModel) ToStoryAyahDomain() story.StoryAyah {
	return story.StoryAyah{
		StoryID:   m.StoryID,
		SurahID:   m.SurahID,
		AyahID:    m.AyahID,
		CreatedAt: m.CreatedAt,
	}
}

// FromStoryAyahDomain converts domain.StoryAyah to StoryAyahModel
func FromStoryAyahDomain(d *story.StoryAyah) *StoryAyahModel {
	return &StoryAyahModel{
		StoryID:   d.StoryID,
		SurahID:   d.SurahID,
		AyahID:    d.AyahID,
		CreatedAt: d.CreatedAt,
	}
}

// ToStoryAyahDomainSlice converts a slice of StoryAyahModel to a slice of domain.StoryAyah
func ToStoryAyahDomainSlice(models []StoryAyahModel) []story.StoryAyah {
	domains := make([]story.StoryAyah, len(models))
	for i, m := range models {
		domains[i] = m.ToStoryAyahDomain()
	}
	return domains
}