package command

import (
	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
)

// FromCreateTranslationRequest converts a DTO to a CreateTranslationCommand.
func FromCreateTranslationRequest(req dto.CreateTranslationRequest) CreateTranslationCommand {
	return CreateTranslationCommand{
		ID:       req.ID,
		Name:     req.Name,
		Author:   req.Author,
		Language: req.Language,
	}
}

// FromUpdateTranslationRequest converts a DTO to an UpdateTranslationCommand.
func FromUpdateTranslationRequest(req dto.UpdateTranslationRequest) UpdateTranslationCommand {
	cmd := UpdateTranslationCommand{ID: req.ID}
	if req.Name != "" {
		cmd.Name = &req.Name
	}
	if req.Author != "" {
		cmd.Author = &req.Author
	}
	if req.Language != "" {
		cmd.Language = &req.Language
	}
	return cmd
}

// FromCreateAyahTranslationRequest converts a DTO to a CreateAyahTranslationCommand.
func FromCreateAyahTranslationRequest(req dto.CreateAyahTranslationRequest) CreateAyahTranslationCommand {
	return CreateAyahTranslationCommand{
		TranslationID: req.TranslationID,
		SurahID:       req.SurahID,
		AyahID:        req.AyahID,
		Text:          req.Text,
	}
}

// FromUpdateAyahTranslationRequest converts a DTO to an UpdateAyahTranslationCommand.
func FromUpdateAyahTranslationRequest(req dto.UpdateAyahTranslationRequest) UpdateAyahTranslationCommand {
	cmd := UpdateAyahTranslationCommand{
		TranslationID: req.TranslationID,
		SurahID:       req.SurahID,
		AyahID:        req.AyahID,
	}
	if req.Text != "" {
		cmd.Text = &req.Text
	}
	return cmd
}
