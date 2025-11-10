package quran

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SurahModel struct {
	ID              uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	NameArabic      string         `gorm:"column:name_ar;not null;size:255" json:"name_ar"`
	NameEnglish     string         `gorm:"column:name_en;not null;size:255" json:"name_en"`
	RevelationPlace string         `gorm:"column:revelation_place;not null;size:50" json:"revelation_place"`
	RevelationOrder int            `gorm:"column:revelation_order;not null" json:"revelation_order"`
	AyahCount       int            `gorm:"column:ayah_count;not null" json:"ayah_count"`
	CreatedBy       string         `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy       string         `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Ayahs []AyahModel `gorm:"foreignKey:surah_id;references:id" json:"ayahs,omitempty"`
}

func (SurahModel) TableName() string {
	return "surahs"
}

type AyahModel struct {
	ID           uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	SurahID      uuid.UUID      `gorm:"column:surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"surah_id"`
	Text         string         `gorm:"column:text;not null;type:text" json:"text"`
	PageNumber   int            `gorm:"column:page_number;not null;index" json:"page_number"`
	JuzNumber    int            `gorm:"column:juz_number;not null;index" json:"juz_number"`
	HizbNumber   int            `gorm:"column:hizb_number;not null" json:"hizb_number"`
	ManzilNumber int            `gorm:"column:manzil_number;not null" json:"manzil_number"`
	CreatedBy    string         `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy    string         `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Surah SurahModel `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
}

func (AyahModel) TableName() string {
	return "ayahs"
}

type JuzModel struct {
	ID           uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	StartSurahID uuid.UUID      `gorm:"column:start_surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"start_surah_id"`
	EndSurahID   uuid.UUID      `gorm:"column:end_surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"end_surah_id"`
	StartAyahID  uuid.UUID      `gorm:"column:start_ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"start_ayah_id"`
	EndAyahID    uuid.UUID      `gorm:"column:end_ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"end_ayah_id"`
	CreatedBy    string         `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy    string         `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	StartSurah SurahModel `gorm:"foreignKey:start_surah_id;references:id" json:"start_surah_model,omitempty"`
	EndSurah   SurahModel `gorm:"foreignKey:end_surah_id;references:id" json:"end_surah_model,omitempty"`
	StartAyah  AyahModel  `gorm:"foreignKey:start_ayah_id;references:id" json:"start_ayah_model,omitempty"`
	EndAyah    AyahModel  `gorm:"foreignKey:end_ayah_id;references:id" json:"end_ayah_model,omitempty"`
}

func (JuzModel) TableName() string {
	return "juz"
}

type ReciterModel struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"column:name;not null;unique;size:255"`
	Style     string         `gorm:"column:style;not null;size:100"`
	Place     string         `gorm:"column:place;not null;size:20;default:Unknown"`
	Picture   string         `gorm:"column:picture;not null;size:255;default:Unknown"`
	CreatedBy string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (ReciterModel) TableName() string {
	return "reciters"
}

type AyahAudioFileModel struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ReciterID uuid.UUID      `gorm:"column:reciter_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"reciter_id"`
	SurahID   uuid.UUID      `gorm:"column:surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"surah_id"`
	AyahID    uuid.UUID      `gorm:"column:ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ayah_id"`
	FilePath  string         `gorm:"column:file_path;not null;size:512"`
	Duration  float64        `gorm:"column:duration"`
	ByteSize  float64        `gorm:"column:byte_size"`
	CreatedBy string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Reciter ReciterModel `gorm:"foreignKey:reciter_id;references:id" json:"reciter,omitempty"`
	Surah   SurahModel   `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
	Ayah    AyahModel    `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
}

func (AyahAudioFileModel) TableName() string {
	return "ayah_audio_files"
}

type TranslationEditionModel struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"column:name;not null;size:255"`
	Author    string         `gorm:"column:author;not null;size:255"`
	Language  string         `gorm:"column:language;not null;size:50"`
	CreatedBy string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (TranslationEditionModel) TableName() string {
	return "translation_editions"
}

type AyahTranslationModel struct {
	ID                   uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	TranslationEditionID uuid.UUID      `gorm:"column:translation_edition_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"translation_edition_id"`
	SurahID              uuid.UUID      `gorm:"column:surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"surah_id"`
	AyahID               uuid.UUID      `gorm:"column:ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ayah_id"`
	Text                 string         `gorm:"column:text;not null;type:text"`
	CreatedBy            string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy            string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt            time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	TranslationEdition TranslationEditionModel `gorm:"foreignKey:translation_edition_id;references:id" json:"translation_edition,omitempty"`
	Surah              SurahModel              `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
	Ayah               AyahModel               `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
}

func (AyahTranslationModel) TableName() string {
	return "ayah_translations"
}

type BookmarkAyahModel struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    string         `gorm:"column:user_id;type:varchar(255);not null;index;" json:"user_id"`
	AyahID    uuid.UUID      `gorm:"column:ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ayah_id"`
	CreatedBy string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Ayah AyahModel `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
}

func (BookmarkAyahModel) TableName() string {
	return "bookmark_ayahs"
}

type LastReadModel struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      string    `gorm:"column:user_id;type:varchar(255);not null;index" json:"user_id"`
	SurahID     uuid.UUID `gorm:"column:surah_id;type:uuid;not null;index:idx_user_surah" json:"surah_id"`
	AyahID      uuid.UUID `gorm:"column:ayah_id;type:uuid;not null;index" json:"ayah_id"`
	AyahNumber  int       `gorm:"column:ayah_number;type:int;not null" json:"ayah_number"`
	ProgressPct float64   `gorm:"column:progress_pct;type:decimal(5,2);default:0" json:"progress_pct"`
	LastReadAt  time.Time `gorm:"column:last_read_at;autoUpdateTime" json:"last_read_at"`

	CreatedBy string         `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Surah SurahModel `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
	Ayah  AyahModel  `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
}

func (LastReadModel) TableName() string {
	return "last_reads"
}

type ProgressHatamModel struct {
	ID          uuid.UUID  `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      string     `gorm:"column:user_id;type:varchar(255);not null;index;" json:"user_id"`
	JuzID       uuid.UUID  `gorm:"column:juz_id;type:uuid;not null;index:idx_user_juz;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"juz_id"`
	StartAyahID uuid.UUID  `gorm:"column:start_ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"start_ayah_id"`
	LastAyahID  uuid.UUID  `gorm:"column:last_ayah_id;type:uuid;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"last_ayah_id"`
	ProgressPct float64    `gorm:"column:progress_pct;type:decimal(5,2);default:0" json:"progress_pct"`
	IsCompleted bool       `gorm:"column:is_completed;default:false" json:"is_completed"`
	StartedAt   time.Time  `gorm:"column:started_at;type:timestamp with time zone" json:"started_at"`
	CompletedAt *time.Time `gorm:"column:completed_at;type:timestamp with time zone" json:"completed_at,omitempty"`

	CreatedBy string         `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Juz       JuzModel   `gorm:"foreignKey:juz_id;references:id" json:"juz,omitempty"`
	StartAyah AyahModel  `gorm:"foreignKey:start_ayah_id;references:id" json:"start_ayah,omitempty"`
	LastAyah  *AyahModel `gorm:"foreignKey:last_ayah_id;references:id" json:"last_ayah,omitempty"`
}

func (ProgressHatamModel) TableName() string {
	return "progress_hatam"
}

// type TopicModel struct {
// 	ID          uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	Name        string         `gorm:"column:name;not null;unique;size:255"`
// 	Description string         `gorm:"column:description;type:text"`
// 	CreatedBy   string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy   string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
// }

// func (TopicModel) TableName() string {
// 	return "topics"
// }

// type TopicAyahModel struct {
// 	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	TopicID   uuid.UUID      `gorm:"primaryKey;column:topic_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"topic_id"`
// 	SurahID   uuid.UUID      `gorm:"primaryKey;column:surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"surah_id"`
// 	AyahID    uuid.UUID      `gorm:"primaryKey;column:ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ayah_id"`
// 	CreatedBy string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

// 	Topic TopicModel `gorm:"foreignKey:topic_id;references:id" json:"topic,omitempty"`
// 	Surah SurahModel `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
// 	Ayah  AyahModel  `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
// }

// func (TopicAyahModel) TableName() string {
// 	return "topic_ayahs"
// }

// type TajweedRuleModel struct {
// 	ID          uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	Rule        string         `gorm:"column:rule;not null;unique;size:255" json:"rule"`
// 	Explanation string         `gorm:"column:explanation;type:text" json:"explanation"`
// 	Color       string         `gorm:"column:color;size:50" json:"color"`
// 	CreatedBy   string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy   string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
// }

// func (TajweedRuleModel) TableName() string {
// 	return "tajweed_rules"
// }

// type AyahTajweedModel struct {
// 	ID            uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	TajweedRuleID uuid.UUID      `gorm:"primaryKey;column:tajweed_rule_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tajweed_rule_id"`
// 	SurahID       uuid.UUID      `gorm:"primaryKey;column:surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"surah_id"`
// 	AyahID        uuid.UUID      `gorm:"primaryKey;column:ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ayah_id"`
// 	Start         int            `gorm:"column:start;not null" json:"start"`
// 	End           int            `gorm:"column:end;not null" json:"end"`
// 	CreatedBy     string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy     string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

// 	TajweedRule TajweedRuleModel `gorm:"foreignKey:tajweed_rule_id;references:id" json:"tajweed_rule,omitempty"`
// 	Surah       SurahModel       `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
// 	Ayah        AyahModel        `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
// }

// func (AyahTajweedModel) TableName() string {
// 	return "ayah_tajweed"
// }

// type TafsirEditionModel struct {
// 	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	Name      string         `gorm:"column:name;not null;size:255"`
// 	Author    string         `gorm:"column:author;not null;size:255"`
// 	Language  string         `gorm:"column:language;not null;size:50"`
// 	CreatedBy string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
// }

// func (TafsirEditionModel) TableName() string {
// 	return "tafsir_editions"
// }

// type AyahTafsirModel struct {
// 	ID              uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	TafsirEditionID uuid.UUID      `gorm:"primaryKey;column:tafsir_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tafsir_id"`
// 	SurahID         uuid.UUID      `gorm:"primaryKey;column:surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"surah_id"`
// 	AyahID          uuid.UUID      `gorm:"primaryKey;column:ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ayah_id"`
// 	Text            string         `gorm:"column:text;not null;type:text"`
// 	CreatedBy       string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy       string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

// 	TafsirEdition TafsirEditionModel `gorm:"foreignKey:tafsir_edition_id;references:id" json:"tafsir_edition,omitempty"`
// 	Surah         SurahModel         `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
// 	Ayah          AyahModel          `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
// }

// func (AyahTafsirModel) TableName() string {
// 	return "ayah_tafsirs"
// }

// type StoryModel struct {
// 	ID         uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	Title      string         `gorm:"column:title;not null;size:255"`
// 	Summary    string         `gorm:"column:summary;type:text"`
// 	Characters datatypes.JSON `gorm:"column:characters;type:jsonb"`
// 	CreatedBy  string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy  string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt  time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt  time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
// }

// func (StoryModel) TableName() string {
// 	return "stories"
// }

// type StoryAyahModel struct {
// 	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
// 	StoryID   uuid.UUID      `gorm:"primaryKey;column:story_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"story_id"`
// 	SurahID   uuid.UUID      `gorm:"primaryKey;column:surah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"surah_id"`
// 	AyahID    uuid.UUID      `gorm:"primaryKey;column:ayah_id;type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ayah_id"`
// 	CreatedBy string         `gorm:"column:create_by;type:varchar(255);not null" json:"created_by"`
// 	UpdatedBy string         `gorm:"column:update_by;type:varchar(255);not null" json:"updated_by"`
// 	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
// 	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

// 	Story StoryModel `gorm:"foreignKey:story_id;references:id" json:"story,omitempty"`
// 	Surah SurahModel `gorm:"foreignKey:surah_id;references:id" json:"surah,omitempty"`
// 	Ayah  AyahModel  `gorm:"foreignKey:ayah_id;references:id" json:"ayah,omitempty"`
// }

// func (StoryAyahModel) TableName() string {
// 	return "story_ayahs"
// }
