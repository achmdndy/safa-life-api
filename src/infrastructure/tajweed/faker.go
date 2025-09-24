package tajweed

import (
	"math/rand"
	"time"
)

// TajweedRuleFaker generates fake TajweedRule data.
type TajweedRuleFaker struct{}

// NewTajweedRuleFaker creates a new TajweedRuleFaker.
func NewTajweedRuleFaker() *TajweedRuleFaker {
	return &TajweedRuleFaker{}
}

// Generate creates a fake TajweedRuleModel.
func (f *TajweedRuleFaker) Generate() TajweedRuleModel {
	rules := []struct {
		ID          string
		Rule        string
		Explanation string
		Color       string
	}{
		{"idgham_bighunnah", "Idgham Bi Ghunnah", "Merging with nasalization.", "#FF0000"},
		{"idgham_bilaghunnah", "Idgham Bila Ghunnah", "Merging without nasalization.", "#00FF00"},
		{"ikhfa", "Ikhfa'", "Hiding the 'n' sound.", "#0000FF"},
		{"iqlab", "Iqlab", "Flipping the 'n' sound to 'm'.", "#FFFF00"},
		{"izhar", "Izhar", "Clear pronunciation.", "#FF00FF"},
	}
	rule := rules[rand.Intn(len(rules))]
	return TajweedRuleModel{
		ID:          rule.ID,
		Rule:        rule.Rule,
		Explanation: rule.Explanation,
		Color:       rule.Color,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// GenerateBatch creates multiple fake TajweedRuleModel instances.
func (f *TajweedRuleFaker) GenerateBatch(count int) []TajweedRuleModel {
	// Ensure we don't generate more unique rules than available
	rules := []struct {
		ID          string
		Rule        string
		Explanation string
		Color       string
	}{
		{"idgham_bighunnah", "Idgham Bi Ghunnah", "Merging with nasalization.", "#FF0000"},
		{"idgham_bilaghunnah", "Idgham Bila Ghunnah", "Merging without nasalization.", "#00FF00"},
		{"ikhfa", "Ikhfa'", "Hiding the 'n' sound.", "#0000FF"},
		{"iqlab", "Iqlab", "Flipping the 'n' sound to 'm'.", "#FFFF00"},
		{"izhar", "Izhar", "Clear pronunciation.", "#FF00FF"},
		{"madd_lazim", "Madd Lazim", "Necessary prolongation (6 harakat).", "#800080"},
		{"madd_wajib", "Madd Wajib Muttasil", "Connected obligatory prolongation (4-5 harakat).", "#FFA500"},
		{"madd_jaiz", "Madd Ja'iz Munfasil", "Separated permissible prolongation (2, 4, or 5 harakat).", "#008080"},
		{"qalqalah", "Qalqalah", "Echoing or bouncing sound.", "#A52A2A"},
	}

	if count > len(rules) {
		count = len(rules)
	}

	// Shuffle rules to get a random batch
	rand.Shuffle(len(rules), func(i, j int) { rules[i], rules[j] = rules[j], rules[i] })

	models := make([]TajweedRuleModel, count)
	for i := 0; i < count; i++ {
		rule := rules[i]
		models[i] = TajweedRuleModel{
			ID:          rule.ID,
			Rule:        rule.Rule,
			Explanation: rule.Explanation,
			Color:       rule.Color,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}
	return models
}

// AyahTajweedFaker generates fake AyahTajweed data.
type AyahTajweedFaker struct {
	rules []TajweedRuleModel
}

// NewAyahTajweedFaker creates a new AyahTajweedFaker.
func NewAyahTajweedFaker(rules []TajweedRuleModel) *AyahTajweedFaker {
	return &AyahTajweedFaker{rules: rules}
}

// Generate creates a fake AyahTajweedModel.
func (f *AyahTajweedFaker) Generate(tajweedID string, surahID, ayahID int, wordCount int) AyahTajweedModel {
	words := make([]TajweedWordModel, wordCount)
	arabicWords := []string{"بِسْمِ", "اللَّهِ", "الرَّحْمَٰنِ", "الرَّحِيمِ", "الْحَمْدُ", "لِلَّهِ", "رَبِّ", "الْعَالَمِينَ"}

	for i := 0; i < wordCount; i++ {
		rule := f.rules[rand.Intn(len(f.rules))]
		words[i] = TajweedWordModel{
			Word:  arabicWords[rand.Intn(len(arabicWords))],
			Rule:  rule.Rule,
			Color: rule.Color,
		}
	}

	return AyahTajweedModel{
		TajweedID: tajweedID,
		SurahID:   surahID,
		AyahID:    ayahID,
		Words:     words,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// GenerateBatch creates multiple fake AyahTajweedModel instances.
func (f *AyahTajweedFaker) GenerateBatch(tajweedID string, surahID int, ayahCount int, avgWordsPerAyah int) []AyahTajweedModel {
	models := make([]AyahTajweedModel, ayahCount)
	for i := 0; i < ayahCount; i++ {
		wordCount := avgWordsPerAyah + rand.Intn(5) - 2 // some variation
		if wordCount < 1 {
			wordCount = 1
		}
		models[i] = f.Generate(tajweedID, surahID, i+1, wordCount)
	}
	return models
}

// TajweedFaker combines all faker functionality.
type TajweedFaker struct {
	RuleFaker *TajweedRuleFaker
	AyahFaker *AyahTajweedFaker
}

// NewTajweedFaker creates a new TajweedFaker instance.
func NewTajweedFaker() *TajweedFaker {
	rand.Seed(time.Now().UnixNano())
	ruleFaker := NewTajweedRuleFaker()
	rules := ruleFaker.GenerateBatch(5) // Generate some rules to use
	ayahFaker := NewAyahTajweedFaker(rules)
	return &TajweedFaker{
		RuleFaker: ruleFaker,
		AyahFaker: ayahFaker,
	}
}
