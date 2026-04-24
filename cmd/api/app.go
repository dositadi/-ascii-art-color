package api

import (
	"acad.learn2earn.ng/git/dositadi/ascii-art-color/internal/transformer"
)

type Transformer interface {
	GetInput() []string
	ReplaceWithNewLine(input string) []string
	TransformTextWithColor(input []string, sub, color, font string) ([][][]string, error)
	TransformTextWithoutColor(input []string, font string) ([][][]string, error)
	PrintAsciiOnTerminal(input [][][]string)
}

type App struct {
	Transformer Transformer
}

func (a *App) Run() {
	a.Transformer = &transformer.Transformer{}
	a.Transform()
}

func (a *App) Transform() {
	parameters := a.Transformer.GetInput()

	if parameters == nil {
		return
	}

	if len(parameters) == 4 {
		a.printWithColor(parameters[0], parameters[1], parameters[2], parameters[3])
	} else if len(parameters) == 2 {
		a.PrintWithoutColor(parameters[0], parameters[1])
	} else {
		transformer.PrintUsage()
	}
}

func (a *App) printWithColor(color, substring, text, font string) {
	cleanedText := a.Transformer.ReplaceWithNewLine(text)

	asciiText, err := a.Transformer.TransformTextWithColor(cleanedText, substring, color, font)
	if err != nil {
		transformer.PrintUsage()
		return
	}

	a.Transformer.PrintAsciiOnTerminal(asciiText)
}

func (a *App) PrintWithoutColor(text, font string) {
	cleanedText := a.Transformer.ReplaceWithNewLine(text)

	asciiText, err := a.Transformer.TransformTextWithoutColor(cleanedText, font)
	if err != nil {
		transformer.PrintUsage()
		return
	}

	a.Transformer.PrintAsciiOnTerminal(asciiText)
}
