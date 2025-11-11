package prayertimes

import "context"

type PrayerTimesProviderInterface interface {
	GetDailyTimings(ctx context.Context, req *TimingsRequest) (*PrayerTimes, error)
	GetMonthlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error)
	GetYearlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error)
}

type PrayerTimesRepositoryInterface interface {
	GetDailyTimings(ctx context.Context, req *TimingsRequest) (*PrayerTimes, error)
	GetMonthlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error)
	GetYearlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error)
}
