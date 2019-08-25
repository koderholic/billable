package utils

import (
	"encoding/csv"
	"strings"
)


func ReadCSV(csvInput string) (error, [][]string) {

	csvContent, err := csv.NewReader(strings.NewReader(csvInput)).ReadAll()
	if err != nil {
		return err, [][]string{}
	}

	return nil, csvContent
	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(record[1])
	// }
}
