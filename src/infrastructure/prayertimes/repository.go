package prayertimes

import (
	"context"

	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

type CompositePrayerTimesRepository struct {
	primary  dom.PrayerTimesProviderInterface
	fallback dom.PrayerTimesProviderInterface
}

func NewCompositePrayerTimesRepository(primary, fallback dom.PrayerTimesProviderInterface) dom.PrayerTimesRepositoryInterface {
	return &CompositePrayerTimesRepository{primary: primary, fallback: fallback}
}

func (r *CompositePrayerTimesRepository) GetDailyTimings(ctx context.Context, req *dom.TimingsRequest) (*dom.PrayerTimes, error) {
	if req == nil {
		return nil, dom.ErrInvalidRequest
	}

	if r.primary != nil {
		if pt, err := r.primary.GetDailyTimings(ctx, req); err == nil && pt != nil {
			return pt, nil
		} else if err != nil {
			if err == dom.ErrInvalidRequest {
				return nil, err
			}
		}
	}

	if r.fallback != nil {
		return r.fallback.GetDailyTimings(ctx, req)
	}

	return nil, dom.ErrProviderUnavailable
}

func (r *CompositePrayerTimesRepository) GetMonthlyTimings(ctx context.Context, req *dom.TimingsRequest) ([]*dom.PrayerTimes, error) {
	if req == nil {
		return nil, dom.ErrInvalidRequest
	}

	if r.primary != nil {
		if pts, err := r.primary.GetMonthlyTimings(ctx, req); err == nil && pts != nil && len(pts) > 0 {
			return pts, nil
		} else if err != nil {
			if err == dom.ErrInvalidRequest {
				return nil, err
			}
		}
	}

	if r.fallback != nil {
		return r.fallback.GetMonthlyTimings(ctx, req)
	}

	return nil, dom.ErrProviderUnavailable
}

func (r *CompositePrayerTimesRepository) GetYearlyTimings(ctx context.Context, req *dom.TimingsRequest) ([]*dom.PrayerTimes, error) {
	if req == nil {
		return nil, dom.ErrInvalidRequest
	}

	if r.primary != nil {
		if pts, err := r.primary.GetYearlyTimings(ctx, req); err == nil && pts != nil && len(pts) > 0 {
			return pts, nil
		} else if err != nil {
			if err == dom.ErrInvalidRequest {
				return nil, err
			}
		}
	}

	if r.fallback != nil {
		return r.fallback.GetYearlyTimings(ctx, req)
	}

	return nil, dom.ErrProviderUnavailable
}
