package generator

import (
	"fmt"
	"go/ast"

	"github.com/godcong/go-inter/generator/code"
)

type Visitor struct {
	withName   bool
	Interfaces map[string]*code.Struct
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
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
			return v
		}
		inter := &code.Struct{Name: s}
		if i, ok := v.Interfaces[s]; ok {
			inter = i
		}
		v.Interfaces[s] = parseMethod(inter, n)
	}
	return v
}

func NewVisitor() *Visitor {
	return &Visitor{
		Interfaces: make(map[string]*code.Struct),
	}
}
