package dto

// GetStoryByIDRequest represents the request to get a story by ID.
type GetStoryByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// CreateStoryRequest represents the request to create a new story.
type CreateStoryRequest struct {
	ID         string   `json:"id" binding:"required"`
	Title      string   `json:"title" binding:"required"`
	Summary    string   `json:"summary"`
	Characters []string `json:"characters"`
}

// UpdateStoryRequest represents the request to update a story.
type UpdateStoryRequest struct {
	ID         string   `uri:"id" binding:"required"`
	Title      string   `json:"title" binding:"omitempty"`
	Summary    string   `json:"summary" binding:"omitempty"`
	Characters []string `json:"characters" binding:"omitempty"`
}

// AddAyahToStoryRequest represents the request to link an ayah to a story.
type AddAyahToStoryRequest struct {
	StoryID string `json:"story_id" binding:"required"`
	SurahID int    `json:"surah_id" binding:"required,min=1,max=114"`
	AyahID  int    `json:"ayah_id" binding:"required,min=1"`
}

// RemoveAyahFromStoryRequest represents the request to unlink an ayah from a story.
type RemoveAyahFromStoryRequest struct {
	StoryID string `uri:"story_id" binding:"required"`
	SurahID int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID  int    `uri:"ayah_id" binding:"required,min=1"`
}
