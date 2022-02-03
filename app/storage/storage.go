package storage

import (
	"time"

	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

type Note struct {
	ID        uint `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Note{})
	return db
}

func (u *DBRepository) Create(note Note) {
	u.db.Create(&note)
}

func (u *DBRepository) List(limit int) {
	u.db.Model(&Note{}).Limit(limit)
}
