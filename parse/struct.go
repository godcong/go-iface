package parse

import (
	"go/ast"
	"strings"
)

type Struct struct {
	Name      string
	Variables []*Argument
	Methods   []*Method
	//OverLoad  []*Method
}

func (s *Struct) parseFunc(n *ast.FuncDecl) {
	m := Method{
		Names: []string{n.Name.Name},
	}
	m.Parse(n.Type)
	s.Methods = append(s.Methods, &m)
}

func (s *Struct) parseVariables(n *ast.StructType) {
	if n.Fields != nil {
		s.Variables = FuncArgs(n.Fields)
	}
}

func (s *Struct) Parse(v ast.Node) {
	switch t := v.(type) {
	case *ast.StructType:
		s.parseVariables(t)
	case *ast.FuncDecl:
		s.parseFunc(t)
		s.parseDoc(t)
	}
}

func (s *Struct) String() string {
	//return a struct as string
	if s.Name == "" {
		return " struct {" +
			combineArgs(s.Variables) +
			"}"
	}

	str := "type " + s.Name + " struct {" +
		combineArgs(s.Variables) +
		"}"
	if len(s.Methods) == 0 {
		return str
	}

	var methods string
	for i := range s.Methods {
		methods += "(" + s.Name + ") " + s.Methods[i].String() + "\n"
	}
	return str + "\n" + methods
}

func (s *Struct) parseDoc(t *ast.FuncDecl) {
	if t.Doc != nil {
		for _, c := range t.Doc.List {
			if strings.HasPrefix(c.Text, "//") {
				comment := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
				if strings.HasPrefix(comment, "OVERLOAD(") && strings.HasSuffix(comment, ")") {
					name := strings.TrimSuffix(strings.TrimPrefix(comment, "OVERLOAD("), ")")
					s.Methods[len(s.Methods)-1].Names = append(s.Methods[len(s.Methods)-1].Names, name)
				}
			}
		}
	}
}
