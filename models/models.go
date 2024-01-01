package models

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username  string    `gorm:"size:255;not null;unique" json:"Username"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null;" json:"password"`
	Photos    []Photo   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photos"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}


func (user *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	user.ID = uuid
	return nil
}

type Photo struct {
	ID        UUIDString `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title     string     `gorm:"size:255;not null;" json:"title"`
	Caption   string     `gorm:"not null;" json:"caption"`
	PhotoUrl  string     `gorm:"not null;" json:"photo_url"`
	UserID    UUIDString `gorm:"not null" json:"user_id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}


func (photo *Photo) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	photo.ID = UUIDString(uuid.String())
	return nil
}

type UUIDString string

func (u UUIDString) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *UUIDString) Scan(value interface{}) error {
	*u = UUIDString(value.(string))
	return nil
}
