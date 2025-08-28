.PHONY: all clean test tests

run:
	go run .
watch:
	air
build:
	go build -o build/
tests:
	go test ./... -v
