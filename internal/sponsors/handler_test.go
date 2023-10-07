package sponsors_test

import (
	"errors"
	"testing"

	"renatoaraujo/uk-visa-sponsors/internal/sponsors"
	mocks "renatoaraujo/uk-visa-sponsors/internal/sponsors/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	tests := []struct {
		name           string
		load           bool
		fetcherSetup   func(client *mocks.Fetcher)
		processorSetup func(client *mocks.Processor)
		wantErr        bool
	}{
		{
			name:    "Successfully initialise Handler without preloading data",
			load:    false,
			wantErr: false,
		},
		{
			name: "Successfully to initialise Handler preloading data",
			load: true,
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchDataSource").Return(
					"datasource",
					nil,
				)
				client.On("FetchData", mock.Anything).Return(
					[]byte("some primitive data"),
					nil,
				)
			},
			processorSetup: func(client *mocks.Processor) {
				client.On("ProcessRawData", mock.Anything).Return(
					[]map[string]string{
						{
							"Organisation Name": "Awesome company",
							"Route":             "Skilled Worker Visa",
						},
					},
					nil,
				)
			},
			wantErr: false,
		},
		{
			name: "Failed to initialise Handler preloading due to failure on fetching data",
			load: true,
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchDataSource").Return(
					"datasource",
					nil,
				)
				client.On("FetchData", mock.Anything).Return(
					nil,
					errors.New("failed to fetch data"),
				)
			},
			wantErr: true,
		},
		{
			name: "Failed to initialise Handler preloading due to failure on processing raw data",
			load: true,
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchDataSource").Return(
					"datasource",
					nil,
				)
				client.On("FetchData", mock.Anything).Return(
					[]byte("some primitive data"),
					nil,
				)
			},
			processorSetup: func(client *mocks.Processor) {
				client.On("ProcessRawData", mock.Anything).Return(
					nil,
					errors.New("failed to process data"),
				)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fetcherMock := &mocks.Fetcher{}
			if tt.fetcherSetup != nil {
				tt.fetcherSetup(fetcherMock)
			}

			processorMock := &mocks.Processor{}
			if tt.processorSetup != nil {
				tt.processorSetup(processorMock)
			}

			handler, err := sponsors.NewHandler(fetcherMock, processorMock, tt.load)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if handler != nil {
				require.IsType(t, &sponsors.Handler{}, handler)
			}

			if tt.load && !tt.wantErr {
				require.NotEmpty(t, handler.Organisations)
			}

			mock.AssertExpectationsForObjects(t, fetcherMock)
			mock.AssertExpectationsForObjects(t, processorMock)
		})
	}
}
