package handlers

import (
	"fmt"
	"log"
	"notebook/app/formatters"
	"notebook/app/models"
	"notebook/app/storage"

	"gorm.io/gorm"
)

func ListHandler(db *gorm.DB) {
	result := storage.List(db, 10)

	for _, item := range formatters.ListNotesItems(result) {
		fmt.Println(item)
	}
}

func CreateHandler(db *gorm.DB) {
	new_note := models.Note{Content: "This is test message"}
	db.Create(&new_note)
}

func SearchHandler() {
	fmt.Println("Search Handler")
	// var search string
	// TODO: scan
}

func DeleteHandler() {
	log.Println("Delete Handler")
	// var id int
	// TODO: scan
}
