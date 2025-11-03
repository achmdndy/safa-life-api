package quran

import "errors"

// Domain errors for Quran entities
var (
	// Surah errors
	ErrSurahNotFound         = errors.New("surah not found")
	ErrSurahAlreadyExists    = errors.New("surah already exists")
	ErrInvalidSurahNumber    = errors.New("invalid surah number")
	ErrInvalidRevelationOrder = errors.New("invalid revelation order")

	// Ayah errors
	ErrAyahNotFound       = errors.New("ayah not found")
	ErrAyahAlreadyExists  = errors.New("ayah already exists")
	ErrInvalidAyahNumber  = errors.New("invalid ayah number")
	ErrInvalidPageNumber  = errors.New("invalid page number")
	ErrInvalidJuzNumber   = errors.New("invalid juz number")
	ErrInvalidHizbNumber  = errors.New("invalid hizb number")
	ErrInvalidManzilNumber = errors.New("invalid manzil number")

	// Juz errors
	ErrJuzNotFound        = errors.New("juz not found")
	ErrJuzAlreadyExists   = errors.New("juz already exists")
	ErrInvalidJuzRange    = errors.New("invalid juz range")
	ErrInvalidStartSurah  = errors.New("invalid start surah")
	ErrInvalidEndSurah    = errors.New("invalid end surah")
	ErrInvalidStartAyah   = errors.New("invalid start ayah")
	ErrInvalidEndAyah     = errors.New("invalid end ayah")

	// Validation errors
	ErrInvalidInput       = errors.New("invalid input provided")
	ErrInvalidUUID        = errors.New("invalid UUID format")
	ErrEmptyText          = errors.New("text cannot be empty")
	ErrEmptyName          = errors.New("name cannot be empty")
	ErrInvalidCreatedBy   = errors.New("created by cannot be empty")
)