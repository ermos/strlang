// Package strlang provides a string builder with indentation support,
// it allows you to generate code using a fluent string builder interface.
package strlang

import (
	"strings"
)

// Builder is the base builder of the strlang package
type Builder struct {
	// builder that contains a string builder.
	builder *strings.Builder
	// currIndent is the current indentation level,
	// current indentation start at 0.
	currIndent int
	// indentChar represent the indentation character,
	// default indent character is `\t`.
	indentChar string
}

// NewBuilder returns a new instance of Builder with default settings:
// an empty strings.Builder and an indentation character of two spaces.
func NewBuilder() *Builder {
	return &Builder{
		builder:    &strings.Builder{},
		currIndent: 0,
		indentChar: "\t",
	}
}

// SetIndentChar sets the indentation character for the builder.
// The default indentation character is one tab.
func (b *Builder) SetIndentChar(s string) {
	b.indentChar = s
}

// Block writes the start string, indents the builder,
// calls the provided function f, strips the indent, and writes the end string with a newline character.
// Optionally, it can write additional newline characters with the ln argument.
func (b *Builder) Block(start string, f func(), end string, ln ...int) {
	b.WriteStringln(start)
	b.Indent()
	f()
	b.StripIndent()
	b.WriteStringln(end, ln...)
}

// WriteStringln writes a string with a newline character to the builder.
// Optionally, it can write multiple newline characters by passing an integer argument.
func (b *Builder) WriteStringln(s string, nb ...int) {
	n := 1
	if len(nb) != 0 {
		n = nb[0]
	}

	b.Write(append([]byte(s), []byte(strings.Repeat("\n", n))...))
}

// WriteString writes a string to the builder without adding a newline character.
func (b *Builder) WriteString(s string) {
	b.Write([]byte(s))
}

// Write writes a byte slice to the builder with indentation.
// It adds the current indentation characters before the provided byte slice.
func (b *Builder) Write(p []byte) {
	tab := []byte(strings.Repeat(b.indentChar, b.currIndent))
	b.builder.Write(append(tab, p...))
}

// WriteNoIdentString writes a string to the builder without indentation.
func (b *Builder) WriteNoIdentString(s string) {
	old := b.currIndent
	b.currIndent = 0

	b.Write([]byte(s))

	b.currIndent = old
}

// TrimLeft trims the resulting string from the builder by removing characters
// from the left side of the string that are contained in the provided cutset string.
func (b *Builder) TrimLeft(cutset string) {
	old := strings.TrimLeft(b.builder.String(), cutset)
	b.builder.Reset()
	b.builder.WriteString(old)
}

// TrimRight trims the resulting string from the builder by removing characters
// from the right side of the string that are contained in the provided cutset string.
func (b *Builder) TrimRight(cutset string) {
	old := strings.TrimRight(b.builder.String(), cutset)
	b.builder.Reset()
	b.builder.WriteString(old)
}

// String returns the resulting string built by the builder.
func (b *Builder) String() string {
	return b.builder.String()
}

// Indent increases the current indentation level by the number of spaces
// passed as an optional argument. The default indentation increase is two spaces.
func (b *Builder) Indent(nb ...int) {
	n := 1
	if len(nb) != 0 {
		n = nb[0]
	}
	b.currIndent += n
}

// StripIndent decreases the current indentation level by the number of
// spaces passed as an optional argument. The default indentation decrease is two spaces.
func (b *Builder) StripIndent(nb ...int) {
	n := 1
	if len(nb) != 0 {
		n = nb[0]
	}
	b.currIndent -= n
}
