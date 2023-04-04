package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"text/template"
	"unicode"

	"github.com/godcong/go-iface/generator/parse"
)

// Generator is responsible for generating validation files for the given in a go source file.
type Generator struct {
	fileSet *token.FileSet
	tmpl    *template.Template
	faces   map[string]*parse.Struct
	target  string
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator() *Generator {
	return &Generator{
		tmpl:    addEmbeddedTemplates(template.New("generator")),
		fileSet: token.NewFileSet(),
		faces:   make(map[string]*parse.Struct),
		target:  "",
	}
}

// GenerateFromPath is responsible for orchestrating the Code generation.  It results in a byte array
// that can be written to any file desired.  It has already had goimports run on the code before being returned.
func (g *Generator) GenerateFromPath(path string) (map[string][]byte, error) {
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

func (g *Generator) Generate(f map[string]*ast.Package) (map[string][]byte, error) {
	for name, pkg := range f {
		if g.target == "" {
			g.target = pkg.Name
		}
		fmt.Println("name:", name, ", package", pkg.Name)
		for fname, file := range pkg.Files {
			fmt.Println("filename:", fname)
			ast.Walk(g, file)
		}
	}

	vBuff := bytes.NewBuffer([]byte{})
	err := g.tmpl.ExecuteTemplate(vBuff, "header", map[string]interface{}{
		"version":   "",
		"revision":  "",
		"buildDate": "",
		"builtBy":   "",
	})
	if err != nil {
		return nil, fmt.Errorf("failed write template:%v", err)
	}
	fmt.Println("buffer:", vBuff.String())
	for _, m := range g.faces {
		buf := bytes.NewBuffer(nil)
		buf.WriteString(fmt.Sprintf("type %s interface {\n", camelCase(m.Name)))
		for _, param := range m.Methods {
			buf.WriteString(param.String() + "\n")
		}
		buf.WriteString("}")
		buf.WriteString("\n")
		fmt.Println(buf.String())
	}

	return nil, nil
}

func camelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	cc := []rune(s)
	cc[0] = unicode.ToUpper(cc[0])
	return string(cc)
}

func (g *Generator) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		var s string
		if n.Recv != nil {
			for _, arg := range n.Recv.List {
				s = fmt.Sprintf("%s", arg.Type)
			}
		}
		//skip empty receiver
		if s == "" {
			return g
		}
		inter := &parse.Struct{Name: s}
		if i, ok := g.faces[s]; ok {
			inter = i
		}
		m := parseStructMethod(n)
		if m != nil {
			inter.Methods = append(inter.Methods, m)
		}
		g.faces[s] = inter
	}
	return g
}
