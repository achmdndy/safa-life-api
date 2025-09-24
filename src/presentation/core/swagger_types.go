package core

import "time"

// SuccessResponse represents a successful API response
// @Description Successful API response structure
type SuccessResponse struct {
	Success      bool        `json:"success" example:"true"`
	StatusCode   int         `json:"status_code" example:"200"`
	Message      string      `json:"message" example:"Operation completed successfully"`
	Timestamp    time.Time   `json:"timestamp" example:"2023-01-01T00:00:00Z"`
	ResponseTime string      `json:"response_time" example:"15.234ms"`
	Data         interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error API response
// @Description Error API response structure
type ErrorResponse struct {
	Success      bool        `json:"success" example:"false"`
	StatusCode   int         `json:"status_code" example:"400"`
	Message      string      `json:"message" example:"Operation failed"`
	Timestamp    time.Time   `json:"timestamp" example:"2023-01-01T00:00:00Z"`
	ResponseTime string      `json:"response_time" example:"15.234ms"`
	Errors       interface{} `json:"errors,omitempty"`
}
