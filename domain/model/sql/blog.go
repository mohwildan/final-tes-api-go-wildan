package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"`
	Title     string    `gorm:"type:text;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UrlImage  string    `gorm:"type:text" json:"url_image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}