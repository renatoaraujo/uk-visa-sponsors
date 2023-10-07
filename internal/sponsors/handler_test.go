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
			fetcherMock := mocks.NewFetcher(t)
			if tt.fetcherSetup != nil {
				tt.fetcherSetup(fetcherMock)
			}

			processorMock := mocks.NewProcessor(t)
			if tt.processorSetup != nil {
				tt.processorSetup(processorMock)
			}

			handler, err := sponsors.NewHandler(fetcherMock, processorMock, tt.load)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.IsType(t, &sponsors.Handler{}, handler)

				if tt.load && !tt.wantErr {
					require.NotEmpty(t, handler.Organisations)
				}
			}
		})
	}
}

func TestHandler_Load(t *testing.T) {
	tests := []struct {
		name           string
		datasource     string
		fetcherSetup   func(client *mocks.Fetcher)
		processorSetup func(client *mocks.Processor)
		wantEmptyOrgs  bool
		wantErr        bool
	}{
		{
			name:       "Successfully load data using default datasource",
			datasource: "",
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchDataSource").Once().Return(
					"datasource",
					nil,
				)
				client.On("FetchData", mock.Anything).Once().Return(
					[]byte(""),
					nil,
				)
			},
			processorSetup: func(client *mocks.Processor) {
				client.On("ProcessRawData", mock.Anything).Once().Return(
					[]map[string]string{
						{
							"Organisation Name": "Awesome company",
							"Route":             "Skilled Worker",
						},
					},
					nil,
				)
			},
			wantErr: false,
		},
		{
			name:       "Successfully load data using custom datasource",
			datasource: "custom_datasource",
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchData", mock.Anything).Once().Return(
					[]byte(""),
					nil,
				)
			},
			processorSetup: func(client *mocks.Processor) {
				client.On("ProcessRawData", mock.Anything).Once().Return(
					[]map[string]string{
						{
							"Organisation Name": "Awesome company",
							"Route":             "Skilled Worker",
						},
					},
					nil,
				)
			},
			wantErr: false,
		},
		{
			name:       "Successfully load data using custom datasource, but it's empty",
			datasource: "custom_datasource",
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchData", mock.Anything).Once().Return(
					[]byte(""),
					nil,
				)
			},
			processorSetup: func(client *mocks.Processor) {
				client.On("ProcessRawData", mock.Anything).Once().Return(
					[]map[string]string{},
					nil,
				)
			},
			wantEmptyOrgs: true,
			wantErr:       false,
		},
		{
			name: "Failed to load due to failure to fetch datasource",
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchDataSource").Once().Return(
					"",
					errors.New("failed to fetch datasource"),
				)
			},
			wantErr: true,
		},
		{
			name:       "Failed to load due to failure to fetch data from datasource",
			datasource: "custom_datasource",
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchData", mock.Anything).Once().Return(
					[]byte(""),
					errors.New("failed to fetch datasource"),
				)
			},
			wantErr: true,
		},
		{
			name:       "Failed to process data due to incompatibility",
			datasource: "custom_datasource",
			fetcherSetup: func(client *mocks.Fetcher) {
				client.On("FetchData", mock.Anything).Once().Return(
					[]byte(""),
					nil,
				)
			},
			processorSetup: func(client *mocks.Processor) {
				client.On("ProcessRawData", mock.Anything).Once().Return(
					[]map[string]string{
						{
							"OrganisationName": "Awesome company",
							"Route":            "Skilled Worker",
						},
					},
					nil,
				)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fetcherMock := mocks.NewFetcher(t)
			if tt.fetcherSetup != nil {
				tt.fetcherSetup(fetcherMock)
			}

			processorMock := mocks.NewProcessor(t)
			if tt.processorSetup != nil {
				tt.processorSetup(processorMock)
			}

			handler, err := sponsors.NewHandler(fetcherMock, processorMock, false)
			require.NoError(t, err)

			err = handler.Load(tt.datasource)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				if tt.wantEmptyOrgs {
					require.Empty(t, handler.Organisations.List())
				} else {
					require.NotEmpty(t, handler.Organisations.List())
				}
			}
		})
	}
}
