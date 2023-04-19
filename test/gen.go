package tst

// IFACE(baseGen)
// @PATH:
type baseGen struct {
}

func (baseGen) NoArgNoRet() {

}

func (baseGen) noArgNoRet() {

}

func (baseGen) IntArgRet(i int, j *int) int {
	return 0
}

func (baseGen) StringArgRet(s string) string {
	return ""
}

func (baseGen) IntString(i int, s string) string {
	return ""
}

func (baseGen) Variable(vars string, vari ...int) {
}

func (baseGen) FuncRet() func(string) string {
	return func(s string) string {
		return s
	}
}

func (baseGen) StructRet() struct{ v int } {
	return struct{ v int }{v: 0}
}

func (baseGen) F() {

}

func (baseGen) G(i int) (int, error) {
	return 0, nil
}

func (baseGen) H(s string) (string, error) {
	return "", nil
}
