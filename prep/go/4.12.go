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

/*
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a
request to https://xkcd.com/571/info.0.json produces a detailed description of
comic 571, one of many favorites. Download each URL (once!) and build an offline index.
Write a tool xkcd that, using this index, prints the URL and transcript of each comic
that matches a search term provided on the command line.
*/

const maxID int = 2427

// IKCDComicRecord holds the data for each comic in the search database
type IKCDComicRecord struct {
	Num        int
	URL        string `json:"link"`
	Title      string
	Transcript string
}

// DataAttributes are the attributes returned for search results
type DataAttributes struct {
	URL        string
	Title      string
	Transcript string
}

// BuildIndex retrieves data from the XKCD site to build a local database of
// records
func BuildIndex() {
	allDataRecords := make(map[int]map[string]string)
	termsToIDs := make(map[string][]int)

	for i := 1; i <= maxID; i++ {
		// TODO: there is a problem with ID 404, it returns a 404; need to skip records if
		// they don't return a 200
		if i == 404 {
			continue
		}
		fmt.Printf("Retrieving API response for record with ID: %d\n", i)
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

		// Populate database records
		values := make(map[string]string)
		values["URL"] = result.URL
		values["TItle"] = result.Title
		values["Transcript"] = result.Transcript

		id := int(result.Num)

		allDataRecords[id] = values

		// Build search index based on terms in the title
		wordRegex := regexp.MustCompile("[a-zA-Z]+")
		terms := wordRegex.FindAllString(result.Title, -1)

		for _, term := range terms {
			normTerm := strings.ToLower(term)
			_, pres := termsToIDs[normTerm]

			if pres {
				termsToIDs[normTerm] = append(termsToIDs[normTerm], id)
			} else {
				termsToIDs[normTerm] = []int{id}
			}
		}
	}

	recordsToStore, _ := json.Marshal(allDataRecords)

	// Write database to file
	recordsFile, err := os.Create("./records.json")
	if err != nil {
		panic(err)
	}

	recordsWriter := bufio.NewWriter(recordsFile)

	_, err = recordsWriter.WriteString(string(recordsToStore))
	if err != nil {
		panic(err)
	}

	recordsWriter.Flush()

	// Write Search Index to file
	indexToStore, _ := json.Marshal(termsToIDs)

	indexFile, err := os.Create("./index.json")
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

// MatchResult holds the two attributes displayed in search results
type MatchResult struct {
	url        string
	transcript string
}

// SearchIndex returns the URL and transcript of each matching XKCD record
func SearchIndex(term string) []DataAttributes {

	// Bring index into memory
	indexData, err := ioutil.ReadFile("./index.json")
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
		return []DataAttributes{}
	}

	// Bring datbase into memory
	data, err := ioutil.ReadFile("./records.json")
	if err != nil {
		panic(err)
	}

	var allDataRecords = map[string]DataAttributes{}
	if err := json.Unmarshal(data, &allDataRecords); err != nil {
		panic(err)
	}

	// Return database records that match the search term
	matchData := []DataAttributes{}
	for _, id := range matches {
		record := allDataRecords[strconv.Itoa(id)]
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
		searchResults := SearchIndex(os.Args[2])
		for i, result := range searchResults {
			fmt.Printf("\n\nResult %d:\n", i+1)
			fmt.Printf("Title: %s\n", result.Title)
			fmt.Printf("URL: %s\n", result.URL)
			fmt.Printf("Transcript: %s\n", result.Transcript)
		}
	} else {
		fmt.Println("Unknown argument")
	}
}
