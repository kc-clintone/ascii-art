package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func generateFont(upTo rune) string {
	var b strings.Builder
	for r := ' '; r <= upTo; r++ {
		for i := range 8 {
			b.WriteString(string(r))
			// b.WriteString(byte('0' + i))
			fmt.Fprintf(&b, "%d", i)

			b.WriteString("\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func TestMain(t *testing.T) {

	dummy_font := generateFont('B')

	err := os.WriteFile("standard.txt", []byte(dummy_font), 0644)
	if err != nil {
		t.Fatalf("failed to create standard.txt: %v", err)
	}
	defer os.Remove("standard.txt")

	originalArgs := os.Args
	os.Args = []string{"ascii-art", "AB"}
	defer func() { os.Args = originalArgs }()

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = origStdout

	outBytes, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}

		expected := "" +
		"A0B0\n" +
		"A1B1\n" +
		"A2B2\n" +
		"A3B3\n" +
		"A4B4\n" +
		"A5B5\n" +
		"A6B6\n" +
		"A7B7\n"

	if string(outBytes) != expected {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, string(outBytes))
	}
}