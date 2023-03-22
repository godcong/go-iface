package code

import (
	"fmt"
	"go/ast"
)

type Argument struct {
	Name   string
	Type   Type
	Params []Argument
	Rets   []Argument
}

func (a Argument) String() string {
	fmt.Println("type", a.Type)
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

func getIdentName(names []*ast.Ident) string {
	for _, name := range names {
		return name.Name
	}
	return ""
}

// parseArgsFromFieldList parse Type from ArrayType,StructType,FuncType,InterfaceType,MapType,ChanType
func parseArgsFromFieldList(fl *ast.FieldList) []Argument {
	var args []Argument
	if fl == nil {
		return args
	}
	for _, field := range fl.List {
		var arg Argument
		arg.Name = getIdentName(field.Names)
		arg.Type = ParseType(field.Type)
		arg.Params = arg.Type.Params()
		arg.Rets = arg.Type.Rets()
		if ft, ok := field.Type.(*ast.FuncType); ok {

			arg.Params = parseArgsFromFieldList(ft.Params)
			arg.Rets = parseArgsFromFieldList(ft.Results)
		}
		fmt.Println("filed type", arg.Type)
		args = append(args, arg)
	}
	return args
}
