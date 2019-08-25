package utils

import (
	"os"
	"testing"
)

type Test struct {
	In  string
	Out int
}

var test = Test{
	In:  "../test/validRequest.csv",
	Out: 0,
}

func TestUtilsCsvProcessingAcceptsValidCSVAndReturnsCsvContentCorrectly(t *testing.T) {

	file, err := os.Open(test.In)
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	err, content := ReadCSV(file)
	if err != nil {
		t.Errorf("Expected request to run without error. Got %s\n", err)
	}
	if len(content) <= test.Out {
		t.Errorf("Expected length of content read to be greater than %d. Got %d\n", test.Out, len(content))
	}

}
