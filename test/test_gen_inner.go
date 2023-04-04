package test

type innerInterface struct {
}

func (innerInterface) Put(string2 string) {

}

// IFACE(innerGen)
// @PATH[]
type innerGen struct {
}

func (innerGen) InnerMethod(func(string2 string) string) func(int) int {
	return func(i int) int {
		return i
	}
}

func (innerGen) InnerStruct(i struct{ Any any }) struct{ Int int } {
	return struct{ Int int }{}
}

func (innerGen) InnerInterface(s interface{ Get() int }) interface{ Put(string2 string) } {
	return innerInterface{}
}

func (innerGen) D(i int, s string) string {
	return ""
}

func (innerGen) E(fn func(string) string) (f func(string) string) {
	return fn
}
