package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvFile struct {
	filePath string
	content  [][]string
}

func (c *CsvFile) ReadCsvFile() ([][]string, error) {
	file, err := os.Open(c.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: %w", err)
	}
	return records, nil
}
