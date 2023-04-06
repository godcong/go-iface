package parse

import (
	"go/ast"
)

type Struct struct {
	Name      string
	Variables []*Argument
	Methods   []*Method
}

func (s *Struct) parseFunc(n *ast.FuncDecl) {
	m := Method{
		Name: n.Name.Name,
	}
	if n.Type.Params != nil {
		m.Args = FuncArgs(n.Type.Params)
	}

	if n.Type.Results != nil {
		m.Ret = FuncArgs(n.Type.Results)
	}
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
