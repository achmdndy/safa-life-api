package core

import (
	quran "github.com/achmdndy/safa-life-api/src/infrastructure/quran"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: quran.SurahModel{}},
		{Model: quran.AyahModel{}},
		{Model: quran.JuzModel{}},
	}
}
