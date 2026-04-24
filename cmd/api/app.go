package api

import (
	"fmt"

	"acad.learn2earn.ng/git/dositadi/ascii-art-color/internal/transformer"
)

type Transformer interface {
	GetInput() (string, string, string, string)
	GetAnsiiColor(color string) string
	ReplaceWithNewLine(input string) string
	GetAllSubstringStartIndex(input, sub string) []int
	ReadCharFromFont(char rune, color string, font string) ([]string, error)
	ConvertToAscii(input, sub, color, font string, idxs []int) ([][]string, error)
	PrintAsciiOnTerminal(input [][]string)
}

type App struct {
	Transformer Transformer
}

func (a *App) Run() {
	a.Transformer = &transformer.Transformer{}
	a.Transform()
}

func (a *App) Transform() {
	color, substring, text, font := a.Transformer.GetInput()

	fmt.Println(color, substring, text, font)

	cleanedText := a.Transformer.ReplaceWithNewLine(text)

	idxs := a.Transformer.GetAllSubstringStartIndex(cleanedText, substring)

	asciiWords, err := a.Transformer.ConvertToAscii(cleanedText, substring, color, font, idxs)
	if err != nil {
		transformer.PrintUsage()
		return
	}

	a.Transformer.PrintAsciiOnTerminal(asciiWords)
}
