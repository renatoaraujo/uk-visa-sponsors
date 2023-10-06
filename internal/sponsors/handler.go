package sponsors

import (
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
			org := Organisation{
				Name:     entry["Organisation Name"],
				VisaType: entry["Route"],
			}
			h.Organisations.list = append(h.Organisations.list, org)
		}
	}

	return err
}
