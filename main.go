package main

import (
	"flag"
	"fmt"
)

func main() {
	f_get := flag.String("get", "", "Get note")
	f_create := flag.String("create", "", "Create new note")
	f_delete := flag.Int("delete", 0, "Delete note by id")
	f_update := flag.String("update", "", "Update note. Require using -t/--text flag to fill updated note content")
	f_text := flag.String("text", "", "Content. Require using -u/--update flag to get updated note id")
	f_search := flag.String("search", "", "Search note")
	f_init := flag.String("init", "", "Init database")

	flag.Parse()

	fmt.Printf("%s", *f_get)
	fmt.Printf("%s", *f_create)
	fmt.Printf("%d", *f_delete)
	fmt.Printf("%s", *f_update)
	fmt.Printf("%s", *f_text)
	fmt.Printf("%s", *f_search)
	fmt.Printf("%s", *f_init)

}
