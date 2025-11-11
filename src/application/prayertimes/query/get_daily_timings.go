package query

import (
	"context"

	appdto "github.com/safalife/core-api/src/application/prayertimes/dto"
	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

// GetDailyTimings fetches daily prayer timings via the domain service.
func (h *QueryHandler) GetDailyTimings(ctx context.Context, req *appdto.GetDailyTimingsRequest) (*appdto.PrayerTimesResponse, error) {
	if req == nil {
		return nil, dom.ErrInvalidRequest
	}
	dReq := &dom.TimingsRequest{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Date:      req.Date,
		Timezone:  req.Timezone,
		Method:    req.Method,
	}
	pt, err := h.svc.GetDailyTimings(ctx, dReq)
	if err != nil {
		return nil, err
	}
	return appdto.ToPrayerTimesResponse(pt), nil
}
