package data

import (
	"encoding/csv"
	"strings"
)

type DataProcessor interface {
	ProcessCSVData(data []byte) ([]map[string]string, error)
}

type CSVProcessor struct{}

func NewCSVProcessor() DataProcessor {
	return &CSVProcessor{}
}

func (cp *CSVProcessor) ProcessCSVData(data []byte) ([]map[string]string, error) {
	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var processedData []map[string]string
	for _, record := range records[1:] {
		entry := map[string]string{
			"OrganisationName": record[0],
			"TownCity":         record[1],
			"County":           record[2],
			"TypeAndRating":    record[3],
			"Route":            record[4],
		}
		processedData = append(processedData, entry)
	}

	return processedData, nil
}
