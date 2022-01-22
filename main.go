package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func logger(mode *bool, message string) {
	if *mode {
		log.Println(message)
	}
}

func saveNote() {
	var first string
	fmt.Scanln(&first)
}

// create note require

func main() {

	f_list := flag.String("l", "", "List")
	f_search := flag.String("s", "", "Search note")
	f_delete := flag.Int("d", 0, "Delete note by id")
	debug := flag.Bool("debug", false, "Debug mode")

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

	if len(*f_list) == 0 && len(*f_search) == 0 && *f_delete == 0 {
		// if all flags are empty - create new note.
		saveNote()
	}

	// TODO: if first argument is integer - read/edit this note by id

	fmt.Printf("%s", *f_list)
	fmt.Printf("%d", *f_delete)
	fmt.Printf("%s", *f_search)

}
