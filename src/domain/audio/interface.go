package audio

import "context"

// AudioService defines the business logic for audio file operations.
type AudioService interface {
	GetAyahAudio(ctx context.Context, reciterID string, surahID, ayahID int) (*AyahAudio, error)
	GetAudioFilesForSurah(ctx context.Context, reciterID string, surahID int) ([]AyahAudio, error)
	CreateAyahAudio(ctx context.Context, audio *AyahAudio) (*AyahAudio, error)
	UpdateAyahAudio(ctx context.Context, audio *AyahAudio) (*AyahAudio, error)
	DeleteAyahAudio(ctx context.Context, reciterID string, surahID, ayahID int) error
}

// TransactionManager defines the interface for database transaction management.
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}