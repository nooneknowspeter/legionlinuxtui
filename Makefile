.PHONY: all clean test tests build lint treefmt

run:
	go run .
watch:
	air
build:
	go build -o build/
tests:
	go test ./src/... -v
lint:
	treefmt --ci --config-file treefmt.lint.toml
	golangci-lint run
format:
	treefmt
