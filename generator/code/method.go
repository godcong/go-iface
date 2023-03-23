package code

import (
	"fmt"
	"go/ast"
	"strings"
)

type Method struct {
	Name   string
	Params []Argument
	Rets   []Argument
}

func (m *Method) String() string {
	if m.Name == "" {
		return fmt.Sprintf("(%s) %s", m.ParamString(), m.RetString())
	}
	return fmt.Sprintf("%s(%s) %s", m.Name, m.ParamString(), m.RetString())
}

func (m *Method) ParamString() string {
	return combineArgs(m.Params)
}

func (m *Method) RetString() string {
	if len(m.Rets) == 0 {
		return ""
	}
	if len(m.Rets) == 1 && len(m.Rets[0].Names) == 0 {
		return m.Rets[0].String()
	}
	return combineArgs(m.Rets)
}

func (m *Method) ParseParam(params *ast.FieldList) {
	m.Params = parseArgsFromFieldList(params)
}

func (m *Method) ParseRet(results *ast.FieldList) {
	m.Rets = parseArgsFromFieldList(results)
}

func (m *Method) ParseType(expr ast.Expr) {
	if ft, ok := expr.(*ast.FuncType); ok {
		m.ParseFuncType(ft)
	}
}

func (m *Method) ParseFuncType(ft *ast.FuncType) {
	if ft.Params != nil {
		m.ParseParam(ft.Params)
	}
	if ft.Results != nil {
		m.ParseRet(ft.Results)
	}
}

func combineArgs(args []Argument) string {
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
