package transformer

import (
	"fmt"
	"slices"
	"strings"
)

func (t *Transformer) PrintAsciiOnTerminal(input [][][]string) {
	if len(input) == 0 {
		return
	}

	for i := 0; i < len(input); i++ {
		currentWord := input[i]

		if slices.Compare(currentWord[0], []string{""}) == 0 {
			fmt.Println()
			continue
		}

		for j := 0; j < 8; j++ {
			var result strings.Builder
			for k := 0; k < len(currentWord); k++ {
				result.WriteString(currentWord[k][j])
			}

			fmt.Println(result.String())
		}
	}

}
