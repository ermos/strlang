package strlang

import (
	"fmt"
)

type Javascript struct {
	*Builder
}

func NewJavascript() *Javascript {
	return &Javascript{
		NewBuilder(),
	}
}

func (b *Javascript) If(statement string, inside func()) {
	b.Block(fmt.Sprintf("if (%s) {", statement), inside, "}")
}

func (b *Javascript) Else(inside func()) {
	b.TrimRight("\n")
	b.Block(" else {", inside, "}")
}

func (b *Javascript) ElseIf(statement string, inside func()) {
	b.TrimRight("\n")
	b.Block(fmt.Sprintf(" else if (%s) {", statement), inside, "}")
}

func (b *Javascript) Object(varType, name string, inside func()) {
	b.Block(fmt.Sprintf("%s %s = {", varType, name), inside, "}")
}

func (b *Javascript) Export() *Javascript {
	b.WriteString("export ")
	return b
}
