package dto

// Surah Request DTOs
type CreateSurahRequest struct {
	NameArabic      string `json:"nameArabic" binding:"required,min=1,max=255"`
	NameEnglish     string `json:"nameEnglish" binding:"required,min=1,max=255"`
	RevelationPlace string `json:"revelationPlace" binding:"required,min=1,max=100"`
	RevelationOrder int    `json:"revelationOrder" binding:"required,min=1,max=114"`
	AyahCount       int    `json:"ayahCount" binding:"required,min=1"`
	CreatedBy       string `json:"createdBy" binding:"required,min=1,max=255"`
}

type UpdateSurahRequest struct {
	ID              string `uri:"id" binding:"required"`
	NameArabic      string `json:"nameArabic" binding:"required,min=1,max=255"`
	NameEnglish     string `json:"nameEnglish" binding:"required,min=1,max=255"`
	RevelationPlace string `json:"revelationPlace" binding:"required,min=1,max=100"`
	RevelationOrder int    `json:"revelationOrder" binding:"required,min=1,max=114"`
	AyahCount       int    `json:"ayahCount" binding:"required,min=1"`
	UpdatedBy       string `json:"updatedBy" binding:"required,min=1,max=255"`
}

type GetSurahByIdRequest struct {
	ID      string `uri:"id" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"ayahs" enums:"ayahs" doc:"Include related data (ayahs)"`
}

type GetSurahByNumberRequest struct {
	Number  int    `uri:"number" binding:"required,min=1,max=114"`
	Include string `form:"include" binding:"omitempty" example:"ayahs" enums:"ayahs" doc:"Include related data (ayahs)"`
}

type GetAllSurahsRequest struct {
	Limit   int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"ayahs" enums:"ayahs" doc:"Include related data (ayahs)"`
}

type GetSurahsByRevelationPlaceRequest struct {
	Place  string `uri:"place" binding:"required"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int    `form:"offset" binding:"omitempty,min=0"`
}

type DeleteSurahRequest struct {
	ID string `uri:"id" binding:"required"`
}

// Ayah Request DTOs
type CreateAyahRequest struct {
	SurahID      string `json:"surahId" binding:"required"`
	Text         string `json:"text" binding:"required,min=1"`
	PageNumber   int    `json:"pageNumber" binding:"required,min=1,max=604"`
	JuzNumber    int    `json:"juzNumber" binding:"required,min=1,max=30"`
	HizbNumber   int    `json:"hizbNumber" binding:"required,min=1,max=60"`
	ManzilNumber int    `json:"manzilNumber" binding:"required,min=1,max=7"`
	CreatedBy    string `json:"createdBy" binding:"required,min=1,max=255"`
}

type UpdateAyahRequest struct {
	ID           string `uri:"id" binding:"required"`
	Text         string `json:"text" binding:"required,min=1"`
	PageNumber   int    `json:"pageNumber" binding:"required,min=1,max=604"`
	JuzNumber    int    `json:"juzNumber" binding:"required,min=1,max=30"`
	HizbNumber   int    `json:"hizbNumber" binding:"required,min=1,max=60"`
	ManzilNumber int    `json:"manzilNumber" binding:"required,min=1,max=7"`
	UpdatedBy    string `json:"updatedBy" binding:"required,min=1,max=255"`
}

type GetAyahByIdRequest struct {
	ID      string `uri:"id" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"surah" enums:"surah" doc:"Include related data (surah)"`
}

type GetAyahsBySurahIdRequest struct {
	SurahID string `uri:"surahId" binding:"required"`
	Limit   int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"surah" enums:"surah" doc:"Include related data (surah)"`
}

type GetAyahsByJuzNumberRequest struct {
	JuzNumber int `uri:"juzNumber" binding:"required,min=1,max=30"`
	Limit     int `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset    int `form:"offset" binding:"omitempty,min=0"`
}

type GetAyahsByPageNumberRequest struct {
	PageNumber int `uri:"pageNumber" binding:"required,min=1,max=604"`
	Limit      int `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset     int `form:"offset" binding:"omitempty,min=0"`
}

type GetAyahsByHizbNumberRequest struct {
	HizbNumber int `uri:"hizbNumber" binding:"required,min=1,max=60"`
	Limit      int `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset     int `form:"offset" binding:"omitempty,min=0"`
}

type GetAyahsByManzilNumberRequest struct {
	ManzilNumber int `uri:"manzilNumber" binding:"required,min=1,max=7"`
	Limit        int `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset       int `form:"offset" binding:"omitempty,min=0"`
}

type GetAllAyahsRequest struct {
	Limit   int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"surah" enums:"surah" doc:"Include related data (surah)"`
}

type DeleteAyahRequest struct {
	ID string `uri:"id" binding:"required"`
}

// Juz Request DTOs
type CreateJuzRequest struct {
	StartSurahID string `json:"startSurahId" binding:"required"`
	EndSurahID   string `json:"endSurahId" binding:"required"`
	StartAyahID  string `json:"startAyahId" binding:"required"`
	EndAyahID    string `json:"endAyahId" binding:"required"`
	CreatedBy    string `json:"createdBy" binding:"required,min=1,max=255"`
}

type UpdateJuzRequest struct {
	ID           string `uri:"id" binding:"required"`
	StartSurahID string `json:"startSurahId" binding:"required"`
	EndSurahID   string `json:"endSurahId" binding:"required"`
	StartAyahID  string `json:"startAyahId" binding:"required"`
	EndAyahID    string `json:"endAyahId" binding:"required"`
	UpdatedBy    string `json:"updatedBy" binding:"required,min=1,max=255"`
}

type GetJuzByIdRequest struct {
	ID      string `uri:"id" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes start/end surah and ayah)"`
}

type GetJuzByNumberRequest struct {
	Number  int    `uri:"number" binding:"required,min=1,max=30"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes start/end surah and ayah)"`
}

type GetAllJuzRequest struct {
	Limit   int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes start/end surah and ayah)"`
}

type DeleteJuzRequest struct {
	ID string `uri:"id" binding:"required"`
}
