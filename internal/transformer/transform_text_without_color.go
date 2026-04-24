package transformer

func (t *Transformer) TransformTextWithoutColor(input []string, font string) ([][][]string, error) {
	var output [][][]string

	for i := 0; i < len(input); i++ {
		word := input[i]
		var wordAscii [][]string

		if word == "" {
			wordAscii = append(wordAscii, []string{""})
			output = append(output, wordAscii)
			continue
		}

		for _, char := range word {
			charAscii, err := t.readCharFromFont(char, "", font)
			if err != nil {
				return nil, err
			}
			wordAscii = append(wordAscii, charAscii)
		}
		output = append(output, wordAscii)
	}
	return output, nil
}
