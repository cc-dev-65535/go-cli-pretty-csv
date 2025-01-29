package lib

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvFile struct {
	filePath string
	content  [][]string
}

func (c *CsvFile) readCsvFile() ([][]string, error) {
	file, err := os.Open(c.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: %w", err)
	}
	return records, nil
}

func (c *CsvFile) printCsvFile() {
	for _, record := range c.content {
		for _, column := range record {
			fmt.Printf("%s ", column)
		}
		fmt.Println()
	}
}

func NewCsvFile(path string) (*CsvFile, error) {
	csvFile := &CsvFile{
		filePath: path,
		content:  nil,
	}
	fmt.Println(csvFile)
	records, err := csvFile.readCsvFile()
	if err != nil {
		return nil, err
	}
	// fmt.Println(records)
	csvFile.content = records
	return csvFile, nil
}
