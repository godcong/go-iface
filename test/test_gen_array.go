package tst

type arrayGen struct {
}

func (arrayGen) ArrayIntRet([]int) []int {
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
