package main

import (
	"flag"
	"fmt"
)

func main() {

	flagList := flag.Bool("l", false, "List notes")
	flagDelete := flag.Bool("d", false, "Delete note by id entered next")
	flagSearch := flag.Bool("s", false, "Search note by string entered next")
	flagInitStorage := flag.String("i", "", "Initialize notes storage")

	flag.Parse()

	if *flagList {
		fmt.Println("List notes")
	} else if *flagDelete {
		fmt.Println("Delete note", *flagDelete)
	} else if *flagSearch {
		fmt.Println("Search note by string", *flagSearch)
	} else if len(*flagInitStorage) > 0 {
		fmt.Printf("Init storage in: %s\n", *flagInitStorage)
	} else {
		fmt.Println("Create new note mode")
	}
}
