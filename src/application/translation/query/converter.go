package query

// ToGetAllTranslationsQuery creates a query for getting all translations.
func ToGetAllTranslationsQuery() GetAllTranslationsQuery {
	return GetAllTranslationsQuery{}
}

// ToGetTranslationByIDQuery creates a query for getting a translation by ID.
func ToGetTranslationByIDQuery(id string) GetTranslationByIDQuery {
	return GetTranslationByIDQuery{ID: id}
}

// ToGetAyahTranslationQuery creates a query for getting an ayah translation.
func ToGetAyahTranslationQuery(translationID string, surahID, ayahID int) GetAyahTranslationQuery {
	return GetAyahTranslationQuery{
		TranslationID: translationID,
		SurahID:       surahID,
		AyahID:        ayahID,
	}
}

// ToGetTranslationsForAyahQuery creates a query for getting all translations for an ayah.
func ToGetTranslationsForAyahQuery(surahID, ayahID int) GetTranslationsForAyahQuery {
	return GetTranslationsForAyahQuery{
		SurahID: surahID,
		AyahID:  ayahID,
	}
}

// ToGetTranslationsForSurahQuery creates a query for getting all translations for a surah.
func ToGetTranslationsForSurahQuery(translationID string, surahID int) GetTranslationsForSurahQuery {
	return GetTranslationsForSurahQuery{
		TranslationID: translationID,
		SurahID:       surahID,
	}
}
