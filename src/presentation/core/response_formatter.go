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

type ResponseFormatter[D any, E any] struct {
	Success      bool      `json:"success"`
	StatusCode   int       `json:"status_code"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timestamp"`
	ResponseTime string    `json:"response_time"`
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
