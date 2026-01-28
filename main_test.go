package main

import (
	"bytes"
	"os"
	"testing"
)

// test for capturing output
func captureOutput(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

// testin gthe program -- argumrnsts provided
func withArgs(args []string, fn func()) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = args
	fn()
}
// testing the program -- no arguments provided
func NoArguments(t *testing.T) {
	output := captureOutput(t, func() {
		withArgs([]string{"app"}, main)
	})

	if output != "" {
		t.Fatalf("expected no output, got: %q", output)
	}
}