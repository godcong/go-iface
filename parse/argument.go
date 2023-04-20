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
		return a.Type.Val()
	}
	return fmt.Sprintf("%s %s", CombineNames(a.Names), a.Type.Val())
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
