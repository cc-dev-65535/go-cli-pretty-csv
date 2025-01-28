package cmd

import "fmt"

func PrettyExec(args []string) error {
	// create a new CsvFile struct
	csvFile := &CsvFile{
		filePath: args[0],
		content:  nil,
	}
	fmt.Println(csvFile)
	records, err := csvFile.ReadCsvFile()
	if err != nil {
		return err
	}
	fmt.Println(records)
	csvFile.content = records

	// TODO: implement the pretty print logic here

	// print the contents of the csv file
	for _, record := range csvFile.content {
		for _, field := range record {
			fmt.Printf("%s ", field)
		}
		fmt.Println()
	}

	return nil
}
