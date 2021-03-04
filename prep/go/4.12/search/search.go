package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// Result are the attributes returned for search results
type Result struct {
	URL        string
	Title      string
	Transcript string
}

// Find returns the URL and transcript of each matching XKCD record
func Find(dbFilename, indexFilename, term string) ([]Result, error) {

	// Bring index into memory
	indexData, err := ioutil.ReadFile(indexFilename)
	if err != nil {
		log.Println("Error encountered reading index file")
		return nil, errors.New("Error reading index file")
	}

	indexRecords := make(map[string][]int, 0)
	if err := json.Unmarshal(indexData, &indexRecords); err != nil {
		log.Println("Error encountered unmarshaling index file")
		return nil, errors.New("Error unmarshaling index file")
	}

	// Search index for the ids that match the search term
	matches, pres := indexRecords[term]
	if !pres {
		fmt.Println("No records found.")
		return []Result{}, nil
	}

	// Bring database into memory
	data, err := ioutil.ReadFile(dbFilename)
	if err != nil {
		log.Println("Error encountered reading database file")
		return nil, errors.New("Error reading database file")
	}

	var dataRecords = map[string]Result{}
	if err := json.Unmarshal(data, &dataRecords); err != nil {
		log.Println("Error encountered unmarshaling database file")
		return nil, errors.New("Error unmarshaling database file")
	}

	// Return database records that match the search term
	matchData := []Result{}
	for _, id := range matches {
		record := dataRecords[strconv.Itoa(id)]
		matchData = append(matchData, record)
	}

	return matchData, nil
}
