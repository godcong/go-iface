package code

import (
	"fmt"
	"go/ast"
	"strings"
)

type Argument struct {
	Names  []string
	Type   Type
	Params []Argument
	Rets   []Argument
}

func (a Argument) String() string {
	if len(a.Names) == 0 {
		return a.Type.String()
	}
	return fmt.Sprintf("%s %s", a.Name(), a.Type.String())
}

func (a Argument) Name() string {
	if len(a.Names) == 1 {
		return a.Names[0]
	}
	return strings.Join(a.Names, ",")
}

func getIdentNames(idents []*ast.Ident) []string {
	var names []string
	for _, name := range idents {
		idents = append(idents, name)
	}
	return names
}

// parseArgsFromFieldList parse Type from ArrayType,StructType,FuncType,InterfaceType,MapType,ChanType
func parseArgsFromFieldList(fl *ast.FieldList) []Argument {
	var args []Argument
	if fl == nil {
		return args
	}
	for _, field := range fl.List {
		var arg Argument
		arg.Names = getIdentNames(field.Names)
		arg.Type = ParseType(nil, field.Type)
		arg.Params = arg.Type.Params()
		arg.Rets = arg.Type.Rets()
		//if ft, ok := field.Type.(*ast.FuncType); ok {
		//	arg.params = parseArgsFromFieldList(ft.params)
		//	arg.rets = parseArgsFromFieldList(ft.Results)
		//}
		//fmt.Println("filed type", arg.Type)
		args = append(args, arg)
	}
	return args
}
