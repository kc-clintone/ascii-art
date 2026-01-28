package main

import "strings"

const characterHeight = 8

// Render the ASCII art for a given input string.
func RenderAsciiArt(input string, letterMap map[rune][]string) string {
	outputLines := []string(nil)

	input = strings.ReplaceAll(input, "\\n", "\n")

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			outputLines = append(outputLines, "")
			continue
		}

		for i := range characterHeight {
			artLine := strings.Builder{}

			for _, c := range line {
				artLine.WriteString(letterMap[c][i])
			}

			outputLines = append(outputLines, artLine.String())
		}
	}

	output := strings.Join(outputLines, "\n")

	if len(strings.Trim(output, "\n")) > 0 {
		output += "\n"
	}

	return output
}
