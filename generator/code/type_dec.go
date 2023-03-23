package code

import (
	"fmt"
	"go/ast"
	"strings"
)

type funcDec struct {
	*ast.FuncType
	m Method
}

func (f funcDec) String() string {
	return f.m.String()
}

func newFuncDec(v *ast.FuncType) *funcDec {
	fd := &funcDec{FuncType: v}
	if v.Params != nil {
		fd.m.ParseParam(v.Params)
	}
	if v.Results != nil {
		fd.m.ParseRet(v.Results)
	}
	return fd
}

type structDec struct {
	*ast.StructType
	s Struct
}

func (s structDec) String() string {
	return s.s.String()
}

func newStructDec(v *ast.StructType) *structDec {
	sd := &structDec{StructType: v}
	if v.Fields != nil {
		sd.s.ParseVariables(v.Fields)
	}
	return sd
}

type arrayDec struct {
	*ast.ArrayType
	t Type
}

func (a arrayDec) String() string {
	return a.t.String()
}

func newArrayDec(v *ast.ArrayType) *arrayDec {
	ad := &arrayDec{ArrayType: v}
	ad.t = parseFieldType(v.Elt)
	return ad
}

type interfaceDec struct {
	*ast.InterfaceType
	Methods []Method
	t       Type
}

func (i interfaceDec) String() string {
	var methods []string
	for _, m := range i.Methods {
		methods = append(methods, m.String())
	}
	return "interface {" +
		strings.Join(methods, "\n") +
		"}"
}

func newInterfaceDec(v *ast.InterfaceType) *interfaceDec {
	it := &interfaceDec{InterfaceType: v}
	fmt.Println("interface decode", v.Methods)
	if v.Methods != nil {
		for _, m := range v.Methods.List {
			names := getIdentName(m.Names)
			t := parseFieldType(m.Type)
			if t.InType() == "func" {
				mm := Method{}
				if len(names) > 0 {
					mm.Name = names[0]
				}
				mm.ParseType(m.Type)
				fmt.Println("interface method", names, t.String(), t.InType())
				it.Methods = append(it.Methods, mm)
			}

		}
	}

	return it
}

type mapDec struct {
	*ast.MapType
}

func (m mapDec) String() string {
	//TODO implement me
	panic("implement me")
}

func newMapDec(v *ast.MapType) *mapDec {
	return &mapDec{MapType: v}
}
