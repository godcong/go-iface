package parse

import (
	"fmt"
	"go/ast"
)

type Method struct {
	Names []string
	Type  Type
	Args  []*Argument
	Ret   []*Argument
}

func (m *Method) String() string {
	if m.Type.InType() == "default" {
		return m.Type.String()
	}
	if len(m.Names) == 0 {
		return fmt.Sprintf("func(%s) %s", combineArgs(m.Args), m.retString())
	}
	return fmt.Sprintf("%s(%s) %s", CombineNames(m.Names), combineArgs(m.Args), m.retString())
}

func (m *Method) retString() string {
	switch len(m.Ret) {
	case 0:
		return ""
	case 1:
		if len(m.Ret[0].Names) == 0 {
			return m.Ret[0].Type.String()
		}
		return fmt.Sprintf("(%s %s)", CombineNames(m.Ret[0].Names), m.Ret[0].Type.String())
	default:
		return "(" + combineArgs(m.Ret) + ")"
	}
}

func (m *Method) Parse(expr ast.Node) {
	if ft, ok := expr.(*ast.FuncType); ok {
		if ft.Params != nil {
			//log.Info("parse args")
			m.Args = FuncArgs(ft.Params)
		}
		if ft.Results != nil {
			//log.Info("parse ret")
			m.Ret = FuncArgs(ft.Results)
		}
	}

}
