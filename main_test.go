package main

import (
	"io"
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