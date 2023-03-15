package test

// INTER(Name)
// @PATH:
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

// INTER(Namae)
// @PATH:
type namae struct {
}

func (namae) A() {

}

func (namae) B(i int) int {
	return 0
}

func (namae) C(s string) string {
	return ""
}

func (namae) D(i int, s string) string {
	return ""
}

func (namae) E(fn func(string) string) (f func(string) string) {
	return fn
}
