package command

import (
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// Helper functions to convert from DTOs to commands

// FromCreateSurahRequest converts CreateSurahRequest to CreateSurahCommand
func FromCreateSurahRequest(req dto.CreateSurahRequest) CreateSurahCommand {
	return CreateSurahCommand{
		NameArabic:      req.NameArabic,
		NameEnglish:     req.NameEnglish,
		RevelationPlace: req.RevelationPlace,
		RevelationOrder: req.RevelationOrder,
		AyahCount:       req.AyahCount,
	}
}

// FromUpdateSurahRequest converts UpdateSurahRequest to UpdateSurahCommand
func FromUpdateSurahRequest(req dto.UpdateSurahRequest) UpdateSurahCommand {
	cmd := UpdateSurahCommand{
		ID: req.ID,
	}

	if req.NameArabic != "" {
		cmd.NameArabic = &req.NameArabic
	}
	if req.NameEnglish != "" {
		cmd.NameEnglish = &req.NameEnglish
	}
	if req.RevelationPlace != "" {
		cmd.RevelationPlace = &req.RevelationPlace
	}
	if req.RevelationOrder > 0 {
		cmd.RevelationOrder = &req.RevelationOrder
	}
	if req.AyahCount > 0 {
		cmd.AyahCount = &req.AyahCount
	}

	return cmd
}

// FromCreateAyahRequest converts CreateAyahRequest to CreateAyahCommand
func FromCreateAyahRequest(req dto.CreateAyahRequest) CreateAyahCommand {
	return CreateAyahCommand{
		SurahID:      req.SurahID,
		AyahID:       req.AyahID,
		Text:         req.Text,
		PageNumber:   req.PageNumber,
		JuzNumber:    req.JuzNumber,
		HizbNumber:   req.HizbNumber,
		ManzilNumber: req.ManzilNumber,
	}
}

// FromUpdateAyahRequest converts UpdateAyahRequest to UpdateAyahCommand
func FromUpdateAyahRequest(req dto.UpdateAyahRequest) UpdateAyahCommand {
	cmd := UpdateAyahCommand{
		SurahID: req.SurahID,
		AyahID:  req.AyahID,
	}

	if req.Text != "" {
		cmd.Text = &req.Text
	}
	if req.PageNumber > 0 {
		cmd.PageNumber = &req.PageNumber
	}
	if req.JuzNumber > 0 {
		cmd.JuzNumber = &req.JuzNumber
	}
	if req.HizbNumber > 0 {
		cmd.HizbNumber = &req.HizbNumber
	}
	if req.ManzilNumber > 0 {
		cmd.ManzilNumber = &req.ManzilNumber
	}

	return cmd
}

// FromCreateJuzRequest converts CreateJuzRequest to CreateJuzCommand
func FromCreateJuzRequest(req dto.CreateJuzRequest) CreateJuzCommand {
	return CreateJuzCommand{
		StartSurah: req.StartSurah,
		StartAyah:  req.StartAyah,
		EndSurah:   req.EndSurah,
		EndAyah:    req.EndAyah,
	}
}

// FromUpdateJuzRequest converts UpdateJuzRequest to UpdateJuzCommand
func FromUpdateJuzRequest(req dto.UpdateJuzRequest) UpdateJuzCommand {
	cmd := UpdateJuzCommand{
		ID: req.ID,
	}

	if req.StartSurah > 0 {
		cmd.StartSurah = &req.StartSurah
	}
	if req.StartAyah > 0 {
		cmd.StartAyah = &req.StartAyah
	}
	if req.EndSurah > 0 {
		cmd.EndSurah = &req.EndSurah
	}
	if req.EndAyah > 0 {
		cmd.EndAyah = &req.EndAyah
	}

	return cmd
}