/*
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a
request to https://xkcd.com/571/info.0.json produces a detailed description of
comic 571, one of many favorites. Download each URL (once!) and build an offline index.
Write a tool xkcd that, using this index, prints the URL and transcript of each comic
that matches a search term provided on the command line.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bradfield-csi/prep/go/4.12/index"
	"github.com/bradfield-csi/prep/go/4.12/search"
)

const maxID int = 2427
const dbFilename = "./data.json"
const indexFilename = "./index.json"

var n = flag.Int("n", maxID, "Max number of IDs to fetch from XKCD API")

func printSearchResult(searchResult search.Result) {
	fmt.Println("\nResult:")
	fmt.Printf("Title: %s\n", searchResult.Title)
	fmt.Printf("URL: %s\n", searchResult.URL)
	fmt.Printf("Transcript: %s\n", searchResult.Transcript)
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("You must provide an argument. [build / search term]")
		os.Exit(1)
	}

	command := args[0]

	if command == "build" {
		log.Println("Rebuilding Index")
		err := index.Build(dbFilename, indexFilename, *n)

		if err != nil {
			fmt.Println("There was an error encountered while building the index")
			os.Exit(1)
		}

	} else if command == "search" {

		// Ensure a search term is provided
		if len(args) < 2 {
			fmt.Println("You must provide a term to search for [search term]")
			os.Exit(1)
		}

		fmt.Println("Searching")
		searchTerm := args[1]
		searchResults, err := search.Find(dbFilename, indexFilename, searchTerm)

		if err != nil {
			fmt.Println("There was an error encountered while searching")
			os.Exit(1)
		}

		for _, result := range searchResults {
			printSearchResult(result)
		}
	} else {
		fmt.Println("Unknown argument. Must be [build / search term]")
		os.Exit(1)
	}
}
