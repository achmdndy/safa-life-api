package prayertimes

import "time"

type TimingsRequest struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Date      time.Time `json:"date"`
	Timezone  string    `json:"timezone"`
	Method    int       `json:"method"`
}

type PrayerTimes struct {
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
