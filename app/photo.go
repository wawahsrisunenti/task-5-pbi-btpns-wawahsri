package app

import (
	"database/sql/driver"
	"time"
)

// CustomPhoto struct represents the data model for a photo with modified field names.
type CustomPhoto struct {
	ID        UUIDString `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"custom_id"`
	Title     string     `gorm:"size:255;not null;" json:"custom_title"`
	Caption   string     `gorm:"not null;" json:"custom_caption"`
	PhotoURL  string     `gorm:"not null;" json:"custom_photo_url"`
	UserID    UUIDString `gorm:"not null" json:"custom_user_id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"custom_created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"custom_updated_at"`
}

// UUIDString is a custom type for UUID representation as string.
type UUIDString string

func (u UUIDString) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *UUIDString) Scan(value interface{}) error {
	*u = UUIDString(value.(string))
	return nil
}
