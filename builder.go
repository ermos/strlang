// Package strlang provides a string builder with indentation support.
package strlang

import (
	"strings"
)

type Builder struct {
	builder    *strings.Builder
	currIndent int
	indentChar string
}

func NewBuilder() *Builder {
	return &Builder{
		builder:    &strings.Builder{},
		currIndent: 0,
		indentChar: "  ",
	}
}

func (b *Builder) SetIndentChar(s string) {
	b.indentChar = s
}

func (b *Builder) Block(start string, f func(), end string, ln ...int) {
	b.WriteStringln(start)
	b.Indent()
	f()
	b.StripIndent()
	b.WriteStringln(end, ln...)
}

func (b *Builder) WriteStringln(s string, nb ...int) {
	n := 1
	if len(nb) != 0 {
		n = nb[0]
	}

	b.Write(append([]byte(s), []byte(strings.Repeat("\n", n))...))
}

func (b *Builder) WriteString(s string) {
	b.Write([]byte(s))
}

func (b *Builder) Write(p []byte) {
	tab := []byte(strings.Repeat(b.indentChar, b.currIndent))
	b.builder.Write(append(tab, p...))
}

func (b *Builder) WriteNoIdentString(s string) {
	old := b.currIndent
	b.currIndent = 0

	b.Write([]byte(s))

	b.currIndent = old
}

func (b *Builder) TrimLeft(cutset string) {
	old := strings.TrimLeft(b.builder.String(), cutset)
	b.builder.Reset()
	b.builder.WriteString(old)
}

func (b *Builder) TrimRight(cutset string) {
	old := strings.TrimRight(b.builder.String(), cutset)
	b.builder.Reset()
	b.builder.WriteString(old)
}

func (b *Builder) String() string {
	return b.builder.String()
}

func (b *Builder) Indent(nb ...int) {
	n := 1
	if len(nb) != 0 {
		n = nb[0]
	}
	b.currIndent += n
}

func (b *Builder) StripIndent(nb ...int) {
	n := 1
	if len(nb) != 0 {
		n = nb[0]
	}
	b.currIndent -= n
}
