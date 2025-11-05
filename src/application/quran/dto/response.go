package dto

import (
	"time"

	"github.com/safalife/core-api/src/domain/quran"
)

// Surah Response DTOs
type SurahResponse struct {
	ID              string    `json:"id"`
	NameArabic      string    `json:"nameArabic"`
	NameEnglish     string    `json:"nameEnglish"`
	RevelationPlace string    `json:"revelationPlace"`
	RevelationOrder int       `json:"revelationOrder"`
	AyahCount       int       `json:"ayahCount"`
	CreatedBy       string    `json:"createdBy"`
	UpdatedBy       string    `json:"updatedBy"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// ToSurahResponse converts a domain Surah to its DTO.
func ToSurahResponse(d *quran.Surah) *SurahResponse {
	if d == nil {
		return nil
	}
	return &SurahResponse{
		ID:              d.ID.String(),
		NameArabic:      d.NameArabic,
		NameEnglish:     d.NameEnglish,
		RevelationPlace: d.RevelationPlace,
		RevelationOrder: d.RevelationOrder,
		AyahCount:       d.AyahCount,
		CreatedBy:       d.CreatedBy,
		UpdatedBy:       d.UpdatedBy,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

// ToSurahResponseSlice converts a slice of domain Surahs to a slice of DTOs.
func ToSurahResponseSlice(ds []*quran.Surah) []*SurahResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*SurahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToSurahResponse(d)
	}
	return outs
}

// Ayah Response DTOs
type AyahResponse struct {
	ID           string    `json:"id"`
	SurahID      string    `json:"surahId"`
	Text         string    `json:"text"`
	PageNumber   int       `json:"pageNumber"`
	JuzNumber    int       `json:"juzNumber"`
	HizbNumber   int       `json:"hizbNumber"`
	ManzilNumber int       `json:"manzilNumber"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// ToAyahResponse converts a domain Ayah to its DTO.
func ToAyahResponse(d *quran.Ayah) *AyahResponse {
	if d == nil {
		return nil
	}
	return &AyahResponse{
		ID:           d.ID.String(),
		SurahID:      d.SurahID.String(),
		Text:         d.Text,
		PageNumber:   d.PageNumber,
		JuzNumber:    d.JuzNumber,
		HizbNumber:   d.HizbNumber,
		ManzilNumber: d.ManzilNumber,
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

// ToAyahResponseSlice converts a slice of domain Ayahs to a slice of DTOs.
func ToAyahResponseSlice(ds []*quran.Ayah) []*AyahResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*AyahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahResponse(d)
	}
	return outs
}

// Juz Response DTOs
type JuzResponse struct {
	ID           string    `json:"id"`
	StartSurahID string    `json:"startSurahId"`
	EndSurahID   string    `json:"endSurahId"`
	StartAyahID  string    `json:"startAyahId"`
	EndAyahID    string    `json:"endAyahId"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// ToJuzResponse converts a domain Juz to its DTO.
func ToJuzResponse(d *quran.Juz) *JuzResponse {
	if d == nil {
		return nil
	}
	return &JuzResponse{
		ID:           d.ID.String(),
		StartSurahID: d.StartSurahID.String(),
		EndSurahID:   d.EndSurahID.String(),
		StartAyahID:  d.StartAyahID.String(),
		EndAyahID:    d.EndAyahID.String(),
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

// ToJuzResponseSlice converts a slice of domain Juz to a slice of DTOs.
func ToJuzResponseSlice(ds []*quran.Juz) []*JuzResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*JuzResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToJuzResponse(d)
	}
	return outs
}

// Pagination Response
type PaginationResponse struct {
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
}

// Surah List Response with Pagination
type SurahListResponse struct {
	Data       []*SurahResponse    `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// Ayah List Response with Pagination
type AyahListResponse struct {
	Data       []*AyahResponse     `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// Juz List Response with Pagination
type JuzListResponse struct {
	Data       []*JuzResponse      `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// Count Response
type CountResponse struct {
	Count int64 `json:"count"`
}

// Surah with Ayahs Response DTO
type SurahWithAyahsResponse struct {
	ID              string          `json:"id"`
	NameArabic      string          `json:"nameArabic"`
	NameEnglish     string          `json:"nameEnglish"`
	RevelationPlace string          `json:"revelationPlace"`
	RevelationOrder int             `json:"revelationOrder"`
	AyahCount       int             `json:"ayahCount"`
	CreatedBy       string          `json:"createdBy"`
	UpdatedBy       string          `json:"updatedBy"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	Ayahs           []*AyahResponse `json:"ayahs,omitempty"`
}

// ToSurahWithAyahsResponse converts a domain SurahWithAyahs to its DTO.
func ToSurahWithAyahsResponse(d *quran.SurahWithAyahs) *SurahWithAyahsResponse {
	if d == nil {
		return nil
	}
	return &SurahWithAyahsResponse{
		ID:              d.ID.String(),
		NameArabic:      d.NameArabic,
		NameEnglish:     d.NameEnglish,
		RevelationPlace: d.RevelationPlace,
		RevelationOrder: d.RevelationOrder,
		AyahCount:       d.AyahCount,
		CreatedBy:       d.CreatedBy,
		UpdatedBy:       d.UpdatedBy,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
		Ayahs:           ToAyahResponseSlice(d.Ayahs),
	}
}

// ToSurahWithAyahsResponseSlice converts a slice of domain SurahWithAyahs to a slice of DTOs.
func ToSurahWithAyahsResponseSlice(ds []*quran.SurahWithAyahs) []*SurahWithAyahsResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*SurahWithAyahsResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToSurahWithAyahsResponse(d)
	}
	return outs
}

// Ayah with Surah Response DTO
type AyahWithSurahResponse struct {
	ID           string         `json:"id"`
	SurahID      string         `json:"surahId"`
	Text         string         `json:"text"`
	PageNumber   int            `json:"pageNumber"`
	JuzNumber    int            `json:"juzNumber"`
	HizbNumber   int            `json:"hizbNumber"`
	ManzilNumber int            `json:"manzilNumber"`
	CreatedBy    string         `json:"createdBy"`
	UpdatedBy    string         `json:"updatedBy"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	Surah        *SurahResponse `json:"surah,omitempty"`
}

// ToAyahWithSurahResponse converts a domain AyahWithSurah to its DTO.
func ToAyahWithSurahResponse(d *quran.AyahWithSurah) *AyahWithSurahResponse {
	if d == nil {
		return nil
	}
	return &AyahWithSurahResponse{
		ID:           d.ID.String(),
		SurahID:      d.SurahID.String(),
		Text:         d.Text,
		PageNumber:   d.PageNumber,
		JuzNumber:    d.JuzNumber,
		HizbNumber:   d.HizbNumber,
		ManzilNumber: d.ManzilNumber,
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		Surah:        ToSurahResponse(d.Surah),
	}
}

// ToAyahWithSurahResponseSlice converts a slice of domain AyahWithSurah to a slice of DTOs.
func ToAyahWithSurahResponseSlice(ds []*quran.AyahWithSurah) []*AyahWithSurahResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*AyahWithSurahResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahWithSurahResponse(d)
	}
	return outs
}

// Juz with Relations Response DTO
type JuzWithRelationsResponse struct {
	ID           string         `json:"id"`
	StartSurahID string         `json:"startSurahId"`
	EndSurahID   string         `json:"endSurahId"`
	StartAyahID  string         `json:"startAyahId"`
	EndAyahID    string         `json:"endAyahId"`
	CreatedBy    string         `json:"createdBy"`
	UpdatedBy    string         `json:"updatedBy"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	StartSurah   *SurahResponse `json:"startSurah,omitempty"`
	EndSurah     *SurahResponse `json:"endSurah,omitempty"`
	StartAyah    *AyahResponse  `json:"startAyah,omitempty"`
	EndAyah      *AyahResponse  `json:"endAyah,omitempty"`
}

// ToJuzWithRelationsResponse converts a domain JuzWithRelations to its DTO.
func ToJuzWithRelationsResponse(d *quran.JuzWithRelations) *JuzWithRelationsResponse {
	if d == nil {
		return nil
	}
	return &JuzWithRelationsResponse{
		ID:           d.ID.String(),
		StartSurahID: d.StartSurahID.String(),
		EndSurahID:   d.EndSurahID.String(),
		StartAyahID:  d.StartAyahID.String(),
		EndAyahID:    d.EndAyahID.String(),
		CreatedBy:    d.CreatedBy,
		UpdatedBy:    d.UpdatedBy,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		StartSurah:   ToSurahResponse(d.StartSurah),
		EndSurah:     ToSurahResponse(d.EndSurah),
		StartAyah:    ToAyahResponse(d.StartAyah),
		EndAyah:      ToAyahResponse(d.EndAyah),
	}
}

// ToJuzWithRelationsResponseSlice converts a slice of domain JuzWithRelations to a slice of DTOs.
func ToJuzWithRelationsResponseSlice(ds []*quran.JuzWithRelations) []*JuzWithRelationsResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*JuzWithRelationsResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToJuzWithRelationsResponse(d)
	}
	return outs
}

// List responses with relations
type SurahWithAyahsListResponse struct {
	Data       []*SurahWithAyahsResponse `json:"data"`
	Pagination *PaginationResponse       `json:"pagination"`
}

type AyahWithSurahListResponse struct {
	Data       []*AyahWithSurahResponse `json:"data"`
	Pagination *PaginationResponse      `json:"pagination"`
}

type JuzWithRelationsListResponse struct {
	Data       []*JuzWithRelationsResponse `json:"data"`
	Pagination *PaginationResponse         `json:"pagination"`
}

// Ayah with Translation Response DTO
type AyahWithTranslationResponse struct {
	ID           string                   `json:"id"`
	SurahID      string                   `json:"surahId"`
	Text         string                   `json:"text"`
	PageNumber   int                      `json:"pageNumber"`
	JuzNumber    int                      `json:"juzNumber"`
	HizbNumber   int                      `json:"hizbNumber"`
	ManzilNumber int                      `json:"manzilNumber"`
	CreatedBy    string                   `json:"createdBy"`
	UpdatedBy    string                   `json:"updatedBy"`
	CreatedAt    time.Time                `json:"createdAt"`
	UpdatedAt    time.Time                `json:"updatedAt"`
	Translation  *AyahTranslationResponse `json:"translation,omitempty"`
	Audio        *AyahAudioFileResponse   `json:"audio,omitempty"`
}

// ToAyahWithTranslationResponse builds AyahWithTranslationResponse from domain ayah and translation
func ToAyahWithTranslationResponse(ayah *quran.Ayah, tr *quran.AyahTranslation) *AyahWithTranslationResponse {
	if ayah == nil {
		return nil
	}
	resp := &AyahWithTranslationResponse{
		ID:           ayah.ID.String(),
		SurahID:      ayah.SurahID.String(),
		Text:         ayah.Text,
		PageNumber:   ayah.PageNumber,
		JuzNumber:    ayah.JuzNumber,
		HizbNumber:   ayah.HizbNumber,
		ManzilNumber: ayah.ManzilNumber,
		CreatedBy:    ayah.CreatedBy,
		UpdatedBy:    ayah.UpdatedBy,
		CreatedAt:    ayah.CreatedAt,
		UpdatedAt:    ayah.UpdatedAt,
	}
	if tr != nil {
		resp.Translation = ToAyahTranslationResponse(tr)
	}
	return resp
}

// Surah + Ayahs with per-ayah Translation Response
type SurahAyahsWithTranslationsResponse struct {
	Surah      *SurahResponse                 `json:"surah"`
	Ayahs      []*AyahWithTranslationResponse `json:"ayahs"`
	Pagination *PaginationResponse            `json:"pagination"`
	Reciter    *ReciterResponse               `json:"reciter,omitempty"`
}

// Translation Edition Response DTOs
type TranslationEditionResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Language  string    `json:"language"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ToTranslationEditionResponse converts a domain TranslationEdition to its DTO.
func ToTranslationEditionResponse(d *quran.TranslationEdition) *TranslationEditionResponse {
	if d == nil {
		return nil
	}
	return &TranslationEditionResponse{
		ID:        d.ID.String(),
		Name:      d.Name,
		Author:    d.Author,
		Language:  d.Language,
		CreatedBy: d.CreatedBy,
		UpdatedBy: d.UpdatedBy,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToTranslationEditionResponseSlice converts a slice of domain TranslationEdition to DTOs.
func ToTranslationEditionResponseSlice(ds []*quran.TranslationEdition) []*TranslationEditionResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*TranslationEditionResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToTranslationEditionResponse(d)
	}
	return outs
}

// Translation Edition List Response with Pagination
type TranslationEditionListResponse struct {
	Data       []*TranslationEditionResponse `json:"data"`
	Pagination *PaginationResponse           `json:"pagination"`
}

// Ayah Translation Response DTOs
type AyahTranslationResponse struct {
	ID                   string    `json:"id"`
	TranslationEditionID string    `json:"translationEditionId"`
	SurahID              string    `json:"surahId"`
	AyahID               string    `json:"ayahId"`
	Text                 string    `json:"text"`
	CreatedBy            string    `json:"createdBy"`
	UpdatedBy            string    `json:"updatedBy"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

// ToAyahTranslationResponse converts a domain AyahTranslation to its DTO.
func ToAyahTranslationResponse(d *quran.AyahTranslation) *AyahTranslationResponse {
	if d == nil {
		return nil
	}
	return &AyahTranslationResponse{
		ID:                   d.ID.String(),
		TranslationEditionID: d.TranslationEditionID.String(),
		SurahID:              d.SurahID.String(),
		AyahID:               d.AyahID.String(),
		Text:                 d.Text,
		CreatedBy:            d.CreatedBy,
		UpdatedBy:            d.UpdatedBy,
		CreatedAt:            d.CreatedAt,
		UpdatedAt:            d.UpdatedAt,
	}
}

// ToAyahTranslationResponseSlice converts a slice of domain AyahTranslation to DTOs.
func ToAyahTranslationResponseSlice(ds []*quran.AyahTranslation) []*AyahTranslationResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*AyahTranslationResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahTranslationResponse(d)
	}
	return outs
}

// Ayah Translation List Response with Pagination
type AyahTranslationListResponse struct {
	Data       []*AyahTranslationResponse `json:"data"`
	Pagination *PaginationResponse        `json:"pagination"`
}

// Reciter Response DTOs
type ReciterResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Style     string    `json:"style"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ToReciterResponse converts a domain Reciter to its DTO.
func ToReciterResponse(d *quran.Reciter) *ReciterResponse {
	if d == nil {
		return nil
	}
	return &ReciterResponse{
		ID:        d.ID.String(),
		Name:      d.Name,
		Style:     d.Style,
		CreatedBy: d.CreatedBy,
		UpdatedBy: d.UpdatedBy,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToReciterResponseSlice converts a slice of domain Reciters to a slice of DTOs.
func ToReciterResponseSlice(ds []*quran.Reciter) []*ReciterResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*ReciterResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToReciterResponse(d)
	}
	return outs
}

// Reciter List Response with Pagination
type ReciterListResponse struct {
	Data       []*ReciterResponse  `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// AyahAudioFile Response DTOs
type AyahAudioFileResponse struct {
	ID        string    `json:"id"`
	ReciterID string    `json:"reciterId"`
	SurahID   string    `json:"surahId"`
	AyahID    string    `json:"ayahId"`
	FilePath  string    `json:"filePath"`
	Duration  float64   `json:"duration"`
	ByteSize  float64   `json:"byteSize"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ToAyahAudioFileResponse converts a domain AyahAudioFile to its DTO.
func ToAyahAudioFileResponse(d *quran.AyahAudioFile) *AyahAudioFileResponse {
	if d == nil {
		return nil
	}
	return &AyahAudioFileResponse{
		ID:        d.ID.String(),
		ReciterID: d.ReciterID.String(),
		SurahID:   d.SurahID.String(),
		AyahID:    d.AyahID.String(),
		FilePath:  d.FilePath,
		Duration:  d.Duration,
		ByteSize:  d.ByteSize,
		CreatedBy: d.CreatedBy,
		UpdatedBy: d.UpdatedBy,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToAyahAudioFileResponseSlice converts a slice of domain AyahAudioFiles to a slice of DTOs.
func ToAyahAudioFileResponseSlice(ds []*quran.AyahAudioFile) []*AyahAudioFileResponse {
	if ds == nil {
		return nil
	}
	outs := make([]*AyahAudioFileResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToAyahAudioFileResponse(d)
	}
	return outs
}

// AyahAudioFile List Response with Pagination
type AyahAudioFileListResponse struct {
	Data       []*AyahAudioFileResponse `json:"data"`
	Pagination *PaginationResponse      `json:"pagination"`
}
