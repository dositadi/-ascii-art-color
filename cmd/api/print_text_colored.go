package api

import (
	"fmt"
	"strings"
)

func (a *App) PrintTextColored(cleanedText, substring, color string, subIdxs []int) {
	var res strings.Builder

	i := 0

	for i < len(cleanedText) {
		check := false
		for j := 0; j < len(subIdxs); j++ {
			fmt.Println(i, subIdxs[j])
			if i == subIdxs[j] {
				check = true
			}
		}

		if check {
			res.WriteString(a.Transformer.GetAnsiiColor(color) + cleanedText[i:i+len(substring)] + a.Transformer.GetAnsiiColor("reset"))
			i += len(substring)
			continue
		} else {
			res.WriteString(string(cleanedText[i]))
			i += 1
			continue
		}
	}

	fmt.Println(res.String())
}
