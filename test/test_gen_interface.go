package test

type b interface {
	B() string
}

type gi struct {
}

func (gi) IfRet() interface {
	b
	A() string
} {
	return gi{}
}

func (gi) IfParam(a, b interface {
	b
	A() string
}) {
}

func (gi) B() string {
	return ""
}

func (gi) A() string {
	return ""
}
