package code

import (
	"fmt"
	"go/ast"
)

type Struct struct {
	Name      string
	Variables []Argument
	Methods   []Method
}

func (s *Struct) ParseVariables(st *ast.FieldList) {
	s.Variables = parseArgsFromFieldList(st)
}

func (s *Struct) parseMethod(n *ast.FuncDecl) {
	m := Method{
		Name: n.Name.Name,
	}
	fmt.Printf("struct(%s) func %s()\n", s, m.Name)
	if n.Type.Params != nil {
		m.ParseParam(n.Type.Params)
	}

	if n.Type.Results != nil {
		m.ParseRet(n.Type.Results)
	}

	s.Methods = append(s.Methods, m)
}

func (s *Struct) String() string {
	return ""
}
