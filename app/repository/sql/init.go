package sql

import "gorm.io/gorm"


type repository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &repository{db: db}
}

func NewFaqRepository(db *gorm.DB) FaqRepository {
    return &repository{db: db}
}

