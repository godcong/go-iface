package generator

import (
	"fmt"
	"go/ast"

	"github.com/godcong/go-iface/generator/parse"
)

type Interface struct {
	i        map[string]*parse.Struct
	withName bool
	target   string
	pkg      string
}

func (v *Interface) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		fmt.Println("name:", n.Name.Name, ", package", n.Name.Name)
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
		inter := &parse.Struct{Name: s}
		if i, ok := v.i[s]; ok {
			inter = i
		}
		m := parseStructMethod(n)
		if m != nil {
			inter.Methods = append(inter.Methods, m)
		}
		v.i[s] = inter
	}
	return v
}

func NewVisitor() *Interface {
	return &Interface{
		i: make(map[string]*parse.Struct),
	}
}

func parseStructMethod(n *ast.FuncDecl) *parse.Method {
	return &parse.Method{
		Name: n.Name.Name,
		Args: parse.FuncArgs(n.Type.Params),
		Ret:  parse.FuncArgs(n.Type.Results),
	}
}
