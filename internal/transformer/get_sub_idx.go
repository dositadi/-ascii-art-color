package transformer

import (
	"strings"
)

func (t *Transformer) getAllSubstringStartIndex(input, sub string) []int {
	var output []int
	actualIndex := 0
	offset := 0

	for {
		if offset >= len(input) {
			break
		}

		idx := strings.Index(input[offset:], sub)
		if idx == -1 {
			break
		}

		actualIndex = offset + idx

		output = append(output, actualIndex)

		offset = actualIndex + len(sub)
	}

	return output
}
