package pkg

import (
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

func (s *Scraper) findDataSource() {

}

func (s *Scraper) FetchData() error {
	resp, err := http.Get(s.url)
	if err != nil {
		return fmt.Errorf("error getting content from %s: %v", s.url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	content := string(body)

	// try to find any link that ends with .csv
	pattern := `https?://[^"'\s]+\.csv`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("error compiling regex: %v", err)
	}

	// return the first one or fail
	match := re.FindString(content)
	if match == "" {
		return errors.New("no csv link found")
	}

	fmt.Print(match)

	return nil
}
