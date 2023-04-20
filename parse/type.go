package parse

type Type struct {
	parser  Parser
	typeVal string
	typeStr string
}

func (t Type) TypeStr() string {
	return t.typeStr
}

func (t Type) Val() string {
	return t.typeVal
}
