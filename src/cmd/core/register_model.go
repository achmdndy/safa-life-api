package core

import (
	"github.com/achmdndy/safa-life-api/src/infrastructure/audio"
	"github.com/achmdndy/safa-life-api/src/infrastructure/quran"
	"github.com/achmdndy/safa-life-api/src/infrastructure/reciter"
	"github.com/achmdndy/safa-life-api/src/infrastructure/resource"
	"github.com/achmdndy/safa-life-api/src/infrastructure/story"
	"github.com/achmdndy/safa-life-api/src/infrastructure/tafsir"
	"github.com/achmdndy/safa-life-api/src/infrastructure/tajweed"
	"github.com/achmdndy/safa-life-api/src/infrastructure/topic"
	"github.com/achmdndy/safa-life-api/src/infrastructure/translation"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		// Quran
		{Model: quran.SurahModel{}},
		{Model: quran.AyahModel{}},
		{Model: quran.JuzModel{}},
		// Audio
		{Model: audio.AyahAudioModel{}},
		// Reciter
		{Model: reciter.ReciterModel{}},
		// Resource
		{Model: resource.ResourceModel{}},
		// Story
		{Model: story.StoryModel{}},
		{Model: story.StoryAyahModel{}},
		// Tafsir
		{Model: tafsir.TafsirModel{}},
		{Model: tafsir.AyahTafsirModel{}},
		// Tajweed
		{Model: tajweed.TajweedRuleModel{}},
		{Model: tajweed.AyahTajweedModel{}},
		// Topic
		{Model: topic.TopicModel{}},
		{Model: topic.TopicAyahModel{}},
		// Translation
		{Model: translation.TranslationModel{}},
		{Model: translation.AyahTranslationModel{}},
	}
}