#!/usr/bin/env bash
set -euo pipefail

# ---- Get day number -------------------------------------------------
if [[ $# -eq 1 ]] && [[ $1 =~ ^[0-9]{1,2}$ ]]; then
  DAY=$1
elif [[ $(date +%m) == "12" ]]; then
  DAY=$(date +%-d)
else
  read -p "Day number (1–25): " DAY
fi

[[ $DAY -ge 1 && $DAY -le 25 ]] || { echo "Invalid day"; exit 1; }

DIR=$(printf "Day_%02d" "$DAY")

[[ ! -e "$DIR" ]] || { echo "$DIR already exists!"; exit 1; }

echo "Creating $DIR ..."

mkdir "$DIR"
cd "$DIR"

# Empty input files
touch input.txt
touch test_input.txt

# Four completely blank solution files (you pick the extension later)
touch solution1
touch solution2
touch solution3
touch solution4

# The most generic Makefile ever written
cat > Makefile << 'EOF'
# Advent of Code – Day XX
# Edit the commands below to match your language

.PHONY: all part1 part2 part3 part4 test clean

all: part1 part2

part1:
    @echo "=== Part 1 ==="
    # Example commands (uncomment the one you need):
    # python3 solution1.py input.txt
    # cargo run --release --bin part1
    # go run solution1.go
    # gcc solution1.c -o solution1 && ./solution1
    # nasm -f elf64 solution1.asm && ld -o solution1 solution1.o && ./solution1
    # ./solution1 < input.txt

part2:
    @echo "=== Part 2 ==="
    # (same as above, for part 2)

part3 part4:
    @echo "No code yet"

test:
    @echo "=== Testing on test_input.txt ==="
    # ./solution1 < test_input.txt
    # python3 solution1.py test_input.txt

clean:
    rm -rf __pycache__ *.o *.out *.exe a.out solution1 solution2
EOF
