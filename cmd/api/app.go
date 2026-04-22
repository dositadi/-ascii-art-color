package api

import (
	"fmt"

	"acad.learn2earn.ng/git/dositadi/ascii-art-color/internal/transformer"
)

type Transformer interface {
	GetInput() (string, string, string, string) /*
		CheckThatSubExists(input string, substring string) bool */
	GetAnsiiColor(color string) string
}

type App struct {
	Transformer Transformer
}

func (a *App) Run() {
	a.Transformer = &transformer.Transformer{}
	a.Transform()
	//a.Transformer.CheckThatSubExists("Hello world", "o")

	str := "abc"

	for i := 1; i < len(str); i++ {
		fmt.Println(string(str[i]) + a.Transformer.GetAnsiiColor("red"))
	}
}

func (a *App) Transform() {
	a.Transformer.GetInput()
}
