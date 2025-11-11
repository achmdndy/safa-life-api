package prayertimes

import (
	appQuery "github.com/safalife/core-api/src/application/prayertimes/query"
)

// Handler aggregates prayer times handlers
type Handler struct {
	getDailyTimingsHandler   *GetDailyTimingsHandler
	getMonthlyTimingsHandler *GetMonthlyTimingsHandler
	getYearlyTimingsHandler  *GetYearlyTimingsHandler
}

// NewHandler constructs the prayer times handler aggregator
func NewHandler(queryHandler *appQuery.QueryHandler) *Handler {
	return &Handler{
		getDailyTimingsHandler:   NewGetDailyTimingsHandler(queryHandler),
		getMonthlyTimingsHandler: NewGetMonthlyTimingsHandler(queryHandler),
		getYearlyTimingsHandler:  NewGetYearlyTimingsHandler(queryHandler),
	}
}

// Getters
func (h *Handler) GetDailyTimingsHandler() *GetDailyTimingsHandler {
	return h.getDailyTimingsHandler
}

func (h *Handler) GetMonthlyTimingsHandler() *GetMonthlyTimingsHandler {
	return h.getMonthlyTimingsHandler
}

func (h *Handler) GetYearlyTimingsHandler() *GetYearlyTimingsHandler {
	return h.getYearlyTimingsHandler
}
