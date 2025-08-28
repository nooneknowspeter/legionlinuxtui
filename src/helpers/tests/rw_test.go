package tests

import (
	"fmt"
	"legionlinuxtui/src/helpers"
	"testing"
)

var (
	testFile string = "test_file"
)

func TestRW(t *testing.T) {
	var expectedContents string = ""
	helpers.WriteToFile(testFile, expectedContents) // clear file

	var file string = helpers.ReadFile(testFile)

	fmt.Printf("File Contents: %v\n", file)
	fmt.Printf("Expected Contents: %v\n", expectedContents)
	if file != expectedContents {
		t.Errorf("  contents do not match what was expected in the file")
		return
	}

	expectedContents = "just another fucking test"
	helpers.WriteToFile(testFile, expectedContents)

	file = helpers.ReadFile(testFile)

	fmt.Printf("File Contents: %v\n", file)
	fmt.Printf("Expected Contents: %v\n", expectedContents)

	if file != expectedContents {
		t.Errorf("  contents do not match what was expected in the file")
		return
	}
}
