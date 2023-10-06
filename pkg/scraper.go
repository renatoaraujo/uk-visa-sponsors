package pkg

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type Scraper struct {
	url string
}

func NewScraper(url string) *Scraper {
	return &Scraper{
		url: url,
	}
}

func (s *Scraper) findDataSource() (string, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return "", fmt.Errorf("error getting content from %s: %v", s.url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	content := string(body)

	// try to find any link that ends with .csv
	pattern := `https?://[^"'\s]+\.csv`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("error compiling regex: %v", err)
	}

	// return the first one or fail
	match := re.FindString(content)
	if match == "" {
		return "", errors.New("no csv link found")
	}

	return match, nil
}

func (s *Scraper) FetchData() ([]map[string]interface{}, error) {
	ds, err := s.findDataSource()
	if err != nil {
		return nil, err
	}
	output := fmt.Sprintf("found datasource link: %s", ds)
	fmt.Println(output)

	resp, err := http.Get(ds)
	if err != nil {
		return nil, fmt.Errorf("error getting content from datasource %s: %v", ds, err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	var data []map[string]interface{}
	for _, record := range records[1:] {
		entry := map[string]interface{}{
			"OrganisationName": record[0],
			"TownCity":         record[1],
			"County":           record[2],
			"TypeAndRating":    record[3],
			"Route":            record[4],
		}
		data = append(data, entry)
	}

	return data, nil
}
