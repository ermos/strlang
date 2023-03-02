package strlang

import (
	"fmt"
	"testing"
)

func TestPHP_If(t *testing.T) {
	testPHPInit(func(b *PHP) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false;")
		})

		expected := "" +
			"if (true) {\n" +
			" return false;\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestPHP_Else(t *testing.T) {
	testPHPInit(func(b *PHP) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false;")
		})
		b.Else(func() {
			b.WriteStringln("return true;")
		})

		expected := "" +
			"if (true) {\n" +
			" return false;\n" +
			"} else {\n" +
			" return true;\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestPHP_ElseIf(t *testing.T) {
	testPHPInit(func(b *PHP) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false;")
		})
		b.ElseIf("false", func() {
			b.WriteStringln("return true;")
		})

		expected := "" +
			"if (true) {\n" +
			" return false;\n" +
			"} elseif (false) {\n" +
			" return true;\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestPHP_Class(t *testing.T) {
	testPHPInit(func(b *PHP) (*testing.T, string, string) {
		b.Class("User", func() {
			b.ClassFunc("public", "getEmail", "", "string", func() {
				b.WriteStringln("return $this->email;")
			})
		})

		expected := "" +
			"class User {\n" +
			" public function getEmail(): string {\n" +
			"  return $this->email;\n" +
			" }\n\n" +
			"}\n\n"

		return t, expected, b.String()
	})
}

func testPHPInit(fn func(b *PHP) (t *testing.T, expected, given string)) {
	b := NewPHP("App")
	b.SetIndentChar(" ")

	t, expected, given := fn(b)

	expected = "<?php\n\nnamespace App;\n\n" + expected

	if given != expected {
		fmt.Println(expected)
		fmt.Println(given)
		t.Errorf("Expected %q but got %q", expected, given)
	}
}
