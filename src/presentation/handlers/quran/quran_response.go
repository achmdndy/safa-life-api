package quran

import (
	"time"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// Quran-specific response types for Swagger documentation
// These types provide concrete examples instead of interface{} to avoid "additionalProp1" in Swagger UI

// Surah Response Types
type SurahSuccessResponse struct {
	Success      bool              `json:"success" example:"true"`
	StatusCode   int               `json:"status_code" example:"200"`
	Message      string            `json:"message" example:"Surah retrieved successfully"`
	Timestamp    time.Time         `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string            `json:"responseTime" example:"15.234ms"`
	Data         dto.SurahResponse `json:"data"`
}

type SurahListSuccessResponse struct {
	Success      bool                  `json:"success" example:"true"`
	StatusCode   int                   `json:"status_code" example:"200"`
	Message      string                `json:"message" example:"Surah list retrieved successfully"`
	Timestamp    time.Time             `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string                `json:"responseTime" example:"15.234ms"`
	Data         dto.SurahListResponse `json:"data"`
}

// Ayah Response Types
type AyahSuccessResponse struct {
	Success      bool             `json:"success" example:"true"`
	StatusCode   int              `json:"status_code" example:"200"`
	Message      string           `json:"message" example:"Ayah retrieved successfully"`
	Timestamp    time.Time        `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string           `json:"responseTime" example:"15.234ms"`
	Data         dto.AyahResponse `json:"data"`
}

type AyahListSuccessResponse struct {
	Success      bool                 `json:"success" example:"true"`
	StatusCode   int                  `json:"status_code" example:"200"`
	Message      string               `json:"message" example:"Ayah list retrieved successfully"`
	Timestamp    time.Time            `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string               `json:"responseTime" example:"15.234ms"`
	Data         dto.AyahListResponse `json:"data"`
}

// Juz Response Types
type JuzSuccessResponse struct {
	Success      bool            `json:"success" example:"true"`
	StatusCode   int             `json:"status_code" example:"200"`
	Message      string          `json:"message" example:"Juz retrieved successfully"`
	Timestamp    time.Time       `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string          `json:"responseTime" example:"15.234ms"`
	Data         dto.JuzResponse `json:"data"`
}

type JuzListSuccessResponse struct {
	Success      bool                `json:"success" example:"true"`
	StatusCode   int                 `json:"status_code" example:"200"`
	Message      string              `json:"message" example:"Juz list retrieved successfully"`
	Timestamp    time.Time           `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string              `json:"responseTime" example:"15.234ms"`
	Data         dto.JuzListResponse `json:"data"`
}

// Create Response Types
type CreateSurahSuccessResponse struct {
	Success      bool              `json:"success" example:"true"`
	StatusCode   int               `json:"status_code" example:"201"`
	Message      string            `json:"message" example:"Surah created successfully"`
	Timestamp    time.Time         `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string            `json:"responseTime" example:"15.234ms"`
	Data         dto.SurahResponse `json:"data"`
}

type CreateAyahSuccessResponse struct {
	Success      bool             `json:"success" example:"true"`
	StatusCode   int              `json:"status_code" example:"201"`
	Message      string           `json:"message" example:"Ayah created successfully"`
	Timestamp    time.Time        `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string           `json:"responseTime" example:"15.234ms"`
	Data         dto.AyahResponse `json:"data"`
}

// Update Response Types
type UpdateSurahSuccessResponse struct {
	Success      bool              `json:"success" example:"true"`
	StatusCode   int               `json:"status_code" example:"200"`
	Message      string            `json:"message" example:"Surah updated successfully"`
	Timestamp    time.Time         `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string            `json:"responseTime" example:"15.234ms"`
	Data         dto.SurahResponse `json:"data"`
}

type UpdateAyahSuccessResponse struct {
	Success      bool             `json:"success" example:"true"`
	StatusCode   int              `json:"status_code" example:"200"`
	Message      string           `json:"message" example:"Ayah updated successfully"`
	Timestamp    time.Time        `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string           `json:"responseTime" example:"15.234ms"`
	Data         dto.AyahResponse `json:"data"`
}

// Delete Response Types
type DeleteSurahSuccessResponse struct {
	Success      bool      `json:"success" example:"true"`
	StatusCode   int       `json:"status_code" example:"200"`
	Message      string    `json:"message" example:"Surah deleted successfully"`
	Timestamp    time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string    `json:"responseTime" example:"15.234ms"`
	Data         *string   `json:"data" swaggertype:"string" example:"null"`
}

type DeleteAyahSuccessResponse struct {
	Success      bool      `json:"success" example:"true"`
	StatusCode   int       `json:"status_code" example:"200"`
	Message      string    `json:"message" example:"Ayah deleted successfully"`
	Timestamp    time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string    `json:"responseTime" example:"15.234ms"`
	Data         *string   `json:"data" swaggertype:"string" example:"null"`
}

// Special Response Types
type GetAllSurahsSuccessResponse struct {
	Success      bool                  `json:"success" example:"true"`
	StatusCode   int                   `json:"status_code" example:"200"`
	Message      string                `json:"message" example:"Surahs retrieved successfully"`
	Timestamp    time.Time             `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string                `json:"responseTime" example:"15.234ms"`
	Data         dto.SurahListResponse `json:"data"`
}

type GetSurahByIdSuccessResponse struct {
	Success      bool              `json:"success" example:"true"`
	StatusCode   int               `json:"status_code" example:"200"`
	Message      string            `json:"message" example:"Surah retrieved successfully"`
	Timestamp    time.Time         `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string            `json:"responseTime" example:"15.234ms"`
	Data         dto.SurahResponse `json:"data"`
}

type GetSurahByNumberSuccessResponse struct {
	Success      bool              `json:"success" example:"true"`
	StatusCode   int               `json:"status_code" example:"200"`
	Message      string            `json:"message" example:"Surah retrieved successfully"`
	Timestamp    time.Time         `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string            `json:"responseTime" example:"15.234ms"`
	Data         dto.SurahResponse `json:"data"`
}

type GetAyahByIdSuccessResponse struct {
	Success      bool             `json:"success" example:"true"`
	StatusCode   int              `json:"status_code" example:"200"`
	Message      string           `json:"message" example:"Ayah retrieved successfully"`
	Timestamp    time.Time        `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string           `json:"responseTime" example:"15.234ms"`
	Data         dto.AyahResponse `json:"data"`
}

type GetAyahsBySurahSuccessResponse struct {
	Success      bool                 `json:"success" example:"true"`
	StatusCode   int                  `json:"status_code" example:"200"`
	Message      string               `json:"message" example:"Ayahs retrieved successfully"`
	Timestamp    time.Time            `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string               `json:"responseTime" example:"15.234ms"`
	Data         dto.AyahListResponse `json:"data"`
}

type GetAyahsByJuzSuccessResponse struct {
	Success      bool                 `json:"success" example:"true"`
	StatusCode   int                  `json:"status_code" example:"200"`
	Message      string               `json:"message" example:"Ayahs retrieved successfully"`
	Timestamp    time.Time            `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string               `json:"responseTime" example:"15.234ms"`
	Data         dto.AyahListResponse `json:"data"`
}

// Error Response Type
type QuranErrorResponse struct {
	Success      bool             `json:"success" example:"false"`
	StatusCode   int              `json:"status_code" example:"500"`
	Message      string           `json:"message" example:"Failed to process request"`
	Timestamp    time.Time        `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string           `json:"responseTime" example:"15.234ms"`
	Errors       core.ErrorDetail `json:"errors"`
}
