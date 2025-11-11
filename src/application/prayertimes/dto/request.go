package dto

import (
	"time"
)

// GetDailyTimingsRequest represents the request to fetch daily prayer timings.
type GetDailyTimingsRequest struct {
	Latitude  float64   `form:"latitude" binding:"required"`
	Longitude float64   `form:"longitude" binding:"required"`
	Date      time.Time `form:"date" time_format:"2006-01-02" binding:"omitempty"`
	Timezone  string    `form:"timezone" binding:"required"`
	Method    int       `form:"method" binding:"omitempty"`
}

// GetMonthlyTimingsRequest represents the request to fetch monthly prayer timings.
type GetMonthlyTimingsRequest struct {
	Latitude  float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
	Month     int     `form:"month" binding:"required,min=1,max=12"`
	Year      int     `form:"year" binding:"required,min=1900,max=3000"`
	Timezone  string  `form:"timezone" binding:"required"`
	Method    int     `form:"method" binding:"omitempty"`
}

// GetYearlyTimingsRequest represents the request to fetch yearly prayer timings.
type GetYearlyTimingsRequest struct {
	Latitude  float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
	Year      int     `form:"year" binding:"required,min=1900,max=3000"`
	Timezone  string  `form:"timezone" binding:"required"`
	Method    int     `form:"method" binding:"omitempty"`
}
