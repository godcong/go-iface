package generator

import (
	"fmt"
	"go/ast"

	"github.com/godcong/go-inter/generator/code"
)

func parseMethod(s *code.Struct, n *ast.FuncDecl) *code.Struct {
	m := code.Method{
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
	return s
}
