package main

import (
	"fmt"
	"os"
	"strings"
)

func RetrieveArt(str string) (map[rune][]string, error) {
	// url := "./standard.txt"

	// defer response.Body.Close()
	data, err := os.ReadFile(str)
	if err != nil {
		fmt.Println(err)
		return map[rune][]string{}, err
	}
	// split into lines
	lines := strings.Split(string(data), "\n")
	// io reads eveything from a stream and returns in bytes
	// strings.Split returns a slice

	if lines[0] == "" {
		lines = lines[1:]
	}

	font := make(map[rune][]string)

	charHeight := 8
	ascii := ' ' // first printable ascii character
	lineIndex := 0

	for ascii <= '~' && lineIndex+charHeight <= len(lines) {
		charLines := lines[lineIndex : lineIndex+charHeight]
		font[ascii] = charLines
		ascii++
		lineIndex += charHeight + 1
	}
	return font, nil
}
