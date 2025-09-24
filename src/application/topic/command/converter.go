package command

import (
	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
)

func FromCreateTopicRequest(req dto.CreateTopicRequest) CreateTopicCommand {
	return CreateTopicCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}
}

func FromUpdateTopicRequest(req dto.UpdateTopicRequest) UpdateTopicCommand {
	cmd := UpdateTopicCommand{ID: req.ID}
	if req.Name != "" {
		cmd.Name = &req.Name
	}
	if req.Description != "" {
		cmd.Description = &req.Description
	}
	return cmd
}

func FromAddAyahToTopicRequest(req dto.AddAyahToTopicRequest) AddAyahToTopicCommand {
	return AddAyahToTopicCommand{
		TopicID: req.TopicID,
		SurahID: req.SurahID,
		AyahID:  req.AyahID,
	}
}
