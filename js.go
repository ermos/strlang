package strlang

import (
	"fmt"
)

// Javascript represents a JavaScript code generator that extends the Builder struct.
type Javascript struct {
	// Builder is an embedded struct that represents
	// the base builder for generating the Javascript code.
	*Builder
}

// NewJavascript returns a new instance of Javascript code builder.
func NewJavascript() *Javascript {
	return &Javascript{
		NewBuilder(),
	}
}

// If generates an if statement in the generated code.
func (b *Javascript) If(statement string, inside func()) {
	b.Block(fmt.Sprintf("if (%s) {", statement), inside, "}")
}

// Else generates an else statement in the generated code.
func (b *Javascript) Else(inside func()) {
	b.TrimRight("\n")
	b.Block(" else {", inside, "}")
}

// ElseIf generates an else if statement in the generated code.
func (b *Javascript) ElseIf(statement string, inside func()) {
	b.TrimRight("\n")
	b.Block(fmt.Sprintf(" else if (%s) {", statement), inside, "}")
}

// Object generates an object block with the provided variable type, name, and inside function.
func (b *Javascript) Object(varType, name string, inside func()) {
	b.Block(fmt.Sprintf("%s %s = {", varType, name), inside, "}")
}

// Export appends the "export" keyword and allows to define an exportable object or variable.
func (b *Javascript) Export() *Javascript {
	b.WriteString("export ")
	return b
}
