package app

import (
	"time"

	"github.com/google/uuid"
)

// CustomUser struct represents the data model for a user with modified field names.
type CustomUser struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"custom_id"`
	Username  string    `gorm:"size:255;not null;unique" json:"custom_username"`
	Email     string    `gorm:"not null;unique" json:"custom_email"`
	Password  string    `gorm:"not null;" json:"custom_password"`
	Photos    []CustomPhoto   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"custom_photos"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"custom_created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"custom_updated_at"`
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
