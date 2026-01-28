package main

import (
	"os"
	"testing"
)

func TestMain_EndToEnd(t *testing.T) {
	dummy_font := `
A0
A1
A2
A3
A4
A5
A6
A7

B0
B1
B2
B3
B4
B5
B6
B7
`

	err := os.WriteFile("standard.txt", []byte(dummy_font), 0644)
	if err != nil {
		t.Fatalf("failed to create standard.txt: %v", err)
	}
	defer os.Remove("standard.txt")

}