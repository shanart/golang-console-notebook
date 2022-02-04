package models

import "time"

type Note struct {
	ID        uint `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
