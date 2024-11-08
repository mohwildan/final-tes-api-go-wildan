package migrations

import (
	"app/domain/model/sql"

	"gorm.io/gorm"
)



func Migrate(db *gorm.DB) error {
    db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

    if err := db.Exec("ALTER TABLE users DROP COLUMN id;").Error; err != nil {
        return err
    }

    if err := db.Exec("ALTER TABLE users ADD COLUMN id uuid PRIMARY KEY DEFAULT uuid_generate_v4();").Error; err != nil {
        return err
    } 
    return db.AutoMigrate(
        &sql.User{},
        &sql.Faq{},
        &sql.Blog{},
    )
}
