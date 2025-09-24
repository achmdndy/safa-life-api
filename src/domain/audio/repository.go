package audio

import "context"

// AyahAudioRepository defines the interface for AyahAudio data access.
type AyahAudioRepository interface {
	Get(ctx context.Context, reciterID string, surahID, ayahID int) (*AyahAudio, error)
	GetBySurah(ctx context.Context, reciterID string, surahID int) ([]AyahAudio, error)
	Create(ctx context.Context, audio *AyahAudio) (*AyahAudio, error)
	Update(ctx context.Context, audio *AyahAudio) (*AyahAudio, error)
	Delete(ctx context.Context, reciterID string, surahID, ayahID int) error
	TransactionManager
}