package tst

type arrayGen struct {
}

func (arrayGen) ArrayIntRet([]int) []int {
	return nil
}

func (arrayGen) ArrayIntRetP([]*int) []*int {
	return nil
}

func (arrayGen) ArrayInt5(i [5]int) [5]int {
	return i
}

func (arrayGen) ArrayInt5P(i [5]*int) [5]*int {
	return i
}

func (arrayGen) ArrayIntRetEP() *[...]int {
	return nil
}

//func (arrayGen) ArrayIntRetE() (ret [...]int) {
//	return
//}

//func (arrayGen) ArrayInE(e [...]int) error {
//	return nil
//}

func (arrayGen) ArrayInEP(p *[...]int) error {
	return nil
}

func (arrayGen) ArrayStringIn(a, b []string) []string {
	return a
}

func (arrayGen) ArrayStruct(a, b []struct{ Any, Iface any }) []struct{ Any, Iface any } {
	return a
}

func (arrayGen) ArrayFunc(a, b []func(string2 string) int) (x, y []func(string2 string) int) {
	return a, b
}
