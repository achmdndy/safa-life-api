package command

import (
	"github.com/achmdndy/safa-life-api/src/application/audio/dto"
)

func FromCreateAyahAudioRequest(req dto.CreateAyahAudioRequest) CreateAyahAudioCommand {
	return CreateAyahAudioCommand{
		ReciterID: req.ReciterID,
		SurahID:   req.SurahID,
		AyahID:    req.AyahID,
		FilePath:  req.FilePath,
		Duration:  req.Duration,
	}
}

func FromUpdateAyahAudioRequest(req dto.UpdateAyahAudioRequest) UpdateAyahAudioCommand {
	cmd := UpdateAyahAudioCommand{
		ReciterID: req.ReciterID,
		SurahID:   req.SurahID,
		AyahID:    req.AyahID,
	}
	if req.FilePath != "" {
		cmd.FilePath = &req.FilePath
	}
	if req.Duration > 0 {
		cmd.Duration = &req.Duration
	}
	return cmd
}
