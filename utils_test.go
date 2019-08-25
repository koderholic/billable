package main

import (
	"billable/utils"
	"os"
	"testing"
)

func TestUtilsCsvProcessingAcceptsValidCSVAndReturnsCsvContentCorrectly(t *testing.T) {

	file, err := os.Open("test/validRequest.csv")
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	err, content := utils.ReadCSV(file)
	if err != nil {
		t.Errorf("Expected request to run without error. Got %s\n", err)
	}
	if len(content) <= 0 {
		t.Errorf("Expected length of content read to be greater than %d. Got %d\n", 0, len(content))
	}

}
