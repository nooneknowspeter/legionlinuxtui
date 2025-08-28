package helpers

import (
	"os"
	"strings"
	"sync"
)

var (
	mutex sync.RWMutex
)

func ReadFile(file string) string {
	mutex.RLock()
	defer mutex.RUnlock()

	contents, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	splittedString := strings.Split(string(contents), "\n") // remove newline escape sequences

	return splittedString[0]
}

func ReadFiles(files ...string) []string {
	fileContents := make([]string, len(files))

	mutex.RLock()
	defer mutex.RUnlock()

	for _, file := range files {
		fileContents = append(fileContents, ReadFile(file))
	}

	return fileContents
}
