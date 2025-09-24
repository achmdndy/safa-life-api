package dto

// GetReciterByIDRequest represents the request to get a reciter by ID.
type GetReciterByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// CreateReciterRequest represents the request to create a new reciter.
type CreateReciterRequest struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Style string `json:"style" binding:"required"`
}

// UpdateReciterRequest represents the request to update a reciter.
type UpdateReciterRequest struct {
	ID    string `uri:"id" binding:"required"`
	Name  string `json:"name" binding:"omitempty"`
	Style string `json:"style" binding:"omitempty"`
}
