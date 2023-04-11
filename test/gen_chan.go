package tst

type chanGen struct {
}

func (chanGen) ChanRet() chan<- chan<- string {
	return nil
}

func (chanGen) ChanIn(a, b <-chan <-chan int) {

}

// ChanInDefault if you input chan <-chan the imports will optimize to chan<- chan...:)
func (chanGen) ChanInDefault(a, b chan<- chan int) string {
	return ""
}
