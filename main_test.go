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
	tmpDir := t.TempDir()

	fontPath := tmpDir + "/standard.txt"
	if err := os.WriteFile(fontPath, []byte(generateFont('B')), 0644); err != nil {
		t.Fatal(err)
	}

	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)
	os.Chdir(tmpDir)

	origArgs := os.Args
	os.Args = []string{"ascii-art", "AB"}
	defer func() { os.Args = origArgs }()

	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = origStdout

	out, _ := io.ReadAll(r)

		expected := "" +
		"A0B0\n" +
		"A1B1\n" +
		"A2B2\n" +
		"A3B3\n" +
		"A4B4\n" +
		"A5B5\n" +
		"A6B6\n" +
		"A7B7\n"

	if string(out) != expected {
		t.Fatalf("expected %q, got %q", expected, string(out))
	}
}