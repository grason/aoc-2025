#!/usr/bin/env sh

# Advent of Code — New Day Setup
# Works on: dash, bash, zsh, fish (as source), Alpine, Ubuntu, Arch, macOS, WSL, Git Bash

# ----------------------------------------------------------------------
# Use only POSIX features from here on — no bashisms, no pipefail
# ----------------------------------------------------------------------

# Get day number
if [ "$#" -eq 1 ] && expr "$1" : '^[0-9]\+$' > /dev/null && [ "$1" -ge 1 ] && [ "$1" -le 25 ]; then
    DAY="$1"
elif date '+%m' 2>/dev/null | grep -q '^12$'; then
    # December → use today
    DAY=$(date '+%d' 2>/dev/null | sed 's/^0*//')
    [ "$DAY" -gt 25 ] && DAY=25   # safety
else
    printf 'Day number (1-25): '
    read -r DAY || exit 1
fi

# Validate day (again, POSIX-safe)
case "$DAY" in
    ''|*[!0-9]*) echo "Invalid day: must be a number" >&2; exit 1 ;;
    *)           [ "$DAY" -ge 1 ] && [ "$DAY" -le 25 ] || { echo "Day must be 1–25" >&2; exit 1; } ;;
esac

DIR=$(printf 'Day_%02d' "$DAY")

if [ -e "$DIR" ]; then
    echo "$DIR already exists!" >&2
    exit 1
fi

echo "Creating $DIR ..."
mkdir -p "$DIR" && cd "$DIR"

# Empty input files
> input.txt
> test_input.txt

# Go solution (clean, minimal, useful)
cat > solution.go <<'EOF'
package main

import (
	"fmt"
	"os"
	"strings"
)

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	input := strings.TrimSpace(string(data))

	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
EOF

# Test file
cat > solution_test.go <<'EOF'
package main

import "testing"

func TestPart1(t *testing.T) {
	// Add tests here
	if got := Part1(""); got != 0 {
		t.Errorf("Part1() = %d; want 0", got)
	}
}

func TestPart2(t *testing.T) {
	// Add tests here
}
EOF

# Makefile — works with GNU make and BSD make
cat > Makefile <<'EOF'
.PHONY: all test run

all: test run

test:
	go test -v

run:
	go run solution.go input.txt

clean:
	rm -f solution
EOF

# Quick runner script
cat > run.sh <<'EOF'
#!/usr/bin/env sh
exec go run solution.go "${1:-input.txt}"
EOF
chmod +x run.sh

echo "Done! $DIR is ready."
echo "   • Run tests + solution: make"
echo "   • Quick run:           ./run.sh"
echo "   • With test input:     ./run.sh test_input.txt"