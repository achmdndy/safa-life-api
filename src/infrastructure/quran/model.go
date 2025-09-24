package quran

import (
	"time"
	"gorm.io/gorm"
)

// SurahModel represents the Surah table in the database
type SurahModel struct {
	ID              int            `gorm:"primaryKey;column:surah_id" json:"surah_id"`
	NameArabic      string         `gorm:"column:name_ar;not null;size:255" json:"name_ar"`
	NameEnglish     string         `gorm:"column:name_en;not null;size:255" json:"name_en"`
	RevelationPlace string         `gorm:"column:revelation_place;not null;size:50" json:"revelation_place"`
	RevelationOrder int            `gorm:"column:revelation_order;not null" json:"revelation_order"`
	AyahCount       int            `gorm:"column:ayah_count;not null" json:"ayah_count"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	
	// Relationships
	Ayahs []AyahModel `gorm:"foreignKey:SurahID;references:ID" json:"ayahs,omitempty"`
}

// TableName specifies the table name for SurahModel
func (SurahModel) TableName() string {
	return "surahs"
}

// AyahModel represents the Ayah table in the database
type AyahModel struct {
	SurahID      int            `gorm:"primaryKey;column:surah_id" json:"surah_id"`
	AyahID       int            `gorm:"primaryKey;column:ayah_id" json:"ayah_id"`
	Text         string         `gorm:"column:text;not null;type:text" json:"text"`
	PageNumber   int            `gorm:"column:page_number;not null;index" json:"page_number"`
	JuzNumber    int            `gorm:"column:juz_number;not null;index" json:"juz_number"`
	HizbNumber   int            `gorm:"column:hizb_number;not null" json:"hizb_number"`
	ManzilNumber int            `gorm:"column:manzil_number;not null" json:"manzil_number"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	
	// Relationships
	Surah SurahModel `gorm:"foreignKey:SurahID;references:ID" json:"surah,omitempty"`
	Juz   JuzModel   `gorm:"foreignKey:JuzNumber;references:ID" json:"juz,omitempty"`
}

// TableName specifies the table name for AyahModel
func (AyahModel) TableName() string {
	return "ayahs"
}

// JuzModel represents the Juz table in the database
type JuzModel struct {
	ID         int            `gorm:"primaryKey;column:juz_id" json:"juz_id"`
	StartSurah int            `gorm:"column:start_surah;not null" json:"start_surah"`
	StartAyah  int            `gorm:"column:start_ayah;not null" json:"start_ayah"`
	EndSurah   int            `gorm:"column:end_surah;not null" json:"end_surah"`
	EndAyah    int            `gorm:"column:end_ayah;not null" json:"end_ayah"`
	CreatedAt  time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	
	// Relationships
	StartSurahModel SurahModel  `gorm:"foreignKey:StartSurah;references:ID" json:"start_surah_model,omitempty"`
	EndSurahModel   SurahModel  `gorm:"foreignKey:EndSurah;references:ID" json:"end_surah_model,omitempty"`
	Ayahs           []AyahModel `gorm:"foreignKey:JuzNumber;references:ID" json:"ayahs,omitempty"`
}

// TableName specifies the table name for JuzModel
func (JuzModel) TableName() string {
	return "juz"
}

// BeforeCreate hook for SurahModel
func (s *SurahModel) BeforeCreate(tx *gorm.DB) error {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for SurahModel
func (s *SurahModel) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}

// BeforeCreate hook for AyahModel
func (a *AyahModel) BeforeCreate(tx *gorm.DB) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for AyahModel
func (a *AyahModel) BeforeUpdate(tx *gorm.DB) error {
	a.UpdatedAt = time.Now()
	return nil
}

// BeforeCreate hook for JuzModel
func (j *JuzModel) BeforeCreate(tx *gorm.DB) error {
	j.CreatedAt = time.Now()
	j.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for JuzModel
func (j *JuzModel) BeforeUpdate(tx *gorm.DB) error {
	j.UpdatedAt = time.Now()
	return nil
}
