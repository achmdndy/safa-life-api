package audio

import (
	"fmt"
	"math/rand"
	"time"
)

// AyahAudioFaker generates fake AyahAudio data
type AyahAudioFaker struct{}

// NewAyahAudioFaker creates a new AyahAudioFaker
func NewAyahAudioFaker() *AyahAudioFaker {
	rand.Seed(time.Now().UnixNano())
	return &AyahAudioFaker{}
}

// Generate creates a fake AyahAudioModel
func (f *AyahAudioFaker) Generate(reciterID string, surahID, ayahID int) AyahAudioModel {
	return AyahAudioModel{
		ReciterID: reciterID,
		SurahID:   surahID,
		AyahID:    ayahID,
		FilePath:  fmt.Sprintf("local://audio/%s/%d_%d.mp3", reciterID, surahID, ayahID),
		Duration:  rand.Float64()*10 + 5, // 5 to 15 seconds
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
