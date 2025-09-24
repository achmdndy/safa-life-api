package command

import (
	"github.com/achmdndy/safa-life-api/src/application/story/dto"
)

func FromCreateStoryRequest(req dto.CreateStoryRequest) CreateStoryCommand {
	return CreateStoryCommand{
		ID:         req.ID,
		Title:      req.Title,
		Summary:    req.Summary,
		Characters: req.Characters,
	}
}

func FromUpdateStoryRequest(req dto.UpdateStoryRequest) UpdateStoryCommand {
	cmd := UpdateStoryCommand{ID: req.ID}
	if req.Title != "" {
		cmd.Title = &req.Title
	}
	if req.Summary != "" {
		cmd.Summary = &req.Summary
	}
	if len(req.Characters) > 0 {
		cmd.Characters = &req.Characters
	}
	return cmd
}

func FromAddAyahToStoryRequest(req dto.AddAyahToStoryRequest) AddAyahToStoryCommand {
	return AddAyahToStoryCommand{
		StoryID: req.StoryID,
		SurahID: req.SurahID,
		AyahID:  req.AyahID,
	}
}
