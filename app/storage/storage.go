package storage

import (
	"notebook/app/formatters"
	"notebook/app/models"

	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Note{})
	return db
}

func (u *DBRepository) Create(note *models.Note) {
	u.db.Create(&note)
}

func List(db *gorm.DB) []models.Note {
	list, _ := db.Find(&[]models.Note{}).Rows()
	defer list.Close()

	return formatters.FromDBToNotesList(db, list)
}

func GetById(id int, db *gorm.DB) []models.Note {
	list, _ := db.Find(&[]models.Note{}, []int{id}).Rows()
	defer list.Close()

	return formatters.FromDBToNotesList(db, list)
}
