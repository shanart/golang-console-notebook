package handlers

import (
	"fmt"
	"log"
	"notebook/app/storage"
	"reflect"

	"gorm.io/gorm"
)

func List(db *gorm.DB) {
	list, _ := db.Find(&[]storage.Note{}).Rows()
	defer list.Close()

	var note = storage.Note{}
	for list.Next() {
		db.ScanRows(list, &note)
		v := reflect.ValueOf(note)
		k := v.Type()

		// fmt.Println(note)
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("%s:\t%v\n", k.Field(i).Name, v.Field(i).Interface())
		}
	}
}

func Create(db *gorm.DB) {
	new_note := storage.Note{Content: "This is test message"}
	db.Create(&new_note)
}

func Search() {
	fmt.Println("Search Handler")
	// var search string
	// TODO: scan
}

func Delete() {
	log.Println("Delete Handler")
	// var id int
	// TODO: scan
}
