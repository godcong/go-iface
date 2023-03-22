package code

import (
	"fmt"
	"go/ast"
	"strings"
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

type Type struct {
	source          ast.Expr
	names           []string
	inType          string
	inTypeFormat    string
	inTypeParamList *ast.FieldList
	inTypeRetList   *ast.FieldList
	inTypeFunc      func(format string, param, ret *ast.FieldList) string
}

func (t Type) String() string {
	if t.inType == "default" {
		return t.inTypeFormat
	}
	return t.inType
}

func (t Type) Params() []Argument {
	return parseArgsFromFieldList(t.inTypeParamList)

}

func (t Type) Rets() []Argument {
	return parseArgsFromFieldList(t.inTypeRetList)
}

func ParseType(names []*ast.Ident, expr ast.Expr) Type {
	t := Type{
		source:       expr,
		names:        getIdentNames(names),
		inType:       "default",
		inTypeFormat: fmt.Sprintf("%s", expr),
		inTypeFunc:   parseDefaultTypeString,
	}
	switch v := expr.(type) {
	case *ast.FuncType:
		t.inType = "func"
		t.inTypeParamList = v.Params
		t.inTypeRetList = v.Results
		t.inTypeFormat = funcTypeFormat
		t.inTypeFunc = parseFuncTypeString
	case *ast.StructType:
		t.inType = "struct"
		t.inTypeParamList = v.Fields
		t.inTypeFormat = structTypeFormat
		//t.inTypeFunc = parseFuncTypeString
	default:
		//do nothing
	}
	return t
}

func parseFuncTypeString(format string, param, ret *ast.FieldList) string {
	var fields []string
	for _, field := range param.List {
		fields = append(fields, fmt.Sprintf("%s", field.Type))
	}
	return strings.Join(fields, ", ")
}

func parseDefaultTypeString(format string, param, ret *ast.FieldList) string {
	return format
}