package model

import "github.com/jinzhu/gorm"

// User is user model struct
type User struct {
	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100)"`
}

func MigrateUser(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON’T change existing column’s type or delete unused columns to protect data.
	tx.AutoMigrate(&User{})

	// TODO: modify column from HERE

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	return tx.Commit().Error
}
