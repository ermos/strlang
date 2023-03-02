package strlang

import (
	"testing"
)

func TestBuilder_Block(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.Block("if true {", func() {
			b.WriteStringln("return false")
		}, "}")

		expected := "" +
			"if true {\n" +
			" return false\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestBuilder_Block_Nested(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.Block("if a {", func() {
			b.Block("if b {", func() {
				b.WriteStringln("return c")
			}, "}")
		}, "}")

		expected := "" +
			"if a {\n" +
			" if b {\n" +
			"  return c\n" +
			" }\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestBuilder_WriteStringln(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.WriteStringln("Hello, world!", 2)
		b.WriteStringln("It's a test :)")

		expected := "" +
			"Hello, world!\n\n" +
			"It's a test :)\n"

		return t, expected, b.String()
	})
}

func TestBuilder_Indent(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.WriteStringln("Hello, world!")
		b.Indent()
		b.WriteStringln("==>")
		b.Indent(2)
		b.WriteStringln("==>")

		expected := "" +
			"Hello, world!\n" +
			" ==>\n" +
			"   ==>\n"

		return t, expected, b.String()
	})
}

func TestBuilder_StripIndent(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.WriteStringln("Hello, world!")
		b.Indent(5)
		b.WriteStringln("==>")
		b.StripIndent()
		b.WriteStringln("==>")
		b.StripIndent(2)
		b.WriteStringln("==>")

		expected := "" +
			"Hello, world!\n" +
			"     ==>\n" +
			"    ==>\n" +
			"  ==>\n"

		return t, expected, b.String()
	})
}

func TestBuilder_TrimLeft(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.WriteStringln(">>>>>>>>>>Hello, world!")
		b.TrimLeft(">")

		expected := "Hello, world!\n"

		return t, expected, b.String()
	})
}

func TestBuilder_TrimRight(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.WriteStringln("Hello, world!", 30)
		b.TrimRight("\n")

		expected := "Hello, world!"

		return t, expected, b.String()
	})
}

func TestBuilder_WriteNoIdentString(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.WriteNoIdentString("John Doe: ")
		b.WriteStringln("Hello, world!")

		expected := "John Doe: Hello, world!\n"

		return t, expected, b.String()
	})
}

func TestBuilder_WriteString(t *testing.T) {
	testBuilderInit(func(b *Builder) (*testing.T, string, string) {
		b.WriteString("Hello, world!")

		expected := "Hello, world!"

		return t, expected, b.String()
	})
}

func testBuilderInit(fn func(b *Builder) (t *testing.T, expected, given string)) {
	b := NewBuilder()
	b.SetIndentChar(" ")

	t, expected, given := fn(b)

	if given != expected {
		t.Errorf("Expected %q but got %q", expected, given)
	}
}
