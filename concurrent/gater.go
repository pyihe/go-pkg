package concurrent

// Gater 用于控制最大并发协程数
type Gater chan struct{}

func NewGater(cap int) Gater {
	return make(Gater, cap)
}

func (g Gater) Enter() {
	g <- struct{}{}
}

func (g Gater) Leave() {
	<-g
}
