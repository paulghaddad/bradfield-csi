package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a
request to https://xkcd.com/571/info.0.json produces a detailed description of
comic 571, one of many favorites. Download each URL (once!) and build an offline index.
Write a tool xkcd that, using this index, prints the URL and transcript of each comic
that matches a search term provided on the command line.
*/

// Tasks:
// Store each json object to a json file
// Create a search function that matches a search term
// Create a CLI to search against the index

const maxID int = 2 // 2427

// IKCDComicRecord holds the data for each comic in the search database
type IKCDComicRecord struct {
	Num        int
	URL        string `json:"link"`
	Title      string
	Transcript string
}

func main() {
	for i := 1; i <= maxID; i++ {
		url := fmt.Sprintf("%s%d%s", "https://xkcd.com/", i, "/info.0.json")
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		var result IKCDComicRecord

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", result)
	}
}
