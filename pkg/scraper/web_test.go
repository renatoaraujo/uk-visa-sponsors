package scraper_test

import (
	"errors"
	"testing"

	"renatoaraujo/uk-visa-sponsors/pkg/scraper"
	mocks "renatoaraujo/uk-visa-sponsors/pkg/scraper/mocks"

	"github.com/stretchr/testify/require"
)

func TestWebScraper_FindDataSourceURL(t *testing.T) {
	tests := []struct {
		name          string
		fileType      string
		path          string
		HttpUtilsMock func(mock *mocks.HttpUtils)
		wantURL       string
		wantErr       bool
	}{
		{
			name:     "Successful extraction of a URL for a given file type",
			fileType: "csv",
			path:     "http://example.com",
			HttpUtilsMock: func(mock *mocks.HttpUtils) {
				mock.On("Get", "http://example.com").Once().Return(
					[]byte("data before https://valid-url.com/data.csv data after"),
					nil,
				)
			},
			wantURL: "https://valid-url.com/data.csv",
			wantErr: false,
		},
		{
			name:     "No URL found for the given file type",
			fileType: "csv",
			path:     "http://example.com",
			HttpUtilsMock: func(mock *mocks.HttpUtils) {
				mock.On("Get", "http://example.com").Once().Return(
					[]byte("data without any csv url"),
					nil,
				)
			},
			wantURL: "",
			wantErr: true,
		},
		{
			name:     "Error when fetching content from the provided path",
			fileType: "csv",
			path:     "http://example.com",
			HttpUtilsMock: func(mock *mocks.HttpUtils) {
				mock.On("Get", "http://example.com").Once().Return(
					nil,
					errors.New("failed to fetch content"),
				)
			},
			wantURL: "",
			wantErr: true,
		},
		{
			name:     "Error due to invalid regex pattern",
			fileType: "?",
			path:     "http://example.com",
			HttpUtilsMock: func(mock *mocks.HttpUtils) {
				mock.On("Get", "http://example.com").Once().Return(
					[]byte("data with invalid pattern"),
					nil,
				)
			},
			wantURL: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHttpUtils := mocks.NewHttpUtils(t)
			if tt.HttpUtilsMock != nil {
				tt.HttpUtilsMock(mockHttpUtils)
			}

			scraper := scraper.NewWebScraper(mockHttpUtils)
			url, err := scraper.FindDataSourceURL(tt.fileType, tt.path)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantURL, url)
			}
		})
	}
}

func TestWebScraper_GetContent(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		HttpUtilsMock func(mock *mocks.HttpUtils)
		wantContent   []byte
		wantErr       bool
	}{
		{
			name: "Successful retrieval of content",
			path: "http://example.com",
			HttpUtilsMock: func(mock *mocks.HttpUtils) {
				mock.On("Get", "http://example.com").Once().Return(
					[]byte("sample content"),
					nil,
				)
			},
			wantContent: []byte("sample content"),
			wantErr:     false,
		},
		{
			name: "Error when fetching content from the provided path",
			path: "http://example.com",
			HttpUtilsMock: func(mock *mocks.HttpUtils) {
				mock.On("Get", "http://example.com").Once().Return(
					nil,
					errors.New("failed to fetch content"),
				)
			},
			wantContent: nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHttpUtils := mocks.NewHttpUtils(t)
			if tt.HttpUtilsMock != nil {
				tt.HttpUtilsMock(mockHttpUtils)
			}

			s := scraper.NewWebScraper(mockHttpUtils)
			content, err := s.GetContent(tt.path)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantContent, content)
			}
		})
	}
}
