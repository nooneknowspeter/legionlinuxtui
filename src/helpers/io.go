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

func WriteToFile(path string, content string) {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
