package code

import (
	"fmt"
	"go/ast"
)

const (
	funcTypeFormat   = "func(%s) %s"
	structTypeFormat = `struct {
	%s %s
}`
	interfaceTypeFormat = `interface {
	%s %s
}`
)

type TypeStringer interface {
	String() string
}

type Type struct {
	t      TypeStringer
	name   string
	inType string
}

func (t Type) InType() string {
	return t.inType
}

func (t Type) String() string {
	if t.inType == "default" {
		return t.name
	}
	return t.t.String()
}

func parseFieldType(expr ast.Expr) Type {
	t := Type{
		name:   "",
		inType: "default",
	}
	switch v := expr.(type) {
	case *ast.FuncType:
		t.inType = "func"
		t.t = newFuncDec(v)
	case *ast.StructType:
		t.inType = "struct"
		t.t = newStructDec(v)
	case *ast.ArrayType:
		t.inType = "struct"
		t.t = newArrayDec(v)
	case *ast.InterfaceType:
		t.inType = "interface"
		t.t = newInterfaceDec(v)
	case *ast.MapType:
		t.inType = "map"
		t.t = newMapDec(v)
	default:
		t.name = fmt.Sprintf("%s", expr)
	}
	return t
}
