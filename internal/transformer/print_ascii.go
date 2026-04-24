package transformer

import (
	"fmt"
)

func (t *Transformer) PrintAsciiOnTerminal(input [][]string) {
	if len(input) == 0 {
		return
	}

	for j := 0; j < 8; j++ {
		for _, char := range input {
			if len(char) == 1 {
				continue
			}
			fmt.Print(char[j])
		}
		fmt.Println()
	}
}
