package main

import (
	"bufio"
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
// Store each json object to a file
// Create a search function that matches a search term
// Create a CLI to search against the index

const maxID int = 2 // 2427

func main() {
	for i := 1; i <= maxID; i++ {
		resp, err := http.Get("https://xkcd.com/571/info.0.json")
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)

		fmt.Println(resp.Status)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
