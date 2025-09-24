package dto

// GetSurahRequest represents the request to get a specific surah
type GetSurahRequest struct {
	ID int `uri:"id" binding:"required,min=1,max=114"`
}

// GetAyahRequest represents the request to get a specific ayah
type GetAyahRequest struct {
	SurahID int `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID  int `uri:"ayah_id" binding:"required,min=1"`
}

// GetAyahsByPageRequest represents the request to get ayahs by page
type GetAyahsByPageRequest struct {
	PageNumber int `uri:"page" binding:"required,min=1,max=604"`
}

// GetAyahsByJuzRequest represents the request to get ayahs by juz
type GetAyahsByJuzRequest struct {
	JuzNumber int `uri:"juz" binding:"required,min=1,max=30"`
}

// GetJuzRequest represents the request to get a specific juz
type GetJuzRequest struct {
	ID int `uri:"id" binding:"required,min=1,max=30"`
}

// SearchRequest represents the search request
type SearchRequest struct {
	Query string `form:"q" binding:"required,min=1,max=255"`
	Type  string `form:"type" binding:"omitempty,oneof=surah ayah"`
}

// PaginationRequest represents pagination parameters
type PaginationRequest struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

// CreateSurahRequest represents the request to create a new surah
type CreateSurahRequest struct {
	NameArabic      string `json:"name_ar" binding:"required,max=255"`
	NameEnglish     string `json:"name_en" binding:"required,max=255"`
	RevelationPlace string `json:"revelation_place" binding:"required,oneof=Mecca Medina"`
	RevelationOrder int    `json:"revelation_order" binding:"required,min=1,max=114"`
	AyahCount       int    `json:"ayah_count" binding:"required,min=1"`
}

// UpdateSurahRequest represents the request to update a surah
type UpdateSurahRequest struct {
	ID              int    `uri:"id" binding:"required,min=1,max=114"`
	NameArabic      string `json:"name_ar" binding:"omitempty,max=255"`
	NameEnglish     string `json:"name_en" binding:"omitempty,max=255"`
	RevelationPlace string `json:"revelation_place" binding:"omitempty,oneof=Mecca Medina"`
	RevelationOrder int    `json:"revelation_order" binding:"omitempty,min=1,max=114"`
	AyahCount       int    `json:"ayah_count" binding:"omitempty,min=1"`
}

// CreateAyahRequest represents the request to create a new ayah
type CreateAyahRequest struct {
	SurahID      int    `json:"surah_id" binding:"required,min=1,max=114"`
	AyahID       int    `json:"ayah_id" binding:"required,min=1"`
	Text         string `json:"text" binding:"required"`
	PageNumber   int    `json:"page_number" binding:"required,min=1,max=604"`
	JuzNumber    int    `json:"juz_number" binding:"required,min=1,max=30"`
	HizbNumber   int    `json:"hizb_number" binding:"required,min=1,max=60"`
	ManzilNumber int    `json:"manzil_number" binding:"required,min=1,max=7"`
}

// UpdateAyahRequest represents the request to update an ayah
type UpdateAyahRequest struct {
	SurahID      int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID       int    `uri:"ayah_id" binding:"required,min=1"`
	Text         string `json:"text" binding:"omitempty"`
	PageNumber   int    `json:"page_number" binding:"omitempty,min=1,max=604"`
	JuzNumber    int    `json:"juz_number" binding:"omitempty,min=1,max=30"`
	HizbNumber   int    `json:"hizb_number" binding:"omitempty,min=1,max=60"`
	ManzilNumber int    `json:"manzil_number" binding:"omitempty,min=1,max=7"`
}

// CreateJuzRequest represents the request to create a new juz
type CreateJuzRequest struct {
	StartSurah int `json:"start_surah" binding:"required,min=1,max=114"`
	StartAyah  int `json:"start_ayah" binding:"required,min=1"`
	EndSurah   int `json:"end_surah" binding:"required,min=1,max=114"`
	EndAyah    int `json:"end_ayah" binding:"required,min=1"`
}

// UpdateJuzRequest represents the request to update a juz
type UpdateJuzRequest struct {
	ID         int `uri:"id" binding:"required,min=1,max=30"`
	StartSurah int `json:"start_surah" binding:"omitempty,min=1,max=114"`
	StartAyah  int `json:"start_ayah" binding:"omitempty,min=1"`
	EndSurah   int `json:"end_surah" binding:"omitempty,min=1,max=114"`
	EndAyah    int `json:"end_ayah" binding:"omitempty,min=1"`
}