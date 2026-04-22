package transformer

import (
	"fmt"
	"os"
	"strings"
)

type Transformer struct{}

func (t *Transformer) GetAnsiiColor(color string) string {
	ansiiColor := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"reset":  "\033[0m",
	}

	if val, ok := ansiiColor[color]; ok {
		return val
	}
	return ""
}

func PrintUsage() {
	str := "Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> something"
	fmt.Println(str)
}

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

	fmt.Println(color, text, substring, font)
	return color, substring, text, font
}
