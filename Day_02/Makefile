.PHONY: all test run build benchmark

all: test run
test:
	go test
run-go:
	go run solution.go input.txt
build:
	go build solution.go
benchmark: build
	time -p ./solution

