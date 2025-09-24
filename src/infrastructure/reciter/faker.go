package reciter

import (
	"math/rand"
	"time"
)

// ReciterFaker generates fake Reciter data
type ReciterFaker struct{}

// NewReciterFaker creates a new ReciterFaker
func NewReciterFaker() *ReciterFaker {
	return &ReciterFaker{}
}

// Generate creates a fake ReciterModel
func (f *ReciterFaker) Generate() ReciterModel {
	reciters := []struct {
		ID    string
		Name  string
		Style string
	}{
		{"mishary_rashid", "Mishary Rashid Alafasy", "Hafs"},
		{"abdul_basit", "Abdul Basit Abdus Samad", "Hafs"},
		{"saud_shuraim", "Saud Al-Shuraim", "Hafs"},
	}
	reciter := reciters[rand.Intn(len(reciters))]
	return ReciterModel{
		ID:        reciter.ID,
		Name:      reciter.Name,
		Style:     reciter.Style,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
