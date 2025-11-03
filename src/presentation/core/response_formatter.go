package core

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type EmptyData struct{}

type ErrorDetail struct {
	Code   *string `json:"code,omitempty"`
	Reason string  `json:"reason"`
}

// Swagger response types for documentation
// These types are used specifically for Swagger documentation since it doesn't support Go generics

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Success      bool        `json:"success" example:"true"`
	StatusCode   int         `json:"statusCode" example:"200"`
	Message      string      `json:"message" example:"Operation completed successfully"`
	Timestamp    time.Time   `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string      `json:"responseTime" example:"15.234ms"`
	Data         interface{} `json:"data,omitempty"`
	Errors       interface{} `json:"errors,omitempty"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success      bool         `json:"success" example:"false"`
	StatusCode   int          `json:"statusCode" example:"500"`
	Message      string       `json:"message" example:"An error occurred"`
	Timestamp    time.Time    `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	ResponseTime string       `json:"responseTime" example:"15.234ms"`
	Data         interface{}  `json:"data,omitempty"`
	Errors       *ErrorDetail `json:"errors,omitempty"`
}

type ResponseFormatter[D any, E any] struct {
	Success      bool      `json:"success"`
	StatusCode   int       `json:"statusCode"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timestamp"`
	ResponseTime string    `json:"responseTime"`
	Data         *D        `json:"data,omitempty"`
	Errors       *E        `json:"errors,omitempty"`
}

func Success[D any](c *gin.Context, statusCode int, message string, data D, start time.Time) {
	elapsed := time.Since(start).Seconds() * 1000.0
	resp := ResponseFormatter[D, any]{
		Success:      true,
		StatusCode:   statusCode,
		Message:      message,
		Timestamp:    time.Now().UTC(),
		ResponseTime: FormatMillis(elapsed),
		Data:         &data,
		Errors:       nil,
	}
	c.JSON(statusCode, resp)
}

func Error[E any](c *gin.Context, statusCode int, message string, errors *E, start time.Time) {
	elapsed := time.Since(start).Seconds() * 1000.0
	resp := ResponseFormatter[any, E]{
		Success:      false,
		StatusCode:   statusCode,
		Message:      message,
		Timestamp:    time.Now().UTC(),
		ResponseTime: FormatMillis(elapsed),
		Data:         nil,
		Errors:       errors,
	}
	c.JSON(statusCode, resp)
}

func FormatMillis(ms float64) string {
	return fmt.Sprintf("%.3fms", ms)
}
