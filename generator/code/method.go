package code

import (
	"fmt"
	"strings"
)

type Method struct {
	name   string
	params []Argument
	rets   []Argument
}

func (m Method) String() string {
	return fmt.Sprintf("%s(%s) %s", m.name, m.ParamString(), m.RetString())
}

func (m Method) ParamString() string {
	return combineArgs(m.rets)
}

func (m Method) methodFormat() string {
	f := "%s"
	if len(m.params) == 0 {
		f = f + "()"
	}
	return f
}

func (m Method) RetString() string {
	if len(m.rets) == 0 {
		return ""
	}
	return combineArgs(m.rets)
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
