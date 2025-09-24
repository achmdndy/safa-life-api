package tajweed

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// TajweedWordModel represents a word and its tajweed rule for storage.
type TajweedWordModel struct {
	Word  string `json:"word"`
	Rule  string `json:"rule"`
	Color string `json:"color"`
}

// TajweedWords is a slice of TajweedWordModel to be stored as JSON.
type TajweedWords []TajweedWordModel

// Value implements the driver.Valuer interface for JSON serialization.
func (t TajweedWords) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// Scan implements the sql.Scanner interface for JSON deserialization.
func (t *TajweedWords) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, t)
}

// AyahTajweedModel represents the AyahTajweed table in the database.
type AyahTajweedModel struct {
	TajweedID string         `gorm:"primaryKey;column:tajweed_id;size:255" json:"tajweed_id"`
	SurahID   int            `gorm:"primaryKey;column:surah_id" json:"surah_id"`
	AyahID    int            `gorm:"primaryKey;column:ayah_id" json:"ayah_id"`
	Words     TajweedWords   `gorm:"column:words;type:jsonb" json:"words"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

// TableName specifies the table name for AyahTajweedModel.
func (AyahTajweedModel) TableName() string {
	return "ayah_tajweed"
}

// TajweedRuleModel represents the TajweedRule table in the database.
type TajweedRuleModel struct {
	ID          string         `gorm:"primaryKey;column:id;size:255" json:"id"`
	Rule        string         `gorm:"column:rule;not null;unique;size:255" json:"rule"`
	Explanation string         `gorm:"column:explanation;type:text" json:"explanation"`
	Color       string         `gorm:"column:color;size:50" json:"color"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

// TableName specifies the table name for TajweedRuleModel.
func (TajweedRuleModel) TableName() string {
	return "tajweed_rules"
}
