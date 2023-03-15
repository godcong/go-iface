package generator

import (
	"fmt"
	"go/ast"
	"strings"
)

type Visitor struct {
	withName bool
	Structs  map[string][]Interface
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		funcInfo := Interface{
			Name: n.Name.Name,
		}
		var s string
		if n.Recv != nil {
			for _, arg := range n.Recv.List {
				s = fmt.Sprintf("%s", arg.Type)
			}
		}
		if s == "" {
			return v
		}

		if n.Type.Params != nil {
			for _, arg := range n.Type.Params.List {
				argName := v.getName(arg.Names)
				if ft, ok := arg.Type.(*ast.FuncType); ok {
					if argName != "" {
						argType := fmt.Sprintf("%s func(%s) %s", argName, formatFieldList(ft.Params), formatFieldList(ft.Results))
						funcInfo.Params = append(funcInfo.Params, argType)
					} else {
						argType := fmt.Sprintf("func(%s) %s", formatFieldList(ft.Params), formatFieldList(ft.Results))
						funcInfo.Params = append(funcInfo.Params, argType)
					}
				} else {
					argType := fmt.Sprintf("%s", arg.Type)
					if argName != "" {
						funcInfo.Params = append(funcInfo.Params, fmt.Sprintf("%s %s", argName, argType))
					} else {
						funcInfo.Params = append(funcInfo.Params, fmt.Sprintf("%s", argType))
					}

				}
			}
		}
		if n.Type.Results != nil {
			for _, ret := range n.Type.Results.List {
				retName := v.getName(ret.Names)
				if ft, ok := ret.Type.(*ast.FuncType); ok {
					if retName != "" {
						retType := fmt.Sprintf("%s func(%s) %s", retName, formatFieldList(ft.Params), formatFieldList(ft.Results))
						funcInfo.RetTypes = append(funcInfo.RetTypes, retType)
					} else {
						retType := fmt.Sprintf("func(%s) %s", formatFieldList(ft.Params), formatFieldList(ft.Results))
						funcInfo.RetTypes = append(funcInfo.RetTypes, retType)
					}
				} else {
					retType := fmt.Sprintf("%s", ret.Type)
					if retName != "" {
						funcInfo.RetTypes = append(funcInfo.RetTypes, fmt.Sprintf("%s %s", retName, retType))
					} else {
						funcInfo.RetTypes = append(funcInfo.RetTypes, fmt.Sprintf("%s", retType))
					}

				}
			}
		}
		v.Structs[s] = append(v.Structs[s], funcInfo)
	}
	return v
}

func (v *Visitor) getName(names []*ast.Ident) string {
	if v.withName {
		for _, name := range names {
			return name.Name
		}
	}
	return ""
}

func formatFieldList(fl *ast.FieldList) string {
	var fields []string
	for _, field := range fl.List {
		fields = append(fields, fmt.Sprintf("%s", field.Type))
	}
	return strings.Join(fields, ", ")
}

func NewVisitor() *Visitor {
	return &Visitor{
		Structs: make(map[string][]Interface),
	}
}
