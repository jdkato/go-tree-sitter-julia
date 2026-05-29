# go-tree-sitter-julia

Go bindings for the [tree-sitter-julia](https://github.com/tree-sitter/tree-sitter-julia) parser, powered by [go-tree-sitter](https://github.com/smacker/go-tree-sitter).

## Installation

```
go get github.com/jdkato/go-tree-sitter-julia
```

## Usage

```go
package main

import (
	"fmt"
	julia "github.com/jdkato/go-tree-sitter-julia/julia"
	sitter "github.com/smacker/go-tree-sitter"
)

func main() {
	lang := julia.GetLanguage()
	parser := sitter.NewParser()
	parser.SetLanguage(lang)
	// Use parser...
	fmt.Println("Julia parser loaded.")
}
```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
