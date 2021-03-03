package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// Result are the attributes returned for search results
type Result struct {
	URL        string
	Title      string
	Transcript string
}

// Find returns the URL and transcript of each matching XKCD record
func Find(dbFilename, indexFilename, term string) []Result {

	// Bring index into memory
	indexData, err := ioutil.ReadFile(indexFilename)
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
		return []Result{}
	}

	// Bring database into memory
	data, err := ioutil.ReadFile(dbFilename)
	if err != nil {
		panic(err)
	}

	var dataRecords = map[string]Result{}
	if err := json.Unmarshal(data, &dataRecords); err != nil {
		panic(err)
	}

	// Return database records that match the search term
	matchData := []Result{}
	for _, id := range matches {
		record := dataRecords[strconv.Itoa(id)]
		matchData = append(matchData, record)
	}

	return matchData
}
