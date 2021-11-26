package main

import (
	"flag"
)

var sep = flag.String("s", " ", "divider")

var f_get = flag.String("--get", "", "Get note")
var f_create = flag.String("--create", "", "Create new note")
var f_delete = flag.Int("--delete", 0, "Delete note by id")
var f_update = flag.String("--update", "", "Update note. Require using -t/--text flag to fill updated note content")
var f_text = flag.String("--text", "", "Content. Require using -u/--update flag to get updated note id")
var f_search = flag.String("--search", "", "Search note")
var f_init = flag.String("--init", "", "Init database")

func main() {

}
