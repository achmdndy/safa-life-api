package query

func ToGetAyahAudioQuery(reciterID string, surahID, ayahID int) GetAyahAudioQuery {
	return GetAyahAudioQuery{
		ReciterID: reciterID,
		SurahID:   surahID,
		AyahID:    ayahID,
	}
}

func ToGetAudioForSurahQuery(reciterID string, surahID int) GetAudioForSurahQuery {
	return GetAudioForSurahQuery{
		ReciterID: reciterID,
		SurahID:   surahID,
	}
}
