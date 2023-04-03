package parse

type Type struct {
	t      string
	p      Parser
	inType string
}

func (t Type) InType() string {
	return t.inType
}

func (t Type) String() string {
	if t.inType == "default" {
		return t.t
	}
	return t.p.Val()
}
