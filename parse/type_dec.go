package parse

import (
	"fmt"
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
	symbol string
	t      Type
}

func (a arrayDec) Val() string {
	return a.symbol + a.t.String()
}

func newArrayDec(v *ast.ArrayType) *arrayDec {
	ad := &arrayDec{ArrayType: v}
	ad.symbol = "[]"
	if v.Len != nil {
		switch l := v.Len.(type) {
		case *ast.BasicLit:
			ad.symbol = "[" + l.Value + "]"
		case *ast.Ellipsis:
			ad.symbol = "[...]"
		}
	}
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
	return "interface {\n" +
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

type chanDec struct {
	*ast.ChanType
	Value Type
}

func (c chanDec) Val() string {
	switch c.Dir {
	case ast.RECV:
		return "<-chan " + c.Value.String()
	case ast.SEND:
		return "chan<- " + c.Value.String()
	default:
		return "chan " + c.Value.String()
	}
}

func newChanDec(v *ast.ChanType) *chanDec {
	cd := &chanDec{ChanType: v}
	cd.Value = Parse(v.Value)
	return cd
}

type ellipsisDec struct {
	*ast.Ellipsis
	Value Type
}

func (c ellipsisDec) Val() string {
	return "..." + c.Value.String()
}

func newEllipsisDec(v *ast.Ellipsis) *ellipsisDec {
	ed := &ellipsisDec{Ellipsis: v}
	ed.Value = Parse(v.Elt)
	return ed
}

type defaultDec struct {
	Node    ast.Node
	isStart bool
	t       string
	p       Type
}

func (d defaultDec) Val() string {
	if d.isStart {
		return "*" + d.t
	}
	return d.t
}

func newDefaultDec(node ast.Node) *defaultDec {
	dd := &defaultDec{Node: node}
	switch n := node.(type) {
	case *ast.StarExpr:
		dd.isStart = true
		dd.p = Parse(n.X)
		dd.t = dd.p.String()
	default:
		dd.t = fmt.Sprintf("%s", n)
	}
	return dd
}
