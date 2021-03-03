package index

import (
	"bufio"
	"encoding/json"
	"fmt"
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
}

// Build retrieves data from the XKCD site to build a local database of records
func Build(dbFilename, indexFilename string, maxID int) {
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
	dataFile, err := os.Create(dbFilename)
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

	indexFile, err := os.Create(indexFilename)
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
