// testall.go
// Drop this file in your advent/ root folder.
// Then run: go test ./...
// It will automatically discover and run ALL tests in every Day_XX folder.

package main

import (
	"fmt"
	"os"
	"testing"
)

// This makes `go test` show a nice summary at the end
func TestMain(m *testing.M) {
	fmt.Println("Running tests for all Advent of Code days...")
	code := m.Run()
	os.Exit(code)
}

// Optional: add a fun banner
func TestAdventOfCode(t *testing.T) {
	// This test always passes — just here so `go test` shows something
	// when there are no real tests in some days
	t.Log("All days loaded — running individual day tests...")
}