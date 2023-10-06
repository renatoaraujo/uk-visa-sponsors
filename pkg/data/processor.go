package data

import (
	"encoding/csv"
	"strings"
)

type CSVProcessor struct{}

func NewCSVProcessor() *CSVProcessor {
	return &CSVProcessor{}
}

func (cp *CSVProcessor) ProcessRawData(data []byte) ([]map[string]string, error) {
	reader := csv.NewReader(strings.NewReader(string(data)))

	// Read the header
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Read the rest of the data
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var processedData []map[string]string
	for _, record := range records {
		entry := make(map[string]string)
		for i, head := range header {
			entry[head] = record[i]
		}
		processedData = append(processedData, entry)
	}

	return processedData, nil
}
