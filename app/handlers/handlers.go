package handlers

import (
	"bufio"
	"fmt"
	"log"
	"notebook/app/config"
	"notebook/app/formatters"
	"notebook/app/models"
	"notebook/app/storage"
	"os"
	"strings"
	"time"
	"unsafe"

	"gorm.io/gorm"
)

func ListHandler(db *gorm.DB) {
	result := storage.List(db, 15)

	for _, item := range formatters.ListNotesItems(result) {
		fmt.Println(item)
	}
}

func CreateHandler(db *gorm.DB) {
	fmt.Print("Enter text:\n")
	scanner := bufio.NewScanner(os.Stdin)
	var noteText []string
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			noteText = append(noteText, text)
		} else {
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	if len(noteText) != 0 {
		note_content := strings.Join(noteText, "\n")
		new_note := models.Note{Content: note_content}
		db.Create(&new_note)
		fmt.Printf("\u001b[32mSaved. Size: %d| Date: %s\u001b[0m\n",
			unsafe.Sizeof(note_content),
			time.Now().Format(config.GetConfig().TimeFormat))
	}
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
