package models

import "time"

type Note struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	// TODO: add number of characters
	// TODO: size in bites of content
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
