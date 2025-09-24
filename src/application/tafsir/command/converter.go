package command

import (
	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
)

func FromCreateTafsirRequest(req dto.CreateTafsirRequest) CreateTafsirCommand {
	return CreateTafsirCommand{
		ID:       req.ID,
		Name:     req.Name,
		Author:   req.Author,
		Language: req.Language,
	}
}

func FromUpdateTafsirRequest(req dto.UpdateTafsirRequest) UpdateTafsirCommand {
	cmd := UpdateTafsirCommand{ID: req.ID}
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

func FromCreateAyahTafsirRequest(req dto.CreateAyahTafsirRequest) CreateAyahTafsirCommand {
	return CreateAyahTafsirCommand{
		TafsirID: req.TafsirID,
		SurahID:  req.SurahID,
		AyahID:   req.AyahID,
		Text:     req.Text,
	}
}

func FromUpdateAyahTafsirRequest(req dto.UpdateAyahTafsirRequest) UpdateAyahTafsirCommand {
	cmd := UpdateAyahTafsirCommand{
		TafsirID: req.TafsirID,
		SurahID:  req.SurahID,
		AyahID:   req.AyahID,
	}
	if req.Text != "" {
		cmd.Text = &req.Text
	}
	return cmd
}
