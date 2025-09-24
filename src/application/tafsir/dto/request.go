package dto

// GetTafsirByIDRequest represents the request to get a tafsir edition by ID.
type GetTafsirByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// CreateTafsirRequest represents the request to create a new tafsir edition.
type CreateTafsirRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Language string `json:"language" binding:"required"`
}

// UpdateTafsirRequest represents the request to update a tafsir edition.
type UpdateTafsirRequest struct {
	ID       string `uri:"id" binding:"required"`
	Name     string `json:"name" binding:"omitempty"`
	Author   string `json:"author" binding:"omitempty"`
	Language string `json:"language" binding:"omitempty"`
}

// GetAyahTafsirRequest represents the request to get a specific ayah tafsir.
type GetAyahTafsirRequest struct {
	TafsirID string `uri:"tafsir_id" binding:"required"`
	SurahID  int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID   int    `uri:"ayah_id" binding:"required,min=1"`
}

// GetTafsirsForAyahRequest represents the request to get all tafsirs for a specific ayah.
type GetTafsirsForAyahRequest struct {
	SurahID int `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID  int `uri:"ayah_id" binding:"required,min=1"`
}

// GetTafsirsForSurahRequest represents the request to get all tafsirs for a specific surah.
type GetTafsirsForSurahRequest struct {
	TafsirID string `uri:"tafsir_id" binding:"required"`
	SurahID  int    `uri:"surah_id" binding:"required,min=1,max=114"`
}

// CreateAyahTafsirRequest represents the request to create a new ayah tafsir.
type CreateAyahTafsirRequest struct {
	TafsirID string `json:"tafsir_id" binding:"required"`
	SurahID  int    `json:"surah_id" binding:"required,min=1,max=114"`
	AyahID   int    `json:"ayah_id" binding:"required,min=1"`
	Text     string `json:"text" binding:"required"`
}

// UpdateAyahTafsirRequest represents the request to update an ayah tafsir.
type UpdateAyahTafsirRequest struct {
	TafsirID string `uri:"tafsir_id" binding:"required"`
	SurahID  int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID   int    `uri:"ayah_id" binding:"required,min=1"`
	Text     string `json:"text" binding:"omitempty"`
}
