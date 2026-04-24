package transformer

import "strings"

func (t *Transformer) ReplaceWithNewLine(input string) string {
	if input == "" {
		return ""
	}
	output := strings.ReplaceAll(input, "\\n", "\n")
	return output
}
