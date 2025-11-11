package dto

import (
	"time"

	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

// PrayerTimesResponse represents the response DTO for daily prayer timings.
type PrayerTimesResponse struct {
	Date     time.Time `json:"date"`
	Timezone string    `json:"timezone"`

	Fajr    time.Time `json:"fajr"`
	Sunrise time.Time `json:"sunrise"`
	Dhuhr   time.Time `json:"dhuhr"`
	Asr     time.Time `json:"asr"`
	Maghrib time.Time `json:"maghrib"`
	Isha    time.Time `json:"isha"`

	Source string `json:"source"`
}

// ToPrayerTimesResponse converts a domain PrayerTimes to its DTO.
func ToPrayerTimesResponse(d *dom.PrayerTimes) *PrayerTimesResponse {
	if d == nil {
		return nil
	}
	return &PrayerTimesResponse{
		Date:     d.Date,
		Timezone: d.Timezone,
		Fajr:     d.Fajr,
		Sunrise:  d.Sunrise,
		Dhuhr:    d.Dhuhr,
		Asr:      d.Asr,
		Maghrib:  d.Maghrib,
		Isha:     d.Isha,
		Source:   d.Source,
	}
}

// PrayerTimesListResponse represents a list of prayer timings entries (e.g., monthly/yearly).
type PrayerTimesListResponse struct {
	Items []*PrayerTimesResponse `json:"items"`
}

// ToPrayerTimesListResponse converts a slice of domain PrayerTimes to list DTO.
func ToPrayerTimesListResponse(ds []*dom.PrayerTimes) *PrayerTimesListResponse {
	if ds == nil {
		return &PrayerTimesListResponse{Items: nil}
	}
	items := make([]*PrayerTimesResponse, 0, len(ds))
	for _, d := range ds {
		items = append(items, ToPrayerTimesResponse(d))
	}
	return &PrayerTimesListResponse{Items: items}
}
