package httputils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient httpClient
	baseURI    url.URL
}

func NewClient(baseURI string, timeout int) (*Client, error) {
	parsedBaseURI, err := url.ParseRequestURI(baseURI)
	if err != nil {
		return nil, fmt.Errorf("%w; invalid base uri", err)
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	return &Client{
		httpClient: client,
		baseURI: url.URL{
			Scheme: parsedBaseURI.Scheme,
			Host:   parsedBaseURI.Host,
		},
	}, nil
}

// Get will make the request to an endpoint. The endpoint can be a path or a full url
func (c Client) Get(endpoint string) ([]byte, error) {
	var requestURL *url.URL
	var err error

	if strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
		requestURL, err = url.Parse(endpoint)
		if err != nil {
			return nil, err
		}
	} else {
		requestURL = &url.URL{
			Scheme: c.baseURI.Scheme,
			Host:   c.baseURI.Host,
		}

		if strings.Contains(endpoint, "?") {
			requestURL.Opaque = "//" + c.baseURI.Host + endpoint
		} else {
			requestURL.Path = endpoint
		}
	}

	request, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body; %w", err)
	}

	switch response.StatusCode {
	case http.StatusOK:
		return respBody, nil
	default:
		return nil, fmt.Errorf("response code received: %d", response.StatusCode)
	}
}
