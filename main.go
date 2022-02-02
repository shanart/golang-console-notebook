package main

import (
	"flag"
	"fmt"

	"notebook/config"

	"gorm.io/gorm"
)

func main() {

	db, _ := gorm.Open("root:root@/golang")
	defer db.Close()

	flagList := flag.Bool("l", false, "List notes")
	flagDelete := flag.Bool("d", false, "Delete note by id entered next")
	flagSearch := flag.Bool("s", false, "Search note by string entered next")
	flagInitStorage := flag.Bool("i", false, "Initialize notes storage")

	flag.Parse()

	if *flagList {
		fmt.Println("List notes")
	} else if *flagDelete {
		fmt.Println("Delete note")
	} else if *flagSearch {
		fmt.Println("Search note by string")
	} else if *flagInitStorage {
		config.Configure()
	} else {
		fmt.Println("Create new note mode")
	}
}
