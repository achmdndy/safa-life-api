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
	Limit   int    `form:"limit" binding:"omitempty,min=0,max=100"`
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
	SurahID   string `uri:"surahId" binding:"required"`
	Limit     int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset    int    `form:"offset" binding:"omitempty,min=0"`
	Include   string `form:"include" binding:"omitempty" example:"surah" enums:"surah" doc:"Include related data (surah)"`
	EditionID string `form:"editionId" binding:"omitempty" doc:"Translation edition ID to include per-ayah translation"`
	ReciterID string `form:"reciterId" binding:"omitempty" doc:"Reciter ID to include per-ayah audio"`
}

type GetAyahsByJuzNumberRequest struct {
	JuzNumber int `uri:"juzNumber" binding:"required,min=1,max=30"`
	Limit     int `form:"limit" binding:"omitempty,min=0,max=100"`
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
	Limit   int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes start/end surah and ayah)"`
}

type DeleteJuzRequest struct {
	ID string `uri:"id" binding:"required"`
}

// Translation Edition queries
type GetTranslationEditionsRequest struct {
	Language string `form:"language" binding:"omitempty"`
	Limit    int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset   int    `form:"offset" binding:"omitempty,min=0"`
}

// Ayah Translation queries
type GetAyahTranslationByAyahAndEditionRequest struct {
	AyahID    string `uri:"ayahId" binding:"required"`
	EditionID string `uri:"editionId" binding:"required"`
}

type GetAyahTranslationsBySurahAndEditionRequest struct {
	SurahID   string `uri:"surahId" binding:"required"`
	EditionID string `uri:"editionId" binding:"required"`
	Limit     int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset    int    `form:"offset" binding:"omitempty,min=0"`
}

// Reciter Request DTOs
type CreateReciterRequest struct {
	Name      string `json:"name" binding:"required,min=1,max=255"`
	Style     string `json:"style" binding:"required,min=1,max=255"`
	Place     string `json:"place" binding:"omitempty,max=20"`
	Picture   string `json:"picture" binding:"omitempty,max=255"`
	CreatedBy string `json:"createdBy" binding:"required,min=1,max=255"`
}

type UpdateReciterRequest struct {
	ID        string `uri:"id" binding:"required"`
	Name      string `json:"name" binding:"required,min=1,max=255"`
	Style     string `json:"style" binding:"required,min=1,max=255"`
	Place     string `json:"place" binding:"omitempty,max=20"`
	Picture   string `json:"picture" binding:"omitempty,max=255"`
	UpdatedBy string `json:"updatedBy" binding:"required,min=1,max=255"`
}

type GetReciterByIdRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetReciterByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

type GetAllRecitersRequest struct {
	Limit  int `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset int `form:"offset" binding:"omitempty,min=0"`
}

type DeleteReciterRequest struct {
	ID string `uri:"id" binding:"required"`
}

// AyahAudioFile Request DTOs
type CreateAyahAudioFileRequest struct {
	ReciterID string  `json:"reciterId" binding:"required"`
	SurahID   string  `json:"surahId" binding:"required"`
	AyahID    string  `json:"ayahId" binding:"required"`
	FilePath  string  `json:"filePath" binding:"required,min=1"`
	Duration  float64 `json:"duration" binding:"required,min=0"`
	ByteSize  float64 `json:"byteSize" binding:"required,min=0"`
	CreatedBy string  `json:"createdBy" binding:"required,min=1,max=255"`
}

type UpdateAyahAudioFileRequest struct {
	ID        string  `uri:"id" binding:"required"`
	FilePath  string  `json:"filePath" binding:"required,min=1"`
	Duration  float64 `json:"duration" binding:"required,min=0"`
	ByteSize  float64 `json:"byteSize" binding:"required,min=0"`
	UpdatedBy string  `json:"updatedBy" binding:"required,min=1,max=255"`
}

type GetAyahAudioFileByIdRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetAyahAudioFileByAyahAndReciterRequest struct {
	AyahID    string `uri:"ayahId" binding:"required"`
	ReciterID string `uri:"reciterId" binding:"required"`
}

type GetAyahAudioFilesBySurahAndReciterRequest struct {
	SurahID   string `uri:"surahId" binding:"required"`
	ReciterID string `uri:"reciterId" binding:"required"`
	Limit     int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset    int    `form:"offset" binding:"omitempty,min=0"`
}

type DeleteAyahAudioFileRequest struct {
	ID string `uri:"id" binding:"required"`
}

// BookmarkAyah Request DTOs
type CreateBookmarkAyahRequest struct {
	UserID    string `json:"userId" binding:"required"`
	AyahID    string `json:"ayahId" binding:"required"`
	CreatedBy string `json:"createdBy" binding:"required,min=1,max=255"`
}

type GetBookmarkAyahByIdRequest struct {
	ID      string `uri:"id" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"ayah" enums:"ayah" doc:"Include related data (ayah)"`
}

type GetBookmarkAyahByUserAndAyahRequest struct {
	UserID  string `uri:"userId" binding:"required"`
	AyahID  string `uri:"ayahId" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"ayah" enums:"ayah" doc:"Include related data (ayah)"`
}

type GetBookmarkAyahsByUserRequest struct {
	UserID  string `uri:"userId" binding:"required"`
	Limit   int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"ayah" enums:"ayah" doc:"Include related data (ayah)"`
}

type DeleteBookmarkAyahRequest struct {
	ID string `uri:"id" binding:"required"`
}

// LastRead Request DTOs
type CreateLastReadRequest struct {
	UserID      string  `json:"userId" binding:"required"`
	SurahID     string  `json:"surahId" binding:"required"`
	AyahID      string  `json:"ayahId" binding:"required"`
	AyahNumber  int     `json:"ayahNumber" binding:"required,min=1"`
	ProgressPct float64 `json:"progressPct" binding:"required,min=0,max=100"`
	CreatedBy   string  `json:"createdBy" binding:"required,min=1,max=255"`
}

type UpdateLastReadRequest struct {
	ID          string  `uri:"id" binding:"required"`
	AyahID      string  `json:"ayahId" binding:"required"`
	AyahNumber  int     `json:"ayahNumber" binding:"required,min=1"`
	ProgressPct float64 `json:"progressPct" binding:"required,min=0,max=100"`
	UpdatedBy   string  `json:"updatedBy" binding:"required,min=1,max=255"`
}

type GetLastReadByIdRequest struct {
	ID      string `uri:"id" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes surah and ayah)"`
}

type GetLastReadsByUserRequest struct {
	UserID  string `uri:"userId" binding:"required"`
	Limit   int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes surah and ayah)"`
}

type GetLastReadByUserAndSurahRequest struct {
	UserID  string `uri:"userId" binding:"required"`
	SurahID string `uri:"surahId" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes surah and ayah)"`
}

type DeleteLastReadRequest struct {
	ID string `uri:"id" binding:"required"`
}

// ProgressHatam Request DTOs
type CreateProgressHatamRequest struct {
	UserID      string  `json:"userId" binding:"required"`
	JuzID       string  `json:"juzId" binding:"required"`
	StartAyahID string  `json:"startAyahId" binding:"required"`
	ProgressPct float64 `json:"progressPct" binding:"required,min=0,max=100"`
	CreatedBy   string  `json:"createdBy" binding:"required,min=1,max=255"`
}

type UpdateProgressHatamRequest struct {
	ID          string  `uri:"id" binding:"required"`
	LastAyahID  string  `json:"lastAyahId" binding:"required"`
	ProgressPct float64 `json:"progressPct" binding:"required,min=0,max=100"`
	IsCompleted bool    `json:"isCompleted" binding:"required"`
	CompletedAt string  `json:"completedAt" binding:"omitempty"`
	UpdatedBy   string  `json:"updatedBy" binding:"required,min=1,max=255"`
}

type GetProgressHatamByIdRequest struct {
	ID      string `uri:"id" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes juz and ayahs)"`
}

type GetProgressHatamByUserRequest struct {
	UserID  string `uri:"userId" binding:"required"`
	Limit   int    `form:"limit" binding:"omitempty,min=0,max=100"`
	Offset  int    `form:"offset" binding:"omitempty,min=0"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes juz and ayahs)"`
}

type GetProgressHatamByUserAndJuzRequest struct {
	UserID  string `uri:"userId" binding:"required"`
	JuzID   string `uri:"juzId" binding:"required"`
	Include string `form:"include" binding:"omitempty" example:"relations" enums:"relations" doc:"Include related data (relations - includes juz and ayahs)"`
}

type DeleteProgressHatamRequest struct {
	ID string `uri:"id" binding:"required"`
}
