package transformer

import "strings"

func (t *Transformer) ReplaceWithNewLine(input string) []string {
	if input == "" {
		return nil
	}
	output := strings.Split(input, "\\n")

	return output
}
