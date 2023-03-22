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
		m := code.Method{
			Name: n.Name.Name,
		}
		fmt.Printf("struct(%s) func %s()\n", s, m.Name)
		if n.Type.Params != nil {
			for _, field := range n.Type.Params.List {
				var arg code.Argument
				arg.Name = getIdentName(field.Names)
				arg.Type = code.ParseType(field.Type)
				arg.Params = arg.Type.Params()
				arg.Rets = arg.Type.Rets()
				//if ft, ok := field.Type.(*ast.FuncType); ok {
				//	arg.Params = parseArgsFromFieldList(ft.Params)
				//	arg.Rets = parseArgsFromFieldList(ft.Results)
				//}
				fmt.Sprintf("Arg(%s)\n", arg.Type)
				m.Params = append(m.Params, arg)
			}
		}
		if n.Type.Results != nil {
			for _, field := range n.Type.Results.List {
				var arg code.Argument
				arg.Name = getIdentName(field.Names)
				arg.Type = code.ParseType(field.Type)
				arg.Params = arg.Type.Params()
				arg.Rets = arg.Type.Rets()
				//if ft, ok := field.Type.(*ast.FuncType); ok {
				//	arg.Params = parseArgsFromFieldList(ft.Params)
				//	arg.Rets = parseArgsFromFieldList(ft.Results)
				//}
				fmt.Printf("Ret(%s)\n", arg.Type)
				m.Rets = append(m.Rets, arg)
			}
		}
		//if method, ok := inter.Methods[m.name]; ok {
		//	fmt.Println("method", method.name, "is already exist")
		//}
		//if inter.Methods == nil {
		//	inter.Methods = make(map[string]Method)
		//}
		inter.Methods = append(inter.Methods, m)
		v.Interfaces[s] = inter
	}
	return v
}

func getIdentName(names []*ast.Ident) string {
	for _, name := range names {
		return name.Name
	}
	return ""
}

func NewVisitor() *Visitor {
	return &Visitor{
		Interfaces: make(map[string]*code.Struct),
	}
}
