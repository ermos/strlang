package strlang

import (
	"fmt"
)

// PHP represents a PHP code generator that extends the Builder struct.
type PHP struct {
	// namespace is a string representing the namespace of the generated code
	namespace string
	// Builder is an embedded struct that represents
	// the base builder for generating the PHP code.
	*Builder
}

// NewPHP returns a new instance of PHP code builder.
func NewPHP(namespace string) *PHP {
	return &PHP{
		namespace,
		NewBuilder(),
	}
}

// If generates an if statement in the generated code.
func (b *PHP) If(statement string, inside func(), ln ...int) {
	b.Block(fmt.Sprintf("if (%s) {", statement), inside, "}", ln...)
}

// Else generates an else statement in the generated code.
func (b *PHP) Else(inside func()) {
	b.TrimRight("\n")
	b.Block(" else {", inside, "}")
}

// ElseIf generates an else if statement in the generated code.
func (b *PHP) ElseIf(statement string, inside func()) {
	b.TrimRight("\n")
	b.Block(fmt.Sprintf(" elseif (%s) {", statement), inside, "}")
}

// Class generates a class definition in the generated code.
func (b *PHP) Class(name string, inside func()) {
	b.Block(fmt.Sprintf("class %s {", name), inside, "}", 2)
}

// ClassFunc generates class method inside a class in the generated code.
func (b *PHP) ClassFunc(modifiers, name, parameters, output string, inside func()) {
	if modifiers != "" {
		modifiers += " "
	}
	b.Block(fmt.Sprintf("%sfunction %s(%s): %s {", modifiers, name, parameters, output), inside, "}", 2)
}

// String returns the generated code as a string.
func (b *PHP) String() string {
	sb := NewBuilder()

	sb.SetIndentChar(b.indentChar)
	sb.StripIndent(b.currIndent)

	sb.WriteStringln("<?php", 2)

	if b.namespace != "" {
		sb.WriteStringln(fmt.Sprintf("namespace %s;", b.namespace), 2)
	}

	sb.WriteString(b.builder.String())

	return sb.String()
}
