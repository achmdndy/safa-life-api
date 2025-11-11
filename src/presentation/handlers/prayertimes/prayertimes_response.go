package prayertimes

import (
	"time"

	appdto "github.com/safalife/core-api/src/application/prayertimes/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// Prayer Times Response Types for Swagger documentation

// Daily response
type PrayerTimesSuccessResponse struct {
	Success      bool                        `json:"success" example:"true"`
	StatusCode   int                         `json:"status_code" example:"200"`
	Message      string                      `json:"message" example:"Prayer times retrieved successfully"`
	Timestamp    time.Time                   `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string                      `json:"responseTime" example:"15.234ms"`
	Data         *appdto.PrayerTimesResponse `json:"data"`
}

// Monthly/Yearly list response
type PrayerTimesListSuccessResponse struct {
	Success      bool                            `json:"success" example:"true"`
	StatusCode   int                             `json:"status_code" example:"200"`
	Message      string                          `json:"message" example:"Prayer times list retrieved successfully"`
	Timestamp    time.Time                       `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string                          `json:"responseTime" example:"15.234ms"`
	Data         *appdto.PrayerTimesListResponse `json:"data"`
}

// Error response
type PrayerTimesErrorResponse struct {
	Success      bool             `json:"success" example:"false"`
	StatusCode   int              `json:"status_code" example:"500"`
	Message      string           `json:"message" example:"Failed to process request"`
	Timestamp    time.Time        `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string           `json:"responseTime" example:"15.234ms"`
	Errors       core.ErrorDetail `json:"errors"`
}
