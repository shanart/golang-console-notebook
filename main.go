package main

import (
	"flag"
	"fmt"
)

func main() {

	flagList := flag.Bool("l", false, "List notes")

	flag.Parse()

	if *flagList {
		fmt.Println("List notes")
	} else {
		fmt.Println("Create new note mode")
	}

}
