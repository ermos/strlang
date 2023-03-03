# strlang
> 🏗️ fluent string builder interface for generate code easily with go

`strlang` is a Go package for building string-based programming languages.
It provides a simple way to generate strings of code that can be used
in different programming languages.

[![Go Reference](https://pkg.go.dev/badge/github.com/ermos/strlang.svg)](https://pkg.go.dev/github.com/ermos/strlang)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ermos/strlang?label=version)](https://github.com/ermos/strlang/blob/main/go.mod)
[![Latest tag](https://img.shields.io/github/v/tag/ermos/strlang)](https://github.com/ermos/strlang/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermos/strlang)](https://goreportcard.com/report/github.com/ermos/strlang)
[![Go Coverage](https://github.com/ermos/strlang/wiki/coverage.svg)](https://raw.githack.com/wiki/ermos/strlang/coverage.html)

## 🛠️ Installation

To use `strlang` in your Go project, you need to have Go installed on your system.
Once you have Go installed, you can use the following command to install the package:

```go
go get -u github.com/ermos/strlang
```

## 📚 Examples
Here are some examples of how to use `strlang` to generate code in different programming languages.
If your wanted programming language isn't implemented, you can use directly the default builder.

### Javascript
Here is an example of how to use Strlang to generate JavaScript code:

#### Code
```go
package main

import (
	"github.com/ermos/strlang"
	"fmt"
)

func main() {
	js := strlang.NewJavascript()

	js.Object("const", "person", func() {
		js.WriteStringln(`name: "John",`)
		js.WriteStringln(`age: 31,`)
		js.WriteStringln(`city: "New York"`)
	})

	js.If("person.age > 18", func() {
		js.WriteStringln(`console.log("John is an adult");`)
	})
	js.Else(func() {
		js.WriteStringln("console.log('John is not an adult');")
	})
	
	fmt.Println(js.String())
}
```

#### Output
```javascript
const person = {
    name: "John",
    age: 31,
    city: "New York"
};

if (person.age > 18) {
    console.log("John is an adult");
} else {
    console.log("John is not an adult");
}
```

### PHP
Here is an example of how to use Strlang to generate PHP code:

#### Code
```go
package main

import (
	"github.com/ermos/strlang"
	"fmt"
)

func main() {
	php := strlang.NewPHP("App/Models")

	php.Class("Person", func() {
		php.ClassFunc("public", "__construct", "$name, $age", "", func() {
			php.WriteStringln(`$this->name = $name;`)
			php.WriteStringln(`$this->age = $age;`)
		})
	})

	fmt.Println(php.String())
}
```

#### Output
```php
<?php

namespace MyNamespace;

class Person {
    public function __construct($name, $age) {
        $this->name = $name;
        $this->age = $age;
    }
}
```

## 💡 Usage

`strlang` provides a set of builders for each programming language.
Each builder has methods that allow you to build code constructs in that language.
You can use the `New` method of your wanted language (example for javascript: `NewJavascript()`)
functions to create instances of the corresponding builders.
Once you have a builder instance, you can use its methods to build code constructs.

For more information about the available methods and how to use them, please refer to the package documentation.

## 🤝 Contributing

Contributions to `strlang` are always welcome!
If you find a bug or have a feature request, please open an issue on GitHub.
If you want to contribute code, please fork the repository and submit a pull request.