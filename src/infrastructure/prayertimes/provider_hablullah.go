package prayertimes

import (
	"context"
	"time"

	prayer "github.com/hablullah/go-prayer"
	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

type HablullahProvider struct{}

func NewHablullahProvider() dom.PrayerTimesProviderInterface {
	return &HablullahProvider{}
}

func (p *HablullahProvider) GetDailyTimings(ctx context.Context, req *dom.TimingsRequest) (*dom.PrayerTimes, error) {
	if req == nil || req.Timezone == "" {
		return nil, dom.ErrInvalidRequest
	}

	loc, err := time.LoadLocation(req.Timezone)
	if err != nil {
		return nil, err
	}
	date := req.Date
	if date.IsZero() {
		date = time.Now().In(loc)
	} else {
		date = req.Date.In(loc)
	}

	schedules, err := prayer.Calculate(prayer.Config{
		Latitude:           req.Latitude,
		Longitude:          req.Longitude,
		Timezone:           loc,
		TwilightConvention: prayer.Kemenag(),
		AsrConvention:      prayer.Shafii,
		PreciseToSeconds:   true,
	}, date.Year())
	if err != nil {
		return nil, err
	}

	idx := date.YearDay() - 1
	if idx < 0 || idx >= len(schedules) {
		return nil, dom.ErrProviderUnavailable
	}
	s := schedules[idx]

	pt := &dom.PrayerTimes{
		Date:     date,
		Timezone: req.Timezone,
		Fajr:     s.Fajr,
		Sunrise:  s.Sunrise,
		Dhuhr:    s.Zuhr,
		Asr:      s.Asr,
		Maghrib:  s.Maghrib,
		Isha:     s.Isha,
		Source:   "go-prayer",
	}

	return pt, nil
}

func (p *HablullahProvider) GetMonthlyTimings(ctx context.Context, req *dom.TimingsRequest) ([]*dom.PrayerTimes, error) {
	if req == nil || req.Timezone == "" {
		return nil, dom.ErrInvalidRequest
	}

	loc, err := time.LoadLocation(req.Timezone)
	if err != nil {
		return nil, err
	}
	base := req.Date
	if base.IsZero() {
		base = time.Now().In(loc)
	} else {
		base = req.Date.In(loc)
	}

	schedules, err := prayer.Calculate(prayer.Config{
		Latitude:           req.Latitude,
		Longitude:          req.Longitude,
		Timezone:           loc,
		TwilightConvention: prayer.Kemenag(),
		AsrConvention:      prayer.Shafii,
		PreciseToSeconds:   true,
	}, base.Year())
	if err != nil {
		return nil, err
	}

	start := time.Date(base.Year(), base.Month(), 1, 0, 0, 0, 0, loc)
	var out []*dom.PrayerTimes
	for d := start; d.Month() == base.Month(); d = d.AddDate(0, 0, 1) {
		idx := d.YearDay() - 1
		if idx < 0 || idx >= len(schedules) {
			continue
		}
		s := schedules[idx]
		out = append(out, &dom.PrayerTimes{
			Date:     d,
			Timezone: req.Timezone,
			Fajr:     s.Fajr,
			Sunrise:  s.Sunrise,
			Dhuhr:    s.Zuhr,
			Asr:      s.Asr,
			Maghrib:  s.Maghrib,
			Isha:     s.Isha,
			Source:   "go-prayer",
		})
	}

	if len(out) == 0 {
		return nil, dom.ErrProviderUnavailable
	}
	return out, nil
}

func (p *HablullahProvider) GetYearlyTimings(ctx context.Context, req *dom.TimingsRequest) ([]*dom.PrayerTimes, error) {
	if req == nil || req.Timezone == "" {
		return nil, dom.ErrInvalidRequest
	}

	loc, err := time.LoadLocation(req.Timezone)
	if err != nil {
		return nil, err
	}
	base := req.Date
	if base.IsZero() {
		base = time.Now().In(loc)
	} else {
		base = req.Date.In(loc)
	}

	schedules, err := prayer.Calculate(prayer.Config{
		Latitude:           req.Latitude,
		Longitude:          req.Longitude,
		Timezone:           loc,
		TwilightConvention: prayer.Kemenag(),
		AsrConvention:      prayer.Shafii,
		PreciseToSeconds:   true,
	}, base.Year())
	if err != nil {
		return nil, err
	}

	start := time.Date(base.Year(), time.January, 1, 0, 0, 0, 0, loc)
	var out []*dom.PrayerTimes
	for d := start; d.Year() == base.Year(); d = d.AddDate(0, 0, 1) {
		idx := d.YearDay() - 1
		if idx < 0 || idx >= len(schedules) {
			continue
		}
		s := schedules[idx]
		out = append(out, &dom.PrayerTimes{
			Date:     d,
			Timezone: req.Timezone,
			Fajr:     s.Fajr,
			Sunrise:  s.Sunrise,
			Dhuhr:    s.Zuhr,
			Asr:      s.Asr,
			Maghrib:  s.Maghrib,
			Isha:     s.Isha,
			Source:   "go-prayer",
		})
	}

	if len(out) == 0 {
		return nil, dom.ErrProviderUnavailable
	}
	return out, nil
}
