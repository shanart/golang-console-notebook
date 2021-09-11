package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}

	for {
		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 { /* TODO: this is very primitive method. */
			lines = append(lines, text)
		} else {
			break
		}
	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	// TODO: if path exist
	// TODO: generate file name
	f, err := os.Create("storage/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Write to file
	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
