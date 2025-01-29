package lib

import (
	"fmt"
	"strings"
)

func PrettyExec(args []string) error {
	csvFile, err := NewCsvFile(args[0])
	if err != nil {
		return err
	}

	// TODO: implement the pretty print logic here
	columns := len(csvFile.content[0])
	var maxLengths = make([]int, columns)
	for _, record := range csvFile.content {
		for i, column := range record {
			if len(column) > maxLengths[i] {
				maxLengths[i] = len(column)
			}
		}
	}

	for i, record := range csvFile.content {
		for j := range record {
			if len(csvFile.content[i][j]) < maxLengths[j] {
				csvFile.content[i][j] = fmt.Sprintf("%s%s", csvFile.content[i][j], strings.Repeat(" ", maxLengths[j]-len(csvFile.content[i][j])))
				csvFile.content[i][j] += "|"
			}
		}
	}

	csvFile.printCsvFile()

	return nil
}
