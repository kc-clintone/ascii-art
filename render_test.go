package main

import "testing"

func TestRenderAsciiArt(t *testing.T) {
	letterMap := map[rune][]string{
		' ': {
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
		},
		'H': {
			" _    _  ",
			"| |  | | ",
			"| |__| | ",
			"|  __  | ",
			"| |  | | ",
			"|_|  |_| ",
			"         ",
			"         ",
		},
		'e': {
			"       ",
			"       ",
			"  ___  ",
			` / _ \ `,
			`|  __/ `,
			` \___| `,
			"       ",
			"       ",
		},
		'o': {
			`        `,
			`        `,
			`  ___   `,
			` / _ \  `,
			`| (_) | `,
			` \___/  `,
			`        `,
			`        `,
		},
		'h': {
			` _      `,
			`| |     `,
			`| |__   `,
			`|  _ \  `,
			`| | | | `,
			`|_| |_| `,
			`        `,
			`        `,
		},
		'@': {
			`          `,
			`   ____   `,
			`  / __ \  `,
			" / / _` | ",
			`| | (_| | `,
			` \ \__,_| `,
			`  \____/  `,
			`          `,
		},
		'_': {
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			" ______  ",
			"|______| ",
		},
		'0': {
			`        `,
			`  ___   `,
			` / _ \  `,
			`| | | | `,
			`| |_| | `,
			` \___/  `,
			`        `,
			`        `,
		},
	}

	t.Run("empty input", func(t *testing.T) {
		input := ""
		expected := ""

		output := RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}
	})

	t.Run("newline", func(t *testing.T) {
		input := `\n`
		expected := "\n"

		output := RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}

		input = `\n\n`
		expected = "\n\n"

		output = RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}
	})

	t.Run("space", func(t *testing.T) {
		input := " "
		expected := "      \n" +
			"      \n" +
			"      \n" +
			"      \n" +
			"      \n" +
			"      \n" +
			"      \n" +
			"      \n"

		output := RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}
	})

	t.Run("single character", func(t *testing.T) {
		input := "e"
		expected := "       \n" +
			"       \n" +
			"  ___  \n" +
			" / _ \\ \n" +
			"|  __/ \n" +
			" \\___| \n" +
			"       \n" +
			"       \n"

		output := RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}
	})

	t.Run("sequence of characters", func(t *testing.T) {
		input := "@_@"
		expected := "                             \n" +
			"   ____               ____   \n" +
			"  / __ \\             / __ \\  \n" +
			" / / _` |           / / _` | \n" +
			"| | (_| |          | | (_| | \n" +
			" \\ \\__,_|           \\ \\__,_| \n" +
			"  \\____/   ______    \\____/  \n" +
			"          |______|           \n"

		output := RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}
	})

	t.Run("sandwiched newline", func(t *testing.T) {
		input := "0\\nh"
		expected := "        \n" +
			"  ___   \n" +
			" / _ \\  \n" +
			"| | | | \n" +
			"| |_| | \n" +
			" \\___/  \n" +
			"        \n" +
			"        \n" +
			" _      \n" +
			"| |     \n" +
			"| |__   \n" +
			"|  _ \\  \n" +
			"| | | | \n" +
			"|_| |_| \n" +
			"        \n" +
			"        \n"

		output := RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}

		input = `0\n\nh`
		expected = "        \n" +
			"  ___   \n" +
			" / _ \\  \n" +
			"| | | | \n" +
			"| |_| | \n" +
			" \\___/  \n" +
			"        \n" +
			"        \n" +
			"\n" +
			" _      \n" +
			"| |     \n" +
			"| |__   \n" +
			"|  _ \\  \n" +
			"| | | | \n" +
			"|_| |_| \n" +
			"        \n" +
			"        \n"

		output = RenderAsciiArt(input, letterMap)

		if output != expected {
			t.Errorf("Expected:\n%q\nOutput:\n%q\n", expected, output)
		}
	})
}
