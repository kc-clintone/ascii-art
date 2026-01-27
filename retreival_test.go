package main

import (
	"os"
	//"reflect"
	"testing"
)

func TestRetrieveArt(t *testing.T) {
	// 1. Setup: Create a temporary mock font file
	mockData := "\nline1\nline2\nline3\nline4\nline5\nline6\nline7\nline8\n\nnext1\nnext2\nnext3\nnext4\nnext5\nnext6\nnext7\nnext8"
	tmpFile := "test_font.txt"
	err := os.WriteFile(tmpFile, []byte(mockData), 0644)
	if err != nil {
		t.Fatal("Could not create temporary test file")
	}
	defer os.Remove(tmpFile) // Clean up after the test

	t.Run("Valid file returns correct map", func(t *testing.T) {
		font, err := RetrieveArt(tmpFile)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if ' ' (space) was mapped (the first character in your loop)
		if _, ok := font[' ']; !ok {
			t.Error("Expected character ' ' to be in the map")
		}

		// Check height of the character
		if len(font[' ']) != 8 {
			t.Errorf("Expected character height 8, got %d", len(font[' ']))
		}
	})

	t.Run("Non-existent file returns error", func(t *testing.T) {
		_, err := RetrieveArt("ghost.txt")
		if err == nil {
			t.Error("Expected an error for a non-existent file, but got nil")
		}
	})

	t.Run("Data accuracy", func(t *testing.T) {
		font, _ := RetrieveArt(tmpFile)
		expectedFirstLine := "line1"
		if font[' '][0] != expectedFirstLine {
			t.Errorf("Expected first line to be %s, got %s", expectedFirstLine, font[' '][0])
		}
	})
}