package strlang

import (
	"fmt"
)

type PHP struct {
	*Builder
	namespace string
}

func NewPHP(namespace string) *PHP {
	return &PHP{
		NewBuilder(),
		namespace,
	}
}

func (b *PHP) If(statement string, inside func(), ln ...int) {
	b.Block(fmt.Sprintf("if (%s) {", statement), inside, "}", ln...)
}

func (b *PHP) Else(inside func()) {
	b.TrimRight("\n")
	b.Block(" else {", inside, "}")
}

func (b *PHP) ElseIf(statement string, inside func()) {
	b.TrimRight("\n")
	b.Block(fmt.Sprintf(" elseif (%s) {", statement), inside, "}")
}

func (b *PHP) Class(name string, inside func()) {
	b.Block(fmt.Sprintf("class %s {", name), inside, "}", 2)
}

func (b *PHP) ClassFunc(modifiers, name, parameters, output string, inside func()) {
	if modifiers != "" {
		modifiers += " "
	}
	b.Block(fmt.Sprintf("%sfunction %s(%s): %s {", modifiers, name, parameters, output), inside, "}", 2)
}

func (b *PHP) String() string {
	sb := NewBuilder()

	sb.WriteStringln("<?php", 2)

	if b.namespace != "" {
		sb.WriteStringln(fmt.Sprintf("namespace %s;", b.namespace), 2)
	}

	sb.WriteString(b.builder.String())

	return sb.String()
}
