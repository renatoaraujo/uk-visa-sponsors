package scraper

import (
	"fmt"
	"regexp"
)

type HttpUtils interface {
	Get(resourcePath string) ([]byte, error)
}

type WebScraper struct {
	http HttpUtils
}

func NewWebScraper(http HttpUtils) WebScraper {
	return WebScraper{
		http: http,
	}
}

func (s WebScraper) FindDataSourceURL(fileType, path string) (string, error) {
	resp, err := s.http.Get(path)
	if err != nil {
		return "", fmt.Errorf("error getting content from %s: %v", path, err)
	}

	pattern := fmt.Sprintf(`https?://[^"'\s]+\.%s`, fileType)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("error compiling regex: %v", err)
	}

	match := re.FindString(string(resp))
	if match == "" {
		return "", fmt.Errorf("no %s found in %s; %w", fileType, path, err)
	}

	return match, nil
}

func (s WebScraper) GetContent(path string) ([]byte, error) {
	return s.http.Get(path)
}
