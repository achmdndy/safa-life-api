package dto

import (
	"time"
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// TajweedWordResponse represents a single word with its Tajweed rule.
type TajweedWordResponse struct {
	Word  string `json:"word"`
	Rule  string `json:"rule"`
	Color string `json:"color"`
}

// AyahTajweedResponse represents the Tajweed rules for a specific Ayah.
type AyahTajweedResponse struct {
	TajweedID string                `json:"tajweed_id"`
	SurahID   int                   `json:"surah_id"`
	AyahID    int                   `json:"ayah_id"`
	Words     []TajweedWordResponse `json:"words"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

// TajweedRuleResponse represents a specific Tajweed rule with its explanation.
type TajweedRuleResponse struct {
	ID          string    `json:"id"`
	Rule        string    `json:"rule"`
	Explanation string    `json:"explanation"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// --- Conversion functions from domain to response DTOs ---

// ToTajweedWordResponse converts domain TajweedWord to TajweedWordResponse
func ToTajweedWordResponse(word tajweed.TajweedWord) TajweedWordResponse {
	return TajweedWordResponse{
		Word:  word.Word,
		Rule:  word.Rule,
		Color: word.Color,
	}
}

// ToAyahTajweedResponse converts domain AyahTajweed to AyahTajweedResponse
func ToAyahTajweedResponse(ayahTajweed tajweed.AyahTajweed) AyahTajweedResponse {
	words := make([]TajweedWordResponse, len(ayahTajweed.Words))
	for i, w := range ayahTajweed.Words {
		words[i] = ToTajweedWordResponse(w)
	}
	return AyahTajweedResponse{
		TajweedID: ayahTajweed.TajweedID,
		SurahID:   ayahTajweed.SurahID,
		AyahID:    ayahTajweed.AyahID,
		Words:     words,
		CreatedAt: ayahTajweed.CreatedAt,
		UpdatedAt: ayahTajweed.UpdatedAt,
	}
}

// ToTajweedRuleResponse converts domain TajweedRule to TajweedRuleResponse
func ToTajweedRuleResponse(rule tajweed.TajweedRule) TajweedRuleResponse {
	return TajweedRuleResponse{
		ID:          rule.ID,
		Rule:        rule.Rule,
		Explanation: rule.Explanation,
		Color:       rule.Color,
		CreatedAt:   rule.CreatedAt,
		UpdatedAt:   rule.UpdatedAt,
	}
}

// ToTajweedRuleResponseSlice converts slice of domain TajweedRule to slice of TajweedRuleResponse
func ToTajweedRuleResponseSlice(rules []tajweed.TajweedRule) []TajweedRuleResponse {
	result := make([]TajweedRuleResponse, len(rules))
	for i, rule := range rules {
		result[i] = ToTajweedRuleResponse(rule)
	}
	return result
}
