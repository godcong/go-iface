package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"text/template"
)

// Generator is responsible for generating validation files for the given in a go source file.
type Generator struct {
	fileSet *token.FileSet
	tmpl    *template.Template
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator() *Generator {
	return &Generator{
		tmpl:    template.New("generator"),
		fileSet: token.NewFileSet(),
	}
}

// GenerateFromPath is responsible for orchestrating the Code generation.  It results in a byte array
// that can be written to any file desired.  It has already had goimports run on the code before being returned.
func (g *Generator) GenerateFromPath(path string) ([]byte, error) {
	f, err := g.parsePath(path)
	if err != nil {
		return nil, fmt.Errorf("generate: error parsing input path '%s': %s", path, err)
	}
	return g.Generate(f)
}

// parsePath simply calls the go/parser ParseFile function with an empty token.FileSet
func (g *Generator) parsePath(fileName string) (map[string]*ast.Package, error) {
	// Parse the file given in arguments
	return parser.ParseDir(g.fileSet, fileName, nil, parser.ParseComments)
}

func (g *Generator) Generate(f map[string]*ast.Package) ([]byte, error) {
	v := NewVisitor()
	v.withName = true
	for name, pkg := range f {
		fmt.Println("name:", name, ", package", pkg.Name)
		for fname, file := range pkg.Files {
			fmt.Println("filename:", fname)
			ast.Walk(v, file)
		}
	}
	for _, m := range v.Interfaces {
		fmt.Printf("Struct type [%s] has\n", m.Name)

		for _, param := range m.Methods {
			fmt.Printf("Function name: %s\n", param.name)
			for _, argument := range param.params {
				fmt.Printf("Arguments: %s\n", argument.String())
			}
			for _, argument := range param.rets {
				fmt.Printf("Return types: %s\n", argument.String())
			}

		}
		fmt.Println()
	}
	return nil, nil
}
