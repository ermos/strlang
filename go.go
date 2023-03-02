package strlang

import (
	"fmt"
	"strings"
)

// Golang is a struct that represents a Go code generator.
type Golang struct {
	packageName string
	imports     []string
	*Builder
}

// NewGolang returns a new instance of Golang code builder.
func NewGolang(packageName string) *Golang {
	return &Golang{
		packageName,
		[]string{},
		NewBuilder(),
	}
}

// If generates an if statement in the generated code.
func (b *Golang) If(statement string, inside func(), ln ...int) {
	b.Block(fmt.Sprintf("if %s {", statement), inside, "}", ln...)
}

// Else generates an else statement in the generated code.
func (b *Golang) Else(inside func(), ln ...int) {
	b.TrimRight("\n")
	b.Block(" else {", inside, "}", ln...)
}

// ElseIf generates an else if statement in the generated code.
func (b *Golang) ElseIf(statement string, inside func(), ln ...int) {
	b.TrimRight("\n")
	b.Block(fmt.Sprintf(" else if %s {", statement), inside, "}", ln...)
}

// AddImports adds import statements to the generated code.
func (b *Golang) AddImports(imports ...string) {
	b.imports = append(b.imports, imports...)
}

// Func generates a function definition in the generated code.
func (b *Golang) Func(fromStruct, name, parameters, output string, inside func()) {
	if fromStruct != "" {
		fromStruct = fmt.Sprintf("(%s) ", fromStruct)
	}

	if output != "" {
		output = fmt.Sprintf("(%s) ", output)
	}

	b.Block(
		fmt.Sprintf(
			"func %s%s(%s) %s{",
			fromStruct,
			name,
			parameters,
			output,
		),
		inside,
		"}",
		2,
	)
}

// Struct generates a struct definition in the generated code.
func (b *Golang) Struct(name string, inside func()) {
	b.Block(fmt.Sprintf("type %s struct {", name), inside, "}", 2)
}

// StructField generates a struct field definition in the generated code.
func (b *Golang) StructField(name, goType string, docs ...map[string]string) {
	strDocs := ""
	if len(docs) != 0 {
		doc := docs[0]
		var arrDocs []string
		for key, value := range doc {
			arrDocs = append(arrDocs, fmt.Sprintf(`%s:"%s"`, key, value))
		}
		strDocs = fmt.Sprintf(" `%s`", strings.Join(arrDocs, " "))
	}

	b.WriteStringln(fmt.Sprintf("%s %s%s", name, goType, strDocs))
}

// String returns the generated code as a string.
func (b *Golang) String() string {
	sb := NewBuilder()

	sb.WriteStringln(fmt.Sprintf("package %s", b.packageName), 2)

	if len(b.imports) != 0 {
		var imports []string
		alreadyExists := make(map[string]bool)

		for _, i := range b.imports {
			if !alreadyExists[i] {
				imports = append(imports, i)
				alreadyExists[i] = true
			}
		}

		sb.Block("import (", func() {
			for _, i := range imports {
				sb.WriteStringln(fmt.Sprintf(`"%s"`, i))
			}
		}, ")", 2)
	}

	sb.WriteString(b.builder.String())

	return sb.String()
}
