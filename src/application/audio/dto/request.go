package dto

// GetAyahAudioRequest represents the request to get a specific ayah audio.
type GetAyahAudioRequest struct {
	ReciterID string `uri:"reciter_id" binding:"required"`
	SurahID   int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID    int    `uri:"ayah_id" binding:"required,min=1"`
}

// GetAudioForSurahRequest represents the request to get all audio files for a surah.
type GetAudioForSurahRequest struct {
	ReciterID string `uri:"reciter_id" binding:"required"`
	SurahID   int    `uri:"surah_id" binding:"required,min=1,max=114"`
}

// CreateAyahAudioRequest represents the request to create a new ayah audio file.
type CreateAyahAudioRequest struct {
	ReciterID string  `json:"reciter_id" binding:"required"`
	SurahID   int     `json:"surah_id" binding:"required,min=1,max=114"`
	AyahID    int     `json:"ayah_id" binding:"required,min=1"`
	FilePath  string  `json:"file_path" binding:"required"`
	Duration  float64 `json:"duration" binding:"required,min=0"`
}

// UpdateAyahAudioRequest represents the request to update an ayah audio file.
type UpdateAyahAudioRequest struct {
	ReciterID string  `uri:"reciter_id" binding:"required"`
	SurahID   int     `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID    int     `uri:"ayah_id" binding:"required,min=1"`
	FilePath  string  `json:"file_path" binding:"omitempty"`
	Duration  float64 `json:"duration" binding:"omitempty,min=0"`
}
