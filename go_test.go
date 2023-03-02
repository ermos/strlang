package strlang

import (
	"fmt"
	"testing"
)

func TestGolang_If(t *testing.T) {
	testGolangInit(func(b *Golang) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false")
		})

		expected := "" +
			"if true {\n" +
			" return false\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestGolang_Else(t *testing.T) {
	testGolangInit(func(b *Golang) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false")
		})
		b.Else(func() {
			b.WriteStringln("return true")
		})

		expected := "" +
			"if true {\n" +
			" return false\n" +
			"} else {\n" +
			" return true\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestGolang_ElseIf(t *testing.T) {
	testGolangInit(func(b *Golang) (*testing.T, string, string) {
		b.If("true", func() {
			b.WriteStringln("return false")
		})
		b.ElseIf("false", func() {
			b.WriteStringln("return true")
		})

		expected := "" +
			"if true {\n" +
			" return false\n" +
			"} else if false {\n" +
			" return true\n" +
			"}\n"

		return t, expected, b.String()
	})
}

func TestGolang_AddImports(t *testing.T) {
	testGolangInit(func(b *Golang) (*testing.T, string, string) {
		b.AddImports("fmt", "strings")

		expected := "" +
			"import (\n" +
			"  \"fmt\"\n" +
			"  \"strings\"\n" +
			")\n\n"

		return t, expected, b.String()
	})
}

func TestGolang_Func(t *testing.T) {
	testGolangInit(func(b *Golang) (*testing.T, string, string) {
		b.Func("", "GetEmail", "id int64", "string", func() {
			b.WriteStringln("return \"john@doe.com\"")
		})

		expected := "" +
			"func GetEmail(id int64) (string) {\n" +
			" return \"john@doe.com\"\n" +
			"}\n\n"

		return t, expected, b.String()
	})
}

func TestGolang_Func_WithStruct(t *testing.T) {
	testGolangInit(func(b *Golang) (*testing.T, string, string) {
		b.Func("User", "GetEmail", "id int64", "string", func() {
			b.WriteStringln("return \"john@doe.com\"")
		})

		expected := "" +
			"func (User) GetEmail(id int64) (string) {\n" +
			" return \"john@doe.com\"\n" +
			"}\n\n"

		return t, expected, b.String()
	})
}

func TestGolang_Struct(t *testing.T) {
	testGolangInit(func(b *Golang) (*testing.T, string, string) {
		b.Struct("User", func() {
			b.StructField("FirstName", "string", map[string]string{
				"db": "user.first_name",
			})
		})

		expected := "" +
			"type User struct {\n" +
			" FirstName string `db:\"user.first_name\"`\n" +
			"}\n\n"

		return t, expected, b.String()
	})
}

func testGolangInit(fn func(b *Golang) (t *testing.T, expected, given string)) {
	b := NewGolang("test")
	b.SetIndentChar(" ")

	t, expected, given := fn(b)

	expected = "package test\n\n" + expected

	if given != expected {
		fmt.Println(expected)
		fmt.Println(given)
		t.Errorf("Expected %q but got %q", expected, given)
	}
}
