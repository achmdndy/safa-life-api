package dto

// GetTranslationByIDRequest represents the request to get a translation edition by ID.
type GetTranslationByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// CreateTranslationRequest represents the request to create a new translation edition.
type CreateTranslationRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Language string `json:"language" binding:"required"`
}

// UpdateTranslationRequest represents the request to update a translation edition.
type UpdateTranslationRequest struct {
	ID       string `uri:"id" binding:"required"`
	Name     string `json:"name" binding:"omitempty"`
	Author   string `json:"author" binding:"omitempty"`
	Language string `json:"language" binding:"omitempty"`
}

// GetAyahTranslationRequest represents the request to get a specific ayah translation.
type GetAyahTranslationRequest struct {
	TranslationID string `uri:"translation_id" binding:"required"`
	SurahID       int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID        int    `uri:"ayah_id" binding:"required,min=1"`
}

// GetTranslationsForAyahRequest represents the request to get all translations for a specific ayah.
type GetTranslationsForAyahRequest struct {
	SurahID int `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID  int `uri:"ayah_id" binding:"required,min=1"`
}

// GetTranslationsForSurahRequest represents the request to get all translations for a specific surah.
type GetTranslationsForSurahRequest struct {
	TranslationID string `uri:"translation_id" binding:"required"`
	SurahID       int    `uri:"surah_id" binding:"required,min=1,max=114"`
}

// CreateAyahTranslationRequest represents the request to create a new ayah translation.
type CreateAyahTranslationRequest struct {
	TranslationID string `json:"translation_id" binding:"required"`
	SurahID       int    `json:"surah_id" binding:"required,min=1,max=114"`
	AyahID        int    `json:"ayah_id" binding:"required,min=1"`
	Text          string `json:"text" binding:"required"`
}

// UpdateAyahTranslationRequest represents the request to update an ayah translation.
type UpdateAyahTranslationRequest struct {
	TranslationID string `uri:"translation_id" binding:"required"`
	SurahID       int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID        int    `uri:"ayah_id" binding:"required,min=1"`
	Text          string `json:"text" binding:"omitempty"`
}
