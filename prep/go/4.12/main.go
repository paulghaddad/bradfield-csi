/*
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a
request to https://xkcd.com/571/info.0.json produces a detailed description of
comic 571, one of many favorites. Download each URL (once!) and build an offline index.
Write a tool xkcd that, using this index, prints the URL and transcript of each comic
that matches a search term provided on the command line.
*/

package main

import (
	"fmt"
	"os"

	"github.com/bradfield-csi/prep/go/4.12/index"
	"github.com/bradfield-csi/prep/go/4.12/search"
)

const maxID int = 2427
const dbFilename = "./data.json"
const indexFilename = "./index.json"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You must provide an argument. [build / search term]")
		return
	}

	command := args[0]
	if command == "build" {
		fmt.Println("Rebuilding Index")
		index.Build(dbFilename, indexFilename, maxID)
	} else if command == "search" {
		fmt.Println("Searching")

		searchTerm := os.Args[2]
		if len(searchTerm) == 0 {
			fmt.Println("You must provide a term to search for [search term]")
		}

		searchResults := search.Find(dbFilename, indexFilename, searchTerm)
		for i, result := range searchResults {
			fmt.Printf("\n\nResult %d:\n", i+1)
			fmt.Printf("Title: %s\n", result.Title)
			fmt.Printf("URL: %s\n", result.URL)
			fmt.Printf("Transcript: %s\n", result.Transcript)
		}
	} else {
		fmt.Println("Unknown argument. Must be [build / search term]")
	}
}
