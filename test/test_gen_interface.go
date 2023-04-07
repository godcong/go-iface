package tst

type B interface {
	B() string
}

type interfaceGen struct {
}

func (interfaceGen) IfRet() interface {
	B
	A() string
} {
	return interfaceGen{}
}

func (interfaceGen) IfParam(a, b interface {
	B
	A() string
}) {
}

func (interfaceGen) B() string {
	return ""
}

func (interfaceGen) A() string {
	return ""
}
