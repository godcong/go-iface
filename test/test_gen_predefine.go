package test

type PreDefineMethod func(int) string
type PreDefineString string
type PreDefineInterface any
type PreDefineStruct struct {
	Any any
}

type preDefineGen struct {
}

func (preDefineGen) PreMethodArgRet(s PreDefineMethod) PreDefineMethod {
	return func(i int) string {
		return ""
	}
}

func (preDefineGen) PreStringArgRet(m PreDefineString) (s PreDefineString) {
	return ""
}

func (preDefineGen) PreInterfaceArgRet(m PreDefineInterface) (s PreDefineInterface) {
	return ""
}

func (preDefineGen) PreStructArgRet(m PreDefineStruct) (s PreDefineStruct) {
	return PreDefineStruct{}
}

func (preDefineGen) PreMultiArgRet(m PreDefineMethod, s PreDefineString, st PreDefineStruct) (
	PreDefineStruct, PreDefineString, PreDefineMethod) {
	return st, s, m
}

func (preDefineGen) PreMultiArgRetErr(a, b, c PreDefineMethod, s PreDefineString, st PreDefineStruct) (
	x, y, z PreDefineStruct, sr struct{ r string }, e error) {
	return x, y, z, struct{ r string }{r: ""}, nil
}

func (preDefineGen) PreMultiArgRetErr2(m, n, o, p PreDefineMethod) (s, t, r struct{ r string }, e error) {
	return s, t, r, nil
}
