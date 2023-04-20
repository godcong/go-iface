package parse

import (
	"go/ast"
	"strings"
)

type Parser interface {
	Val() string
}

func Parse(node ast.Node) Type {
	var t Type

	switch v := node.(type) {
	case *ast.FuncType:
		t.typeStr = "FuncType"
		t.parser = newFuncDec(v)
	case *ast.StructType:
		t.typeStr = "StructType"
		t.parser = newStructDec(v)
	case *ast.ArrayType:
		t.typeStr = "ArrayType"
		t.parser = newArrayDec(v)
	case *ast.InterfaceType:
		t.typeStr = "InterfaceType"
		t.parser = newInterfaceDec(v)
	case *ast.MapType:
		t.typeStr = "MapType"
		t.parser = newMapDec(v)
	case *ast.ChanType:
		t.typeStr = "ChanType"
		t.parser = newChanDec(v)
	case *ast.ParenExpr:
		t.typeStr = "ParenExpr"
		t.parser = newParenDec(v)
	case *ast.Ellipsis:
		t.typeStr = "Ellipsis"
		t.parser = newEllipsisDec(v)
	default:
		t.typeStr = "Default"
		t.parser = newDefaultDec(node)
	}
	t.typeVal = t.parser.Val()
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
