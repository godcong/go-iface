package parse

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

type decoder[T any] struct {
	s T
	t Type
}

func newDecoder[T any](v T) *decoder[T] {
	log.Debug("AnyType", "type", reflect.TypeOf(v))
	gd := &decoder[T]{s: v}
	return gd
}

type funcDec struct {
	*ast.FuncType
	m Method
}

func (f funcDec) Val() string {
	return f.m.String()
}

func newFuncDec(v *ast.FuncType) *funcDec {
	log.Debug("FuncType", "type", v)
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
	log.Debug("StructType", "type", v)
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
	return a.symbol + a.t.Val()
}

func newArrayDec(v *ast.ArrayType) *arrayDec {
	log.Debug("ArrayType", "type", v)
	ad := &arrayDec{ArrayType: v}
	switch l := v.Len.(type) {
	case *ast.BasicLit:
		ad.symbol = "[" + l.Value + "]"
	case *ast.Ellipsis:
		ad.symbol = "[...]"
	default:
		ad.symbol = "[]"
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
	log.Debug("InterfaceType", "type", v)
	it := &interfaceDec{InterfaceType: v}
	if v.Methods != nil {
		for _, method := range v.Methods.List {
			var m Method
			t := Parse(method.Type)
			if t.TypeStr() != "default" {
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
	return "map[" + m.Key.Val() + "]" + m.Value.Val()
}

func newMapDec(v *ast.MapType) *mapDec {
	log.Debug("MapType", "type", v)
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
		return "<-chan " + c.Value.Val()
	case ast.SEND:
		return "chan<- " + c.Value.Val()
	default:
		return "chan " + c.Value.Val()
	}
}

func newChanDec(v *ast.ChanType) *chanDec {
	log.Debug("ChanType", "type", v)
	cd := &chanDec{ChanType: v}
	cd.Value = Parse(v.Value)
	return cd
}

type ellipsisDec struct {
	*ast.Ellipsis
	Value Type
}

func (c ellipsisDec) Val() string {
	return "..." + c.Value.Val()
}

func newEllipsisDec(v *ast.Ellipsis) *ellipsisDec {
	log.Debug("Ellipsis", "type", v)
	ed := &ellipsisDec{Ellipsis: v}
	ed.Value = Parse(v.Elt)
	return ed
}

type parenDec struct {
	Paren *ast.ParenExpr
	Value Type
}

func newParenDec(v *ast.ParenExpr) *parenDec {
	ed := &parenDec{Paren: v}
	ed.Value = Parse(v.X)
	return ed
}

func (d parenDec) Val() string {
	return "(" + d.Value.Val() + ")"
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
		log.Debug("StarExpr", "type", n)
		dd.isStart = true
		dd.p = Parse(n.X)
		dd.t = dd.p.Val()
	default:
		log.Debug("Default", "type", n)
		dd.t = fmt.Sprintf("%s", n)
	}

	return dd
}
