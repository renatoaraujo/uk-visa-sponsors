package sponsors

import (
	"errors"
	"fmt"
)

type Fetcher interface {
	FetchDataSource() (string, error)
	FetchData(url string) ([]byte, error)
}

type Processor interface {
	ProcessRawData(data []byte) ([]map[string]string, error)
}

type Handler struct {
	Fetcher       Fetcher
	Processor     Processor
	Organisations Organisations
}

func NewHandler(f Fetcher, p Processor, load bool) (*Handler, error) {
	h := &Handler{
		Fetcher:   f,
		Processor: p,
	}

	if load {
		if err := h.Load(""); err != nil {
			return nil, fmt.Errorf("failed to load organisations; %w", err)
		}
	}

	return h, nil
}

func (h *Handler) Load(dataSource string) error {
	var err error

	if h.Organisations.list == nil {
		if dataSource == "" {
			dataSource, err = h.Fetcher.FetchDataSource()
			if err != nil {
				return err
			}
		}

		rawData, err := h.Fetcher.FetchData(dataSource)
		if err != nil {
			return err
		}

		processedData, err := h.Processor.ProcessRawData(rawData)
		if err != nil {
			return err
		}

		// TODO: configure the dynamic header mapper
		for _, entry := range processedData {
			if len(entry["Organisation Name"]) <= 0 || len(entry["Route"]) <= 0 {
				return errors.New("incompatible data, expecting headers `Organisation Name` and `Route`")
			}

			h.Organisations.AddOrUpdateVisaType(entry["Organisation Name"], entry["Route"])
		}
	}

	return err
}
