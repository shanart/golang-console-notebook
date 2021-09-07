package main

import (
	"bufio"
	"log"
	"os"
)

var lines = []string{
	"This is just for example",
	"New line",
}

func main() {
	// create file
	f, err := os.Create("./storage/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	// create new buffer
	buffer := bufio.NewWriter(f)

	for _, line := range lines {
		_, err := buffer.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// flush buffered data to the file
	if err := buffer.Flush(); err != nil {
		log.Fatal(err)
	}
}
