package dto

// TajweedWordRequest represents a word and its tajweed rule in a request.
type TajweedWordRequest struct {
	Word  string `json:"word" binding:"required"`
	Rule  string `json:"rule" binding:"required"`
	Color string `json:"color" binding:"required"`
}

// GetAyahTajweedRequest represents the request to get tajweed for a specific ayah.
type GetAyahTajweedRequest struct {
	TajweedID string `uri:"tajweed_id" binding:"required"`
	SurahID   int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID    int    `uri:"ayah_id" binding:"required,min=1"`
}

// CreateAyahTajweedRequest represents the request to create tajweed for an ayah.
type CreateAyahTajweedRequest struct {
	TajweedID string               `json:"tajweed_id" binding:"required"`
	SurahID   int                  `json:"surah_id" binding:"required,min=1,max=114"`
	AyahID    int                  `json:"ayah_id" binding:"required,min=1"`
	Words     []TajweedWordRequest `json:"words" binding:"required,min=1,dive"`
}

// UpdateAyahTajweedRequest represents the request to update tajweed for an ayah.
type UpdateAyahTajweedRequest struct {
	TajweedID string               `uri:"tajweed_id" binding:"required"`
	SurahID   int                  `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID    int                  `uri:"ayah_id" binding:"required,min=1"`
	Words     []TajweedWordRequest `json:"words" binding:"omitempty,min=1,dive"`
}

// GetTajweedRuleByIDRequest represents the request to get a specific tajweed rule by ID.
type GetTajweedRuleByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// GetTajweedRuleRequest represents the request to get a specific tajweed rule.
type GetTajweedRuleRequest struct {
	RuleName string `uri:"rule_name" binding:"required"`
}

// CreateTajweedRuleRequest represents the request to create a new tajweed rule.
type CreateTajweedRuleRequest struct {
	ID          string `json:"id" binding:"required"`
	Rule        string `json:"rule" binding:"required"`
	Explanation string `json:"explanation" binding:"required"`
	Color       string `json:"color" binding:"required"`
}

// UpdateTajweedRuleRequest represents the request to update a tajweed rule.
type UpdateTajweedRuleRequest struct {
	ID          string `uri:"id" binding:"required"`
	Rule        string `json:"rule" binding:"omitempty"`
	Explanation string `json:"explanation" binding:"omitempty"`
	Color       string `json:"color" binding:"omitempty"`
}
