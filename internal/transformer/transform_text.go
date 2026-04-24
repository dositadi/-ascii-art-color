package transformer

func (t *Transformer) TransformTextWithColor(input []string, sub, color, font string) ([][][]string, error) {
	var output [][][]string

	for i := 0; i < len(input); i++ {
		word := input[i]

		if word == "" {
			output = append(output, [][]string{{""}})
			continue
		}

		idxS := t.getAllSubstringStartIndex(word, sub)

		wordAscii, err := t.convertToAsciiWithColor(word, sub, color, font, idxS)
		if err != nil {
			return nil, err
		}

		output = append(output, wordAscii)
	}

	return output, nil
}
