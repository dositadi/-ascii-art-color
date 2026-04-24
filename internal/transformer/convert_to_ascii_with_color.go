package transformer

func (t *Transformer) convertToAsciiWithColor(input, sub, color, font string, idxs []int) ([][]string, error) {
	var output [][]string
	i := 0

	for i < len(input) {
		check := false

		if rune(input[i]) == '\n' {
			output = append(output, []string{""})

			i += 1
			continue
		}

		for j := 0; j < len(idxs); j++ {
			if i == idxs[j] {
				check = true
			}
		}

		if check {
			for k := i; k < i+len(sub); k++ {
				asciiChar, err := t.readCharFromFont(rune(input[k]), color, font)
				if err != nil {
					return nil, err
				}
				output = append(output, asciiChar)
			}

			i += len(sub)
		} else {
			asciiChar, err := t.readCharFromFont(rune(input[i]), "", font)
			if err != nil {
				return nil, err
			}
			output = append(output, asciiChar)

			i += 1
		}
	}
	return output, nil
}
