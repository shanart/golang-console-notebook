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

func main() {

	f_get := flag.String("get", "", "Get note")
	f_create := flag.String("create", "", "Create new note")
	f_delete := flag.Int("delete", 0, "Delete note by id")
	f_update := flag.String("update", "", "Update note. Require using -t/--text flag to fill updated note content")
	f_text := flag.String("text", "", "Content. Require using -u/--update flag to get updated note id")
	f_search := flag.String("search", "", "Search note")
	f_init := flag.String("init", "", "Init database")
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

	fmt.Printf("%s", *f_get)
	fmt.Printf("%s", *f_create)
	fmt.Printf("%d", *f_delete)
	fmt.Printf("%s", *f_update)
	fmt.Printf("%s", *f_text)
	fmt.Printf("%s", *f_search)
	fmt.Printf("%s", *f_init)

}
