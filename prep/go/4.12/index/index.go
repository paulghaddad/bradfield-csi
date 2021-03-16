package index

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// XKCDApiRecord holds the data for each comic in the search database
type XKCDApiRecord struct {
	Num        int
	URL        string `json:"link"`
	Title      string
	Transcript string
	AltText    string `json:"alt"`
}

func fetchAPIRecord(i int) (map[string]string, error) {
	log.Printf("Retrieving API response for record with ID: %d\n", i)

	url := fmt.Sprintf("%s%d%s", "https://xkcd.com/", i, "/info.0.json")
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed HTTP GET request for URL: %s\n", url)
		return map[string]string{}, errors.New("Failed HTTP GET request")
	}

	// ID 404 always returns a 404
	if resp.StatusCode != http.StatusOK {
		log.Printf("Unsuccessful HTTP request for URL with status code %s: %d\n", url, resp.StatusCode)
		return map[string]string{}, errors.New("HTTP GET request with unsuccessful status code")
	}

	defer resp.Body.Close()

	// Decode API response
	var result XKCDApiRecord
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error encountered while decoding HTTP response for ID: %d", i)
		return map[string]string{}, errors.New("Failed decoding HTTP response")
	}

	values := make(map[string]string)
	values["URL"] = result.URL
	values["Title"] = result.Title

	if result.Transcript == "" {
		values["Transcript"] = result.AltText
	} else {
		values["Transcript"] = result.Transcript
	}

	return values, nil
}

func createDB(filename string, apiRecords map[int]map[string]string) error {
	dataRecords, err := json.Marshal(apiRecords)
	if err != nil {
		log.Println("Error encountered while marshalling API Records")
		return errors.New("Error creating database file")
	}

	dataFile, err := os.Create(filename)
	if err != nil {
		log.Println("Error encountered while creating Database file.")
		return errors.New("Error creating database file")
	}

	recordsWriter := bufio.NewWriter(dataFile)

	_, err = recordsWriter.WriteString(string(dataRecords))
	if err != nil {
		log.Println("Error encountered while writing records to database file")
		return errors.New("Error writing to database file")
	}

	recordsWriter.Flush()

	return nil
}

func createIndex(filename string, termsToIDs map[string][]int) error {
	indexToStore, _ := json.Marshal(termsToIDs)

	indexFile, err := os.Create(filename)
	if err != nil {
		log.Println("Error encountered while creating index file")
		return errors.New("Error creating index file")
	}

	indexWriter := bufio.NewWriter(indexFile)

	_, err = indexWriter.WriteString(string(indexToStore))
	if err != nil {
		log.Println("Error encountered while writing records to index file")
		return errors.New("Error writing to index file")
	}

	indexWriter.Flush()

	return nil
}

func mapTermsToResults(termsToIDs map[string][]int, id int, title string) {
	wordRegex := regexp.MustCompile("[a-zA-Z]+")
	terms := wordRegex.FindAllString(title, -1)

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

// Build retrieves data from the XKCD site to build a local database of records
func Build(dbFilename, indexFilename string, maxID int) error {
	// Store comic id (num) to its values
	apiRecords := make(map[int]map[string]string)

	// map of search terms to matching record ids for search index
	termsToIDs := make(map[string][]int)

	// Populate data records in memory
	for i := 1; i <= maxID; i++ {
		record, err := fetchAPIRecord(i)
		if err != nil {
			continue
		}

		apiRecords[int(i)] = record

		// Build search index based on individual words in the title
		mapTermsToResults(termsToIDs, int(i), record["Title"])
	}

	// Create Database File
	dbErr := createDB(dbFilename, apiRecords)
	if dbErr != nil {
		log.Println("Error encountered while creating database file")
		return dbErr
	}

	// Create Index File
	indErr := createIndex(indexFilename, termsToIDs)
	if indErr != nil {
		log.Println("Error encountered while creating index file")
		return indErr
	}

	return nil
}
