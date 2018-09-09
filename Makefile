.PHONY: all test benchmark
.DEFAULT_GOAL: test

all: test benchmark

test: 
	@echo "\n=== Running Tests... ===\n"
	go test -v ./...

benchmark:
	@echo "\n=== Running Benchmarks... ===\n"
	go test ./... -bench=. -benchmem