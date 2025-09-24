package dto

import (
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/reciter"
)

// ReciterResponse represents a reciter.
type ReciterResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Style     string    `json:"style"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToReciterResponse converts a domain Reciter to its DTO.
func ToReciterResponse(d reciter.Reciter) ReciterResponse {
	return ReciterResponse{
		ID:        d.ID,
		Name:      d.Name,
		Style:     d.Style,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToReciterResponseSlice converts a slice of domain Reciters to a slice of DTOs.
func ToReciterResponseSlice(ds []reciter.Reciter) []ReciterResponse {
	outs := make([]ReciterResponse, len(ds))
	for i, d := range ds {
		outs[i] = ToReciterResponse(d)
	}
	return outs
}
