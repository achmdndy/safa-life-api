package topic

import "github.com/achmdndy/safa-life-api/src/domain/topic"

// ToTopicDomain converts TopicModel to domain.Topic
func (m *TopicModel) ToTopicDomain() topic.Topic {
	return topic.Topic{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// FromTopicDomain converts domain.Topic to TopicModel
func FromTopicDomain(d *topic.Topic) *TopicModel {
	return &TopicModel{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

// ToTopicDomainSlice converts a slice of TopicModel to a slice of domain.Topic
func ToTopicDomainSlice(models []TopicModel) []topic.Topic {
	domains := make([]topic.Topic, len(models))
	for i, m := range models {
		domains[i] = m.ToTopicDomain()
	}
	return domains
}

// ToTopicAyahDomain converts TopicAyahModel to domain.TopicAyah
func (m *TopicAyahModel) ToTopicAyahDomain() topic.TopicAyah {
	return topic.TopicAyah{
		TopicID:   m.TopicID,
		SurahID:   m.SurahID,
		AyahID:    m.AyahID,
		CreatedAt: m.CreatedAt,
	}
}

// FromTopicAyahDomain converts domain.TopicAyah to TopicAyahModel
func FromTopicAyahDomain(d *topic.TopicAyah) *TopicAyahModel {
	return &TopicAyahModel{
		TopicID:   d.TopicID,
		SurahID:   d.SurahID,
		AyahID:    d.AyahID,
		CreatedAt: d.CreatedAt,
	}
}

// ToTopicAyahDomainSlice converts a slice of TopicAyahModel to a slice of domain.TopicAyah
func ToTopicAyahDomainSlice(models []TopicAyahModel) []topic.TopicAyah {
	domains := make([]topic.TopicAyah, len(models))
	for i, m := range models {
		domains[i] = m.ToTopicAyahDomain()
	}
	return domains
}