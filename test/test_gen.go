package test

type name struct {
}

func (name) A() {

}

func (name) B(i int) int {
	return 0
}

func (name) C(s string) string {
	return ""
}

func (name) D(i int, s string) string {
	return ""
}

func (name) E() func(string) string {
	return func(s string) string {
		return s
	}
}
