package query

import (
	"context"
	"time"

	appdto "github.com/safalife/core-api/src/application/prayertimes/dto"
	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

// GetMonthlyTimings fetches monthly prayer timings via the domain service.
func (h *QueryHandler) GetMonthlyTimings(ctx context.Context, req *appdto.GetMonthlyTimingsRequest) (*appdto.PrayerTimesListResponse, error) {
	if req == nil {
		return nil, dom.ErrInvalidRequest
	}
	// Build base date at the first day of requested month/year in the specified timezone.
	loc := time.UTC
	if req.Timezone != "" {
		if l, err := time.LoadLocation(req.Timezone); err == nil {
			loc = l
		}
	}
	if req.Month <= 0 || req.Month > 12 || req.Year == 0 {
		return nil, dom.ErrInvalidRequest
	}
	base := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, loc)
	dReq := &dom.TimingsRequest{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Date:      base,
		Timezone:  req.Timezone,
		Method:    req.Method,
	}
	pts, err := h.svc.GetMonthlyTimings(ctx, dReq)
	if err != nil {
		return nil, err
	}
	return appdto.ToPrayerTimesListResponse(pts), nil
}
