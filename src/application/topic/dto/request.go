package dto

// GetTopicByIDRequest represents the request to get a topic by ID.
type GetTopicByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// CreateTopicRequest represents the request to create a new topic.
type CreateTopicRequest struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// UpdateTopicRequest represents the request to update a topic.
type UpdateTopicRequest struct {
	ID          string `uri:"id" binding:"required"`
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
}

// AddAyahToTopicRequest represents the request to link an ayah to a topic.
type AddAyahToTopicRequest struct {
	TopicID string `json:"topic_id" binding:"required"`
	SurahID int    `json:"surah_id" binding:"required,min=1,max=114"`
	AyahID  int    `json:"ayah_id" binding:"required,min=1"`
}

// RemoveAyahFromTopicRequest represents the request to unlink an ayah from a topic.
type RemoveAyahFromTopicRequest struct {
	TopicID string `uri:"topic_id" binding:"required"`
	SurahID int    `uri:"surah_id" binding:"required,min=1,max=114"`
	AyahID  int    `uri:"ayah_id" binding:"required,min=1"`
}
