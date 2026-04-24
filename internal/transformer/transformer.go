package transformer

import (
	"fmt"
)

type Transformer struct{}

func (t *Transformer) GetAnsiiColor(color string) string {
	ansiiColor := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"reset":  "\033[0m",
	}

	if val, ok := ansiiColor[color]; ok {
		return val
	}
	return ""
}

func PrintUsage() {
	str := "Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> something"
	fmt.Println(str)
}
