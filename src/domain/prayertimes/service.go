package prayertimes

import "context"

type PrayerTimesService struct {
	repo PrayerTimesRepositoryInterface
}

func NewPrayerTimesService(repo PrayerTimesRepositoryInterface) PrayerTimesServiceInterface {
	return &PrayerTimesService{repo: repo}
}

func (s *PrayerTimesService) GetDailyTimings(ctx context.Context, req *TimingsRequest) (*PrayerTimes, error) {
	return s.repo.GetDailyTimings(ctx, req)
}

func (s *PrayerTimesService) GetMonthlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error) {
	return s.repo.GetMonthlyTimings(ctx, req)
}

func (s *PrayerTimesService) GetYearlyTimings(ctx context.Context, req *TimingsRequest) ([]*PrayerTimes, error) {
	return s.repo.GetYearlyTimings(ctx, req)
}
