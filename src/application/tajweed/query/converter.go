package query

// ToGetAyahTajweedQuery converts DTO to GetAyahTajweedQuery
func ToGetAyahTajweedQuery(tajweedID string, surahID, ayahID int) GetAyahTajweedQuery {
	return GetAyahTajweedQuery{
		TajweedID: tajweedID,
		SurahID:   surahID,
		AyahID:    ayahID,
	}
}

// ToGetAllTajweedRulesQuery converts DTO to GetAllTajweedRulesQuery
func ToGetAllTajweedRulesQuery() GetAllTajweedRulesQuery {
	return GetAllTajweedRulesQuery{}
}

// ToGetTajweedRuleByIDQuery converts DTO to GetTajweedRuleByIDQuery
func ToGetTajweedRuleByIDQuery(id string) GetTajweedRuleByIDQuery {
	return GetTajweedRuleByIDQuery{
		ID: id,
	}
}

// ToGetTajweedRuleByNameQuery converts DTO to GetTajweedRuleByNameQuery
func ToGetTajweedRuleByNameQuery(ruleName string) GetTajweedRuleByNameQuery {
	return GetTajweedRuleByNameQuery{
		RuleName: ruleName,
	}
}
