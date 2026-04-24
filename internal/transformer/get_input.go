package transformer

import (
	"fmt"
	"os"
	"strings"
)

func (t *Transformer) GetInput() []string {
	if len(os.Args) == 1 {
		PrintUsage()
		return nil
	}

	args := os.Args[1:]

	var color, substring, text, font string

	switch len(args) {
	case 1:
		text = args[0]
		font = "standard"

		return []string{text, font}

	case 2:
		text = args[0]
		if !(args[1] == "standard" || args[1] == "tinkertoy" || args[1] == "shadow") {
			PrintUsage()
			return nil
		}
		font = args[1]

		return []string{text, font}

	case 3:
		if !strings.HasPrefix(args[0], "--color=") {
			PrintUsage()
			return nil
		}
		color = strings.TrimPrefix(args[0], "--color=")
		if !t.colorValid(color) {
			PrintUsage()
			t.availableColors()
			return nil
		}
		substring = args[1]
		text = args[2]
		font = "standard"

		return []string{color, substring, text, font}

	case 4:
		if !strings.HasPrefix(args[0], "--color=") {
			PrintUsage()
			return nil
		}
		color = strings.TrimPrefix(args[0], "--color=")
		fmt.Println(color)
		if !t.colorValid(color) {
			PrintUsage()
			t.availableColors()
			return nil
		}

		substring = args[1]
		text = args[2]
		if !(args[3] == "standard" || args[3] == "tinkertoy" || args[3] == "shadow") {
			PrintUsage()
			return nil
		}
		font = args[3]

		return []string{color, substring, text, font}
	}
	return nil
}

func (t *Transformer) colorValid(color string) bool {
	return strings.ToLower(color) == "green" || strings.ToLower(color) == "yellow" || strings.ToLower(color) == "blue" || strings.ToLower(color) == "purple" || strings.ToLower(color) == "cyan" || strings.ToLower(color) == "white" || strings.ToLower(color) == "red" || strings.ToLower(color) == "reset"
}

func (t *Transformer) availableColors() {
	fmt.Println("The available colors are:\nGreen\nYellow\nBlue\nPurple\nCyan\nWhite")
}
