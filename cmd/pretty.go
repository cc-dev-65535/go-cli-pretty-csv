package cmd

import "fmt"

func PrettyExec() error {
	csvFile := &CsvFile{
		filePath: "aa",
		content:  nil,
	}
	fmt.Println(csvFile)
	records, err := csvFile.ReadCsvFile()
	if err != nil {
		return err
	}
	fmt.Println(records)
	return nil
}
