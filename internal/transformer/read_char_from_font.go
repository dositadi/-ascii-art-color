package transformer

import (
	"bufio"
	"os"
	"strings"
)

func (t *Transformer) readCharFromFont(char rune, color string, font string) ([]string, error) {
	var path string
	var output []string

	switch font {
	case "standard":
		path = "font/standard.txt"
	case "shadow":
		path = "font/shadow.txt"
	case "tinkertoy":
		path = "font/tinkertoy.txt"
	default:
		path = "font/standard.txt"
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	startLine := ((char - ' ') * 9) + 1
	endLine := int(startLine) + 8
	currentLine := 0

	for scanner.Scan() {
		if currentLine >= int(startLine) && currentLine <= endLine {
			var coloredText strings.Builder
			if color == "" {
				line := scanner.Text()
				output = append(output, line)
			} else {
				coloredText.WriteString(t.GetAnsiiColor(color) + scanner.Text() + t.GetAnsiiColor("reset"))
				output = append(output, coloredText.String())
			}
		} else if currentLine > endLine {
			break
		}
		currentLine += 1
	}
	return output, nil
}
