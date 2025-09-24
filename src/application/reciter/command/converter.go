package command

import (
	"github.com/achmdndy/safa-life-api/src/application/reciter/dto"
)

// FromCreateReciterRequest converts a DTO to a CreateReciterCommand.
func FromCreateReciterRequest(req dto.CreateReciterRequest) CreateReciterCommand {
	return CreateReciterCommand{
		ID:    req.ID,
		Name:  req.Name,
		Style: req.Style,
	}
}

// FromUpdateReciterRequest converts a DTO to an UpdateReciterCommand.
func FromUpdateReciterRequest(req dto.UpdateReciterRequest) UpdateReciterCommand {
	cmd := UpdateReciterCommand{ID: req.ID}
	if req.Name != "" {
		cmd.Name = &req.Name
	}
	if req.Style != "" {
		cmd.Style = &req.Style
	}
	return cmd
}
