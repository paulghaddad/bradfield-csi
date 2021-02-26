/*
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a
request to https://xkcd.com/571/info.0.json produces a detailed description of
comic 571, one of many favorites. Download each URL (once!) and build an offline index.
Write a tool xkcd that, using this index, prints the URL and transcript of each comic
that matches a search term provided on the command line.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const maxID int = 2427
const dataFile = "./data.json"
const indexFile = "./index.json"

// XKCDApiRecord holds the data for each comic in the search database
type XKCDApiRecord struct {
	Num        int
	URL        string `json:"link"`
	Title      string
	Transcript string
}

// SearchResult are the attributes returned for search results
type SearchResult struct {
	URL        string
	Title      string
	Transcript string
}

// BuildIndex retrieves data from the XKCD site to build a local database of
// records
func BuildIndex() {
	// Store comic id (num) to its values
	apiRecords := make(map[int]map[string]string)

	// map of search terms to matching record ids for search index
	termsToIDs := make(map[string][]int)

	for i := 1; i <= maxID; i++ {
		fmt.Printf("Retrieving API response for record with ID: %d\n", i)

		url := fmt.Sprintf("%s%d%s", "https://xkcd.com/", i, "/info.0.json")
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		// ID 404 always returns a 404
		if resp.StatusCode != 200 {
			continue
		}

		defer resp.Body.Close()

		// Decode API response
		var result XKCDApiRecord
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			panic(err)
		}

		// Populate data records in memory
		values := make(map[string]string)
		values["URL"] = result.URL
		values["TItle"] = result.Title
		values["Transcript"] = result.Transcript
		id := int(result.Num)

		apiRecords[id] = values

		// Build search index based on individual words in the title
		wordRegex := regexp.MustCompile("[a-zA-Z]+")
		terms := wordRegex.FindAllString(result.Title, -1)

		for _, term := range terms {
			normalizedTerm := strings.ToLower(term)
			_, pres := termsToIDs[normalizedTerm]

			if pres {
				termsToIDs[normalizedTerm] = append(termsToIDs[normalizedTerm], id)
			} else {
				termsToIDs[normalizedTerm] = []int{id}
			}
		}
	}

	dataRecords, _ := json.Marshal(apiRecords)

	// Write database to file
	dataFile, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}

	recordsWriter := bufio.NewWriter(dataFile)

	_, err = recordsWriter.WriteString(string(dataRecords))
	if err != nil {
		panic(err)
	}

	recordsWriter.Flush()

	// Write Search Index to file
	indexToStore, _ := json.Marshal(termsToIDs)

	indexFile, err := os.Create(indexFile)
	if err != nil {
		panic(err)
	}

	indexWriter := bufio.NewWriter(indexFile)

	_, err = indexWriter.WriteString(string(indexToStore))
	if err != nil {
		panic(err)
	}

	indexWriter.Flush()
}

// SearchIndex returns the URL and transcript of each matching XKCD record
func SearchIndex(term string) []SearchResult {

	// Bring index into memory
	indexData, err := ioutil.ReadFile(indexFile)
	if err != nil {
		panic(err)
	}

	indexRecords := make(map[string][]int, 0)
	if err := json.Unmarshal(indexData, &indexRecords); err != nil {
		panic(err)
	}

	// Search index for the ids that match the search term
	matches, pres := indexRecords[term]
	if !pres {
		fmt.Println("No records found.")
		return []SearchResult{}
	}

	// Bring database into memory
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		panic(err)
	}

	var dataRecords = map[string]SearchResult{}
	if err := json.Unmarshal(data, &dataRecords); err != nil {
		panic(err)
	}

	// Return database records that match the search term
	matchData := []SearchResult{}
	for _, id := range matches {
		record := dataRecords[strconv.Itoa(id)]
		matchData = append(matchData, record)
	}

	return matchData
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You must provide an argument. [build / search term]")
		return
	}

	command := args[0]
	if command == "build" {
		fmt.Println("Rebuilding Index")
		BuildIndex()
	} else if command == "search" {
		fmt.Println("Searching")

		searchTerm := os.Args[2]
		if len(searchTerm) == 0 {
			fmt.Println("You must provide a term to search for [search term]")
		}

		searchResults := SearchIndex(searchTerm)
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
