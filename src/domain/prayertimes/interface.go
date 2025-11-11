package prayertimes

import (
	"context"
)

type PrayerTimesServiceInterface interface {
	GetDailyTimings(ctx context.Context, req *TimingsRequest) (*PrayerTimes, error)
	GetMonthlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error)
	GetYearlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error)
}
