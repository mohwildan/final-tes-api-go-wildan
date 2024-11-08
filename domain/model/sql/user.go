package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)




type User struct {
    ID        uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"`
    Name      string         `gorm:"type:varchar(100);not null" json:"name"`
    Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
    Password  string         `gorm:"type:varchar(100);not null" json:"-"`
    Address   string         `gorm:"type:text" json:"address"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
//     user.ID = uuid.New()
//     return
// }