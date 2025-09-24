package audio

import "github.com/achmdndy/safa-life-api/src/domain/audio"

// ToAyahAudioDomain converts AyahAudioModel to domain.AyahAudio
func (m *AyahAudioModel) ToAyahAudioDomain() audio.AyahAudio {
	return audio.AyahAudio{
		ReciterID: m.ReciterID,
		SurahID:   m.SurahID,
		AyahID:    m.AyahID,
		FilePath:  m.FilePath,
		Duration:  m.Duration,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromAyahAudioDomain converts domain.AyahAudio to AyahAudioModel
func FromAyahAudioDomain(d *audio.AyahAudio) *AyahAudioModel {
	return &AyahAudioModel{
		ReciterID: d.ReciterID,
		SurahID:   d.SurahID,
		AyahID:    d.AyahID,
		FilePath:  d.FilePath,
		Duration:  d.Duration,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// ToAyahAudioDomainSlice converts a slice of AyahAudioModel to a slice of domain.AyahAudio
func ToAyahAudioDomainSlice(models []AyahAudioModel) []audio.AyahAudio {
	domains := make([]audio.AyahAudio, len(models))
	for i, m := range models {
		domains[i] = m.ToAyahAudioDomain()
	}
	return domains
}