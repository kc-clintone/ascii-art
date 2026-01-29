package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

const (
	firstChar = ' '
	lastChar  = '~'
	height    = 8
)

func generateFont(from, to rune) string {
	var b strings.Builder

	for r := from; r <= to; r++ {
		for i := 0; i < height; i++ {
			b.WriteRune(r)
			b.WriteString(strconv.Itoa(i))
			b.WriteString(" \n")
		}
		b.WriteByte('\n')
	}

	return b.String()
}

func buildBinary(t *testing.T, dir string) string {
	bin := filepath.Join(dir, "ascii-art")

	cmd := exec.Command("go", "build", "-o", bin, ".")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}

	return bin
}

func runProgram(t *testing.T, bin, dir, input string) string {
	cmd := exec.Command(bin, input)
	cmd.Dir = dir

	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("program failed: %v", err)
	}

	return string(out)
}

func expectedSingleChar(r rune) string {
	var b strings.Builder

	for i := 0; i < height; i++ {
		b.WriteRune(r)
		b.WriteString(strconv.Itoa(i))
		b.WriteString(" \n")
	}

	return b.String()
}

func expectedAllChars() string {
	var b strings.Builder

	for i := 0; i < height; i++ {
		for r := firstChar; r <= lastChar; r++ {
			b.WriteRune(r)
			b.WriteString(strconv.Itoa(i))
			b.WriteRune(' ')
		}
		b.WriteByte('\n')
	}

	return b.String()
}

func TestMainExec(t *testing.T) {
	tmpDir := t.TempDir()

	// Write fake font
	fontPath := filepath.Join(tmpDir, "standard.txt")
	if err := os.WriteFile(fontPath, []byte(generateFont(firstChar, lastChar)), 0o644); err != nil {
		t.Fatal(err)
	}

	// Build program once
	bin := buildBinary(t, tmpDir)

	t.Run("each character individually", func(t *testing.T) {
		for r := firstChar; r <= lastChar; r++ {
			out := runProgram(t, bin, tmpDir, string(r))
			exp := expectedSingleChar(r)

			if out != exp {
				t.Fatalf("char %q: expected %q, got %q", r, exp, out)
			}
		}
	})

	t.Run("all characters on one line", func(t *testing.T) {
		var input strings.Builder
		for r := firstChar; r <= lastChar; r++ {
			input.WriteRune(r)
		}

		out := runProgram(t, bin, tmpDir, input.String())
		exp := expectedAllChars()

		if out != exp {
			t.Fatalf("expected %q, got %q", exp, out)
		}
	})
}
