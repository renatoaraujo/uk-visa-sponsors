package sponsors_test

import (
	"errors"
	"testing"

	"renatoaraujo/uk-visa-sponsors/internal/sponsors"
	mocks "renatoaraujo/uk-visa-sponsors/internal/sponsors/mocks"

	"github.com/stretchr/testify/require"
)

func TestHandler_Load_New(t *testing.T) {
	tests := []struct {
		name           string
		datasource     string
		scraperSetup   func(*mocks.Scraper)
		processorSetup func(*mocks.Processor)
		expectError    bool
	}{
		{
			name:       "Successfully load data using default datasource",
			datasource: "",
			scraperSetup: func(s *mocks.Scraper) {
				s.On("FindDataSourceURL", "csv", "/government/publications/register-of-licensed-sponsors-workers").Return("default_url", nil)
				s.On("GetContent", "default_url").Return([]byte("data"), nil)
			},
			processorSetup: func(p *mocks.Processor) {
				p.On("ProcessRawData", []byte("data"), &[]sponsors.Organisation{}).Return(nil)
			},
		},
		{
			name:       "Successfully load data using custom datasource",
			datasource: "custom_url",
			scraperSetup: func(s *mocks.Scraper) {
				s.On("GetContent", "custom_url").Return([]byte("data_custom"), nil)
			},
			processorSetup: func(p *mocks.Processor) {
				p.On("ProcessRawData", []byte("data_custom"), &[]sponsors.Organisation{}).Return(nil)
			},
		},
		{
			name:       "Error during data fetching",
			datasource: "",
			scraperSetup: func(s *mocks.Scraper) {
				s.On("FindDataSourceURL", "csv", "/government/publications/register-of-licensed-sponsors-workers").Return("", errors.New("fetching error"))
			},
			expectError: true,
		},
		{
			name:       "Error during data processing",
			datasource: "custom_url",
			scraperSetup: func(s *mocks.Scraper) {
				s.On("GetContent", "custom_url").Return([]byte("data_error"), nil)
			},
			processorSetup: func(p *mocks.Processor) {
				p.On("ProcessRawData", []byte("data_error"), &[]sponsors.Organisation{}).Return(errors.New("processing error"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraperMock := mocks.NewScraper(t)
			processorMock := mocks.NewProcessor(t)

			if tt.scraperSetup != nil {
				tt.scraperSetup(scraperMock)
			}
			if tt.processorSetup != nil {
				tt.processorSetup(processorMock)
			}

			handler := sponsors.NewHandler(scraperMock, processorMock)
			err := handler.Load(tt.datasource)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
