package reciter

import "time"

// Reciter represents a Quran reciter.
type Reciter struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Style     string    `json:"style"` // e.g., Hafs, Warsh
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}