package parse

import (
	"go/ast"
	"strings"
)

var fmtOutput = map[string]string{
	"default": "%s",
}

type funcDec struct {
	*ast.FuncType
	m Method
}

func (f funcDec) Val() string {
	return f.m.String()
}

func newFuncDec(v *ast.FuncType) *funcDec {
	fd := &funcDec{FuncType: v}
	fd.m.Parse(v)
	return fd
}

type structDec struct {
	*ast.StructType
	s Struct
}

func (s structDec) Val() string {
	return s.s.String()
}

func newStructDec(v *ast.StructType) *structDec {
	sd := &structDec{StructType: v}
	sd.s.Parse(v)
	return sd
}

type arrayDec struct {
	*ast.ArrayType
	t Type
}

func (a arrayDec) Val() string {
	return "[]" + a.t.String()
}

func newArrayDec(v *ast.ArrayType) *arrayDec {
	ad := &arrayDec{ArrayType: v}
	ad.t = Parse(v.Elt)
	return ad
}

type interfaceDec struct {
	*ast.InterfaceType
	t       Type
	Methods []*Method
}

func (i interfaceDec) Val() string {
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
	if v.Methods != nil {
		for _, method := range v.Methods.List {
			var m Method
			t := Parse(method.Type)
			if t.InType() != "default" {
				m.Parse(method.Type)
				for i := range method.Names {
					m.Names = append(m.Names, method.Names[i].String())
				}
			} else {
				m.Type = t
			}
			it.Methods = append(it.Methods, &m)
		}
	}
	return it
}

type mapDec struct {
	*ast.MapType
	Key   Type
	Value Type
}

func (m mapDec) Val() string {
	return "map[" + m.Key.String() + "]" + m.Value.String()
}

func newMapDec(v *ast.MapType) *mapDec {
	md := &mapDec{MapType: v}
	md.Key = Parse(v.Key)
	md.Value = Parse(v.Value)
	return md
}
