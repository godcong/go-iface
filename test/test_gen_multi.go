package test

type FuncMethod func(int) string

func (name) F() {

}

func (name) G(i int) (int, error) {
	return 0, nil
}

func (name) H(s string) (string, error) {
	return "", nil
}

func (name) I(s FuncMethod) (FuncMethod, error) {
	return func(i int) string {
		return ""
	}, nil
}

func (name) J(m FuncMethod) (f func(int) string, e error) {
	return func(i int) string {
		return ""
	}, nil
}

func (name) K(m FuncMethod) (s struct{ r string }, e error) {
	return struct{ r string }{r: ""}, nil
}

func (name) L(m, n, o, p FuncMethod) (s struct{ r string }, e error) {
	return struct{ r string }{r: ""}, nil
}
