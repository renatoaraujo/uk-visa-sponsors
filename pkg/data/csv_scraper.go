package data

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type CSVScraper struct {
	url string
}

func NewCSVScraper(url string) *CSVScraper {
	return &CSVScraper{
		url: url,
	}
}

func (s *CSVScraper) FetchDataSource() (string, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return "", fmt.Errorf("error getting content from %s: %v", s.url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	pattern := `https?://[^"'\s]+\.csv`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("error compiling regex: %v", err)
	}

	match := re.FindString(string(body))
	if match == "" {
		return "", errors.New("no csv link found")
	}

	return match, nil
}

func (s *CSVScraper) FetchData(url string) ([]byte, error) {
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
