.PHONY: all clean test

run:
	go run .
watch:
	air
build:
	go build -o build/