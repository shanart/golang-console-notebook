package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

type Note struct {
	gorm.Model
	Content string
}

func Initialize(storagePath string) {
	db, err := gorm.Open(sqlite.Open(storagePath), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Note{})
}

func (u *DBRepository) Create(note Note) error {
	u.db.Create(&note)
	return nil
}
