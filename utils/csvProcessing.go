package utils

import (
	"encoding/csv"
	"io"
)

func ReadCSV(file io.Reader) (error, [][]string) {
	csvContent, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err, [][]string{}
	}

	return nil, csvContent
}
