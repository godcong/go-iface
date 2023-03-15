package generator

import (
	"fmt"
	"go/ast"
)

type InterfaceInfo struct {
	Name string
	Args []string
	Ret  []string
}

type InterfaceVisitor struct {
	Methods []InterfaceInfo
}

func (v *InterfaceVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		funcInfo := InterfaceInfo{
			Name: n.Name.Name,
		}
		if n.Type.Params != nil {
			for _, arg := range n.Type.Params.List {
				argType := fmt.Sprintf("%s", arg.Type)
				for _, name := range arg.Names {
					funcInfo.Args = append(funcInfo.Args, fmt.Sprintf("%s %s", name.Name, argType))
				}
			}
		}
		if n.Type.Results != nil {
			for _, ret := range n.Type.Results.List {
				if ft, ok := ret.Type.(*ast.FuncType); ok {
					retType := fmt.Sprintf("func(%s) %s", formatFieldList(ft.Params), formatFieldList(ft.Results))
					funcInfo.Ret = append(funcInfo.Ret, retType)
				} else {
					retType := fmt.Sprintf("%s", ret.Type)
					funcInfo.Ret = append(funcInfo.Ret, retType)
				}
			}
		}
		v.Methods = append(v.Methods, funcInfo)
	}
	return v
}

func formatFieldList(fl *ast.FieldList) string {
	var fields []string
	for _, field := range fl.List {
		fields = append(fields, fmt.Sprintf("%s", field.Type))
	}
	return "(" + fmt.Sprintf("%s", fields) + ")"
}
