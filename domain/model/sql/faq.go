package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)





type Faq struct {
	ID        uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"`
	Question  string         `gorm:"type:text;not null" json:"question"`
	Answer    string         `gorm:"type:text;not null" json:"answer"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}