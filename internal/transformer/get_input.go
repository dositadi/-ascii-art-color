package transformer

import (
	"os"
	"strings"
)

func (t *Transformer) GetInput() (string, string, string, string) {
	if len(os.Args) == 1 {
		PrintUsage()
		return "", "", "", ""
	}

	args := os.Args[1:]

	var color, substring, text, font string

	switch len(args) {
	case 1:
		text = args[0]
		font = "standard"
	case 2:
		text = args[0]
		if !(args[1] == "standard" || args[1] == "tinkertoy" || args[1] == "shadow") {
			PrintUsage()
			return "", "", "", ""
		}
		font = args[1]
	case 3:
		if !strings.HasPrefix(args[0], "--color=") {
			PrintUsage()
			return "", "", "", ""
		}
		color = strings.TrimPrefix(args[0], "--color=")
		// Check that the color is valid
		substring = args[1]
		text = args[2]
	case 4:
		if !strings.HasPrefix(args[0], "--color=") {
			PrintUsage()
			return "", "", "", ""
		}
		color = strings.TrimPrefix(args[0], "--color=")
		// Check that the color is valid
		substring = args[1]
		text = args[2]
		if !(args[3] == "standard" || args[3] == "tinkertoy" || args[3] == "shadow") {
			PrintUsage()
			return "", "", "", ""
		}
		font = args[3]
	}
	return color, substring, text, font
}
