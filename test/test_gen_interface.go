package test

type b interface {
	B() string
}

type interfaceGen struct {
}

func (interfaceGen) IfRet() interface {
	b
	A() string
} {
	return interfaceGen{}
}

func (interfaceGen) IfParam(a, b interface {
	b
	A() string
}) {
}

func (interfaceGen) B() string {
	return ""
}

func (interfaceGen) A() string {
	return ""
}
