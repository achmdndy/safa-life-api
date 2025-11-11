package query

import (
	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

type QueryHandler struct {
	svc dom.PrayerTimesServiceInterface
}

func NewQueryHandler(svc dom.PrayerTimesServiceInterface) *QueryHandler {
	return &QueryHandler{svc: svc}
}
