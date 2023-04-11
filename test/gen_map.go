package tst

type mapGen struct {
}

func (mapGen) MapRet() map[string]map[string]int {
	return nil
}

func (mapGen) MapRetP() *map[string]*map[string]*int {
	return nil
}

func (mapGen) MapIn(a, b map[int]map[string]any) {

}

func (mapGen) MapInP(a, b *map[int]*map[string]*any) {

}
