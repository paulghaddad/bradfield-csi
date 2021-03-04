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

func readIndex(filename string) (map[string][]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error encountered reading index file")
		return nil, errors.New("Error reading index file")
	}

	records := make(map[string][]int)
	if err := json.Unmarshal(data, &records); err != nil {
		log.Println("Error encountered unmarshaling index file")
		return nil, errors.New("Error unmarshaling index file")
	}

	return records, nil
}

func readDB(filename string) (map[string]Result, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error encountered reading database file")
		return nil, errors.New("Error reading database file")
	}

	var records = map[string]Result{}
	if err := json.Unmarshal(data, &records); err != nil {
		log.Println("Error encountered unmarshaling database file")
		return nil, errors.New("Error unmarshaling database file")
	}

	return records, nil
}

// Find returns the URL and transcript of each matching XKCD record
func Find(dbFilename, indexFilename, term string) ([]Result, error) {

	// Bring index into memory
	indexRecords, err := readIndex(indexFilename)
	if err != nil {
		return nil, err
	}

	// Search index for the ids that match the search term
	indexMatches, pres := indexRecords[term]
	if !pres {
		fmt.Println("No records found.")
		return []Result{}, nil
	}

	// Bring database into memory
	dataRecords, err := readDB(dbFilename)
	if err != nil {
		return nil, err
	}

	// Return database records that match the search term
	matchData := []Result{}
	for _, id := range indexMatches {
		record := dataRecords[strconv.Itoa(id)]
		matchData = append(matchData, record)
	}

	return matchData, nil
}
