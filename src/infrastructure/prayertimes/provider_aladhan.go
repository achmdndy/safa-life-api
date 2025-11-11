package prayertimes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	dom "github.com/safalife/core-api/src/domain/prayertimes"
)

type AladhanProvider struct {
	httpClient *http.Client
	baseURL    string
}

func NewAladhanProvider(client *http.Client, baseURL string) dom.PrayerTimesProviderInterface {
	if client == nil {
		client = &http.Client{Timeout: 15 * time.Second}
	}
	return &AladhanProvider{httpClient: client, baseURL: strings.TrimRight(baseURL, "/")}
}

func (p *AladhanProvider) GetDailyTimings(ctx context.Context, req *dom.TimingsRequest) (*dom.PrayerTimes, error) {
	if req == nil || req.Timezone == "" {
		return nil, dom.ErrInvalidRequest
	}
	if p.baseURL == "" {
		return nil, dom.ErrProviderUnavailable
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
	dateStr := date.Format("02-01-2006")

	u, _ := url.Parse(p.baseURL + "/timings")
	q := u.Query()
	q.Set("latitude", fmt.Sprintf("%f", req.Latitude))
	q.Set("longitude", fmt.Sprintf("%f", req.Longitude))
	method := req.Method
	if method == 0 {
		method = 2
	}
	q.Set("method", fmt.Sprintf("%d", method))
	q.Set("date", dateStr)
	u.RawQuery = q.Encode()

	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, dom.ErrExternalAPI
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, dom.ErrExternalAPI
	}

	var payload struct {
		Code int `json:"code"`
		Data struct {
			Timings map[string]string `json:"timings"`
			Date    struct {
				Gregorian struct {
					Date string `json:"date"`
				} `json:"gregorian"`
			} `json:"date"`
			Meta struct {
				Timezone string `json:"timezone"`
			} `json:"meta"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, dom.ErrParsingAPIResponse
	}

	tz := req.Timezone
	if tz == "" && payload.Data.Meta.Timezone != "" {
		tz = payload.Data.Meta.Timezone
	}
	if tz == "" {
		tz = "UTC"
	}
	locOut, e := time.LoadLocation(tz)
	if e != nil {
		locOut = time.UTC
	}

	if gd := payload.Data.Date.Gregorian.Date; gd != "" {
		if t, e := time.ParseInLocation("02-01-2006", gd, locOut); e == nil {
			date = t
		}
	}

	parseClock := func(hm string) time.Time {
		hm = strings.TrimSpace(hm)
		if sp := strings.Split(hm, " "); len(sp) > 0 {
			hm = sp[0]
		}
		t, e := time.ParseInLocation("15:04", hm, locOut)
		if e != nil {
			return time.Time{}
		}
		return time.Date(date.Year(), date.Month(), date.Day(), t.Hour(), t.Minute(), 0, 0, locOut)
	}

	pt := &dom.PrayerTimes{
		Date:     date,
		Timezone: tz,
		Fajr:     parseClock(payload.Data.Timings["Fajr"]),
		Sunrise:  parseClock(payload.Data.Timings["Sunrise"]),
		Dhuhr:    parseClock(payload.Data.Timings["Dhuhr"]),
		Asr:      parseClock(payload.Data.Timings["Asr"]),
		Maghrib:  parseClock(payload.Data.Timings["Maghrib"]),
		Isha:     parseClock(payload.Data.Timings["Isha"]),
		Source:   "aladhan",
	}

	return pt, nil
}

func (p *AladhanProvider) GetMonthlyTimings(ctx context.Context, req *dom.TimingsRequest) ([]*dom.PrayerTimes, error) {
	if req == nil || req.Timezone == "" {
		return nil, dom.ErrInvalidRequest
	}
	if p.baseURL == "" {
		return nil, dom.ErrProviderUnavailable
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

	u, _ := url.Parse(p.baseURL + "/calendar")
	q := u.Query()
	q.Set("latitude", fmt.Sprintf("%f", req.Latitude))
	q.Set("longitude", fmt.Sprintf("%f", req.Longitude))
	method := req.Method
	if method == 0 {
		method = 2
	}
	q.Set("method", fmt.Sprintf("%d", method))
	q.Set("month", fmt.Sprintf("%d", int(base.Month())))
	q.Set("year", fmt.Sprintf("%d", base.Year()))
	u.RawQuery = q.Encode()

	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, dom.ErrExternalAPI
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, dom.ErrExternalAPI
	}

	var payload struct {
		Code int `json:"code"`
		Data []struct {
			Timings map[string]string `json:"timings"`
			Date    struct {
				Gregorian struct {
					Date string `json:"date"`
				} `json:"gregorian"`
			} `json:"date"`
			Meta struct {
				Timezone string `json:"timezone"`
			} `json:"meta"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, dom.ErrParsingAPIResponse
	}

	tz := req.Timezone
	if tz == "" && len(payload.Data) > 0 && payload.Data[0].Meta.Timezone != "" {
		tz = payload.Data[0].Meta.Timezone
	}
	if tz == "" {
		tz = "UTC"
	}
	locOut, e := time.LoadLocation(tz)
	if e != nil {
		locOut = time.UTC
	}

	parseClock := func(day time.Time, hm string) time.Time {
		hm = strings.TrimSpace(hm)
		if sp := strings.Split(hm, " "); len(sp) > 0 {
			hm = sp[0]
		}
		t, e := time.ParseInLocation("15:04", hm, locOut)
		if e != nil {
			return time.Time{}
		}
		return time.Date(day.Year(), day.Month(), day.Day(), t.Hour(), t.Minute(), 0, 0, locOut)
	}

	var out []*dom.PrayerTimes
	for _, d := range payload.Data {
		day := base
		if gd := d.Date.Gregorian.Date; gd != "" {
			if tt, e := time.ParseInLocation("02-01-2006", gd, locOut); e == nil {
				day = tt
			}
		}
		out = append(out, &dom.PrayerTimes{
			Date:     day,
			Timezone: tz,
			Fajr:     parseClock(day, d.Timings["Fajr"]),
			Sunrise:  parseClock(day, d.Timings["Sunrise"]),
			Dhuhr:    parseClock(day, d.Timings["Dhuhr"]),
			Asr:      parseClock(day, d.Timings["Asr"]),
			Maghrib:  parseClock(day, d.Timings["Maghrib"]),
			Isha:     parseClock(day, d.Timings["Isha"]),
			Source:   "aladhan",
		})
	}

	if len(out) == 0 {
		return nil, dom.ErrProviderUnavailable
	}
	return out, nil
}

func (p *AladhanProvider) GetYearlyTimings(ctx context.Context, req *dom.TimingsRequest) ([]*dom.PrayerTimes, error) {
	if req == nil || req.Timezone == "" {
		return nil, dom.ErrInvalidRequest
	}
	if p.baseURL == "" {
		return nil, dom.ErrProviderUnavailable
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

	var all []*dom.PrayerTimes
	for m := time.January; m <= time.December; m++ {
		monthDate := time.Date(base.Year(), m, 1, 0, 0, 0, 0, loc)
		subReq := &dom.TimingsRequest{
			Latitude:  req.Latitude,
			Longitude: req.Longitude,
			Date:      monthDate,
			Timezone:  req.Timezone,
			Method:    req.Method,
		}
		pts, err := p.GetMonthlyTimings(ctx, subReq)
		if err != nil {
			// Bila satu bulan gagal, lanjut ke bulan berikutnya
			continue
		}
		all = append(all, pts...)
	}

	if len(all) == 0 {
		return nil, dom.ErrProviderUnavailable
	}
	return all, nil
}
