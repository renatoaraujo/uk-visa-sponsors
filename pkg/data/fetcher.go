package data

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type DataFetcher interface {
	FetchDataSource() (string, error)
	FetchCSVData(url string) ([]byte, error)
}

type Scraper struct {
	url string
}

func NewDataFetcher(url string) DataFetcher {
	return &Scraper{
		url: url,
	}
}

func (s *Scraper) FetchDataSource() (string, error) {
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

func (s *Scraper) FetchCSVData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting content from %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}
