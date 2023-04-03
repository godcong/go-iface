package parse

import (
	"fmt"
	"go/ast"
	"strings"
)

type Argument struct {
	Names []string
	Type  Type
}

// String returns the string representation of the argument.
func (a Argument) String() string {
	if len(a.Names) == 0 {
		return a.Type.String()
	}
	return fmt.Sprintf("%s %s", a.NameString(), a.Type.String())
}

func (a Argument) NameString() string {
	if len(a.Names) == 1 {
		return a.Names[0]
	}
	return strings.Join(a.Names, ",")
}

func IdentNames(idents []*ast.Ident) []string {
	var names []string
	for _, name := range idents {
		names = append(names, name.Name)
	}
	return names
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

// parseArgsFromFieldList parse Type from ArrayType,StructType,FuncType,InterfaceType,MapType,ChanType
func parseArgsFromFieldList(fl *ast.FieldList) []*Argument {
	var args []*Argument
	if fl == nil {
		return args
	}
	for _, field := range fl.List {
		args = append(args, argFromField(field))
	}
	return args
}

func argFromField(field *ast.Field) *Argument {
	return &Argument{
		Names: IdentNames(field.Names),
		Type:  Parse(field.Type),
	}
}

func combineArgs(args []*Argument) string {
	if len(args) == 0 {
		return ""
	}
	if len(args) == 1 {
		return args[0].String()
	}
	var rets []string
	for i := range args {
		rets = append(rets, args[i].String())
	}
	return strings.Join(rets, ", ")
}
