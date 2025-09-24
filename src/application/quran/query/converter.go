package query

// ToGetAllSurahsQuery converts DTO to GetAllSurahsQuery
func ToGetAllSurahsQuery() GetAllSurahsQuery {
	return GetAllSurahsQuery{}
}

// ToGetSurahByIDQuery converts DTO to GetSurahByIDQuery
func ToGetSurahByIDQuery(id int) GetSurahByIDQuery {
	return GetSurahByIDQuery{
		ID: id,
	}
}

// ToGetSurahWithAyahsQuery converts DTO to GetSurahWithAyahsQuery
func ToGetSurahWithAyahsQuery(id int) GetSurahWithAyahsQuery {
	return GetSurahWithAyahsQuery{
		ID: id,
	}
}

// ToGetAyahsBySurahQuery converts DTO to GetAyahsBySurahQuery
func ToGetAyahsBySurahQuery(surahID, page, limit int) GetAyahsBySurahQuery {
	return GetAyahsBySurahQuery{
		SurahID: surahID,
		Page:    page,
		Limit:   limit,
	}
}

// ToGetAyahByIDQuery converts DTO to GetAyahByIDQuery
func ToGetAyahByIDQuery(surahID, ayahID int) GetAyahByIDQuery {
	return GetAyahByIDQuery{
		SurahID: surahID,
		AyahID:  ayahID,
	}
}

// ToGetAyahsByPageQuery converts DTO to GetAyahsByPageQuery
func ToGetAyahsByPageQuery(page int) GetAyahsByPageQuery {
	return GetAyahsByPageQuery{
		Page: page,
	}
}

// ToGetAyahsByJuzQuery converts DTO to GetAyahsByJuzQuery
func ToGetAyahsByJuzQuery(juzID int) GetAyahsByJuzQuery {
	return GetAyahsByJuzQuery{
		JuzID: juzID,
	}
}

// ToGetAllJuzQuery converts DTO to GetAllJuzQuery
func ToGetAllJuzQuery() GetAllJuzQuery {
	return GetAllJuzQuery{}
}

// ToGetJuzByIDQuery converts DTO to GetJuzByIDQuery
func ToGetJuzByIDQuery(id int) GetJuzByIDQuery {
	return GetJuzByIDQuery{
		ID: id,
	}
}

// ToGetJuzWithContentQuery converts DTO to GetJuzWithContentQuery
func ToGetJuzWithContentQuery(id int) GetJuzWithContentQuery {
	return GetJuzWithContentQuery{
		ID: id,
	}
}

// ToSearchQuery converts DTO to SearchQuery
func ToSearchQuery(query string, page, limit int) SearchQuery {
	return SearchQuery{
		Query: query,
		Page:  page,
		Limit: limit,
	}
}
