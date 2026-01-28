# ASCII Art Generator (using Golang)

A small but well-structured command‑line program written in Go that converts normal text into banner‑style ASCII art using a font definition file.

It demonstrates:

* File parsing
* Rune‑based text processing
* CLI argument handling
* Terminal output formatting
* Proper unit testing (including `main`)

---

## Project Structure

```
.
├── main.go             # The program entry point
├── retrieval.go        # Font file parser
├── retreival_test.go   # Tests for retrieval function
├── render.go           # ASCII rendering engine
├── render_test.go      # Tests for the rendering function
├── main_test.go        # General test suite for main()
├── standard.txt        # ASCII font definition file
└── go.mod
```

All Go files belong to the same package:

```go
package main
```

This allows the functions to be shared across files automatically.

---

## How It Works

1. The user runs the program with text as arguments
2. The program loads `standard.txt`
3. The font file is parsed into a map of characters → ASCII art
4. The input text is converted line‑by‑line into ASCII art
5. The result is printed to stdout

Pipeline view:

```
User_Input → main.go → RetrieveArt() → AsciiArt() → Terminal_Output
```

---

## 🧠 Core Components

---

## main.go – Program Entry Point

This file controls the application flow.

### Responsibilities

* Read command‑line arguments
* Load the font file
* Call the renderer
* Print the result
* Handle failures safely

### Code Flow

```go
if len(os.Args) < 2 {
    return
}
```

No arguments → program exits silently.

---

```go
input := strings.Join(os.Args[1:], " ")
```

All arguments are combined into a single string:

```
./ascii Hello World
→ "Hello World"
```

---

```go
letterMap, err := RetrieveArt("standard.txt")
```

Loads and parses the font definition file.

---

```go
result := AsciiArt(input, letterMap)
fmt.Print(result)
```

Renders and prints the ASCII art.

---

## retrieval.go – Font File Parser

This file reads and interprets the ASCII font file.

### Purpose

Convert this:

```
  AAAAA
 A     A
 AAAAAAA
```

Into this Go structure:

```go
map[rune][]string
```

Example entry:

```go
'A' → ["  AAAAA", " A     A", ...]
```

---

### How Parsing Works

1. Read entire file into memory
2. Split into lines
3. Skip optional first empty line
4. Read characters in blocks of 8 lines
5. Map each block to a rune

```go
ascii := ' '
charHeight := 8
```

Characters are read in ASCII order from space (`32`) to tilde (`126`).

---

### Return Value

```go
(map[rune][]string, error)
```

This map becomes the rendering dictionary.

---

## render.go – ASCII Rendering Engine

This file converts user input into formatted ASCII art.

---

### Rendering Algorithm

For each line of input:

1. Loop 8 times (the font height)
2. For each character in the line:

   * Fetch its ASCII art line
   * Append it horizontally
3. Append completed row to output

---

### Core Loop (Conceptual)

```
Input: "Hi"

H line 0 + i line 0
H line 1 + i line 1
...
```

Result becomes a vertical stack of composed lines.

---

### Output Formatting

* Lines joined using `\n`
* Ends with newline if output is not empty

---

## Testing – main_test.go

The program includes a **automated test suite** that verifies real behavior instead of mocking everything.

It tests:

* No arguments
* Missing font file
* Single character rendering
* Multiple word rendering
* Output formatting
* Newline correctness

---

### Testing Techniques Used

| Technique                   | Purpose                         |
| --------------------------- | ------------------------------- |
| `os.Args` override          | Simulate CLI input              |
| `os.Pipe`                   | Capture stdout                  |
| `t.TempDir()`               | Create isolated test filesystem |
| Temporary font file         | Avoid dependency on real files  |
| Working directory switching | Test file loading logic         |

---

### Example: Capturing Output

```go
r, w, _ := os.Pipe()
os.Stdout = w
```

Allows verifying what the program prints.

---

## standard.txt – Font Definition File

This file defines how every printable ASCII character looks.

Format:

```
<optional empty line>
<8 lines for space>
<empty line>
<8 lines for !>
<empty line>
...
```

Each character block is exactly 8 lines tall.

---

## Running the Program

```bash
go run . "Hello World"
```

Or after building:

```bash
go build
./ascii "Hello World"
```

---

## Running Tests

```bash
go test -v
```

---

## Summary

This project is a clean example of how to build:

* A CLI program
* With structured parsing
* Safe rendering
* And testing