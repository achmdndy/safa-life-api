package command

import (
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// FromCreateAyahTajweedRequest converts CreateAyahTajweedRequest to CreateAyahTajweedCommand
func FromCreateAyahTajweedRequest(req dto.CreateAyahTajweedRequest) CreateAyahTajweedCommand {
	words := make([]tajweed.TajweedWord, len(req.Words))
	for i, w := range req.Words {
		words[i] = tajweed.TajweedWord{Word: w.Word, Rule: w.Rule, Color: w.Color}
	}
	return CreateAyahTajweedCommand{
		TajweedID: req.TajweedID,
		SurahID:   req.SurahID,
		AyahID:    req.AyahID,
		Words:     words,
	}
}

// FromUpdateAyahTajweedRequest converts UpdateAyahTajweedRequest to UpdateAyahTajweedCommand
func FromUpdateAyahTajweedRequest(req dto.UpdateAyahTajweedRequest) UpdateAyahTajweedCommand {
	words := make([]tajweed.TajweedWord, len(req.Words))
	for i, w := range req.Words {
		words[i] = tajweed.TajweedWord{Word: w.Word, Rule: w.Rule, Color: w.Color}
	}
	return UpdateAyahTajweedCommand{
		TajweedID: req.TajweedID,
		SurahID:   req.SurahID,
		AyahID:    req.AyahID,
		Words:     words,
	}
}

// FromCreateTajweedRuleRequest converts CreateTajweedRuleRequest to CreateTajweedRuleCommand
func FromCreateTajweedRuleRequest(req dto.CreateTajweedRuleRequest) CreateTajweedRuleCommand {
	return CreateTajweedRuleCommand{
		ID:          req.ID,
		Rule:        req.Rule,
		Explanation: req.Explanation,
		Color:       req.Color,
	}
}

// FromUpdateTajweedRuleRequest converts UpdateTajweedRuleRequest to UpdateTajweedRuleCommand
func FromUpdateTajweedRuleRequest(req dto.UpdateTajweedRuleRequest) UpdateTajweedRuleCommand {
	cmd := UpdateTajweedRuleCommand{
		ID: req.ID,
	}
	if req.Rule != "" {
		cmd.Rule = &req.Rule
	}
	if req.Explanation != "" {
		cmd.Explanation = &req.Explanation
	}
	if req.Color != "" {
		cmd.Color = &req.Color
	}
	return cmd
}
