package app

import (
	"database/sql/driver"
	"time"
)

// Photo struct represents the data model for a photo.
type Photo struct {
	ID        UUIDString `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title     string     `gorm:"size:255;not null;" json:"title"`
	Caption   string     `gorm:"not null;" json:"caption"`
	PhotoUrl  string     `gorm:"not null;" json:"photo_url"`
	UserID    UUIDString `gorm:"not null" json:"user_id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UUIDString string

func (u UUIDString) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *UUIDString) Scan(value interface{}) error {
	*u = UUIDString(value.(string))
	return nil
}
