package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// handle input
	if len(os.Args) < 2 {
		return
	}

	input := strings.Join(os.Args[1:], " ")

	// load the ascii standard.txt file
	file_name := "standard.txt"
	letterMap, err := RetrieveArt(file_name)
	if err != nil {
		fmt.Println("Failed to load file: ", err)
		return
	}

	// rendering the Art
	result := RenderAsciiArt(input, letterMap)
	fmt.Print(result)
}
