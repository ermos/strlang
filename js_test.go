package strlang

import (
	"fmt"
	"testing"
)

func TestJavascript_If(t *testing.T) {
	testJavascriptInit(func(b *Javascript) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false")
		})

		expected := "" +
			"if (true) {\n" +
			" return false\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestJavascript_Else(t *testing.T) {
	testJavascriptInit(func(b *Javascript) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false")
		})
		b.Else(func() {
			b.WriteStringln("return true")
		})

		expected := "" +
			"if (true) {\n" +
			" return false\n" +
			"} else {\n" +
			" return true\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestJavascript_ElseIf(t *testing.T) {
	testJavascriptInit(func(b *Javascript) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false")
		})
		b.ElseIf("false", func() {
			b.WriteStringln("return true")
		})

		expected := "" +
			"if (true) {\n" +
			" return false\n" +
			"} else if (false) {\n" +
			" return true\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestJavascript_Export(t *testing.T) {
	testJavascriptInit(func(b *Javascript) (*testing.T, string, string) {
		b.Export()

		expected := "export "

		return t, expected, b.String()
	})
}

func TestJavascript_Object(t *testing.T) {
	testJavascriptInit(func(b *Javascript) (*testing.T, string, string) {
		b.Export().Object("const", "Test", func() {
			b.WriteStringln("isTest: true,")
		})

		expected := "" +
			"export const Test = {\n" +
			" isTest: true,\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func testJavascriptInit(fn func(b *Javascript) (t *testing.T, expected, given string)) {
	b := NewJavascript()
	b.SetIndentChar(" ")

	t, expected, given := fn(b)

	if given != expected {
		fmt.Println(expected)
		fmt.Println(given)
		t.Errorf("Expected %q but got %q", expected, given)
	}
}
