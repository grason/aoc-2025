.PHONY: all test test-v

all: test
test-v:
	go test ./... -v
test:
	go test ./...