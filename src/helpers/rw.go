package helpers

import (
	"os"
	"strings"
	"sync"
)

var (
	mutex sync.RWMutex
)

// ReadFile -> read file from fs
// file: "/file/path"
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

// ReadFiles -> read files from fs
// example
// files: "path1", "path2"
func ReadFiles(files ...string) []string {
	fileContents := make([]string, len(files))

	mutex.RLock()
	defer mutex.RUnlock()

	for _, file := range files {
		fileContents = append(fileContents, ReadFile(file))
	}

	return fileContents
}

// WriteToFile -> write to file in fs
// path: "/path/to/file"
// content: "fileContents"
func WriteToFile(path string, content string) {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
