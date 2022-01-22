package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var f_list = flag.Bool("l", false, "List")
var f_search = flag.String("s", "", "Search note")
var f_delete = flag.Int("d", 0, "Delete note by id")
var debug = flag.Bool("debug", false, "Debug mode")

func logger(mode *bool, message string) {
	if *mode {
		log.Println(message)
	}
}

func saveNote() {
	var first string
	fmt.Scanln(&first)
	logger(debug, "Note saved")
}

func main() {

	flag.Parse()

	if *debug {
		f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		log.SetOutput(f)

		logger(debug, "Application startup")
	}

	logger(debug, "Read flags")

	if !*f_list && len(*f_search) == 0 && *f_delete == 0 {
		// if all flags are empty - create new note.
		saveNote()
	}

	// TODO: if first argument is integer - read/edit this note by id

	if *f_list {
		fmt.Println(*f_list)
	}

	// fmt.Println(*f_delete)
	// fmt.Println(*f_search)

}
