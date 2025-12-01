.PHONY: all run-part1 run-part2 test

all: run-part1 run-part2 run-go
run-part1:
	python solution.py input.txt

run-part2:
	python solution2.py input.txt

test: test-part1 test-part2

test-part1: test_input.txt
	@if [ "$$(python solution.py test_input.txt)" = "3" ]; then \
		echo "Part 1 test passed!"; \
	else \
		echo "Part 1 test failed!"; \
	fi

test-part2: test_input.txt
	@if [ "$$(python solution2.py test_input.txt)" = "6" ]; then \
		echo "Part 2 test passed!"; \
	else \
		echo "Part 2 test failed!"; \
	fi
run-go:
	go run solution.go input.txt