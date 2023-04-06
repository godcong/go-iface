package parse

import (
	"fmt"
	"go/ast"
	"strings"
)

type Parser interface {
	Val() string
}

func Parse(node ast.Node) Type {
	t := Type{
		inType: "default",
	}
	switch v := node.(type) {
	case *ast.FuncType:
		t.inType = "func"
		t.p = newFuncDec(v)
	case *ast.StructType:
		t.inType = "struct"
		t.p = newStructDec(v)
	case *ast.ArrayType:
		t.inType = "struct"
		t.p = newArrayDec(v)
	case *ast.InterfaceType:
		t.inType = "interface"
		t.p = newInterfaceDec(v)
	case *ast.MapType:
		t.inType = "map"
		t.p = newMapDec(v)
	default:
		t.t = fmt.Sprintf("%s", node)
	}
	return t
}

func IdentNames(idents []*ast.Ident) []string {
	var names []string
	for _, name := range idents {
		names = append(names, name.Name)
	}
	return names
}

func CombineNames(names []string) string {
	switch len(names) {
	case 0:
		return ""
	case 1:
		return names[0]
	default:
		return strings.Join(names, ",")
	}
}

func FuncArgs(params *ast.FieldList) []*Argument {
	if params != nil {
		var args []*Argument
		for _, field := range params.List {
			args = append(args, argFromField(field))
		}
		return args
	}
	return nil
}
