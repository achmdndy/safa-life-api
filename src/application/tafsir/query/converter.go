package query

func ToGetAllTafsirsQuery() GetAllTafsirsQuery {
	return GetAllTafsirsQuery{}
}

func ToGetTafsirByIDQuery(id string) GetTafsirByIDQuery {
	return GetTafsirByIDQuery{ID: id}
}

func ToGetAyahTafsirQuery(tafsirID string, surahID, ayahID int) GetAyahTafsirQuery {
	return GetAyahTafsirQuery{
		TafsirID: tafsirID,
		SurahID:  surahID,
		AyahID:   ayahID,
	}
}

func ToGetTafsirsForAyahQuery(surahID, ayahID int) GetTafsirsForAyahQuery {
	return GetTafsirsForAyahQuery{
		SurahID: surahID,
		AyahID:  ayahID,
	}
}

func ToGetTafsirsForSurahQuery(tafsirID string, surahID int) GetTafsirsForSurahQuery {
	return GetTafsirsForSurahQuery{
		TafsirID: tafsirID,
		SurahID:  surahID,
	}
}
