package query

import (
	"context"
	"time"

	appdto "github.com/safalife/core-api/src/application/prayertimes/dto"
	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

// GetYearlyTimings fetches yearly prayer timings via the domain service.
func (h *QueryHandler) GetYearlyTimings(ctx context.Context, req *appdto.GetYearlyTimingsRequest) (*appdto.PrayerTimesListResponse, error) {
	if req == nil {
		return nil, dom.ErrInvalidRequest
	}
	loc := time.UTC
	if req.Timezone != "" {
		if l, err := time.LoadLocation(req.Timezone); err == nil {
			loc = l
		}
	}
	if req.Year == 0 {
		return nil, dom.ErrInvalidRequest
	}
	base := time.Date(req.Year, time.January, 1, 0, 0, 0, 0, loc)
	dReq := &dom.TimingsRequest{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Date:      base,
		Timezone:  req.Timezone,
		Method:    req.Method,
	}
	pts, err := h.svc.GetYearlyTimings(ctx, dReq)
	if err != nil {
		return nil, err
	}
	return appdto.ToPrayerTimesListResponse(pts), nil
}
