package tajweed

import (
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// toTajweedWordDomain converts TajweedWordModel to domain TajweedWord.
func toTajweedWordDomain(model TajweedWordModel) tajweed.TajweedWord {
	return tajweed.TajweedWord{
		Word:  model.Word,
		Rule:  model.Rule,
		Color: model.Color,
	}
}

// fromTajweedWordDomain converts domain TajweedWord to TajweedWordModel.
func fromTajweedWordDomain(domain tajweed.TajweedWord) TajweedWordModel {
	return TajweedWordModel{
		Word:  domain.Word,
		Rule:  domain.Rule,
		Color: domain.Color,
	}
}

// ToAyahTajweedDomain converts AyahTajweedModel to domain AyahTajweed.
func (m *AyahTajweedModel) ToAyahTajweedDomain() tajweed.AyahTajweed {
	words := make([]tajweed.TajweedWord, len(m.Words))
	for i, w := range m.Words {
		words[i] = toTajweedWordDomain(w)
	}
	return tajweed.AyahTajweed{
		TajweedID: m.TajweedID,
		SurahID:   m.SurahID,
		AyahID:    m.AyahID,
		Words:     words,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromAyahTajweedDomain converts domain AyahTajweed to AyahTajweedModel.
func FromAyahTajweedDomain(d *tajweed.AyahTajweed) *AyahTajweedModel {
	words := make([]TajweedWordModel, len(d.Words))
	for i, w := range d.Words {
		words[i] = fromTajweedWordDomain(w)
	}
	return &AyahTajweedModel{
		TajweedID: d.TajweedID,
		SurahID:   d.SurahID,
		AyahID:    d.AyahID,
		Words:     words,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToTajweedRuleDomain converts TajweedRuleModel to domain TajweedRule.
func (m *TajweedRuleModel) ToTajweedRuleDomain() tajweed.TajweedRule {
	return tajweed.TajweedRule{
		ID:          m.ID,
		Rule:        m.Rule,
		Explanation: m.Explanation,
		Color:       m.Color,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// FromTajweedRuleDomain converts domain TajweedRule to TajweedRuleModel.
func FromTajweedRuleDomain(d *tajweed.TajweedRule) *TajweedRuleModel {
	return &TajweedRuleModel{
		ID:          d.ID,
		Rule:        d.Rule,
		Explanation: d.Explanation,
		Color:       d.Color,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

// ToTajweedRuleDomainSlice converts a slice of TajweedRuleModel to a slice of domain TajweedRule.
func ToTajweedRuleDomainSlice(models []TajweedRuleModel) []tajweed.TajweedRule {
	domains := make([]tajweed.TajweedRule, len(models))
	for i, model := range models {
		domains[i] = model.ToTajweedRuleDomain()
	}
	return domains
}