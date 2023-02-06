package contest

type extMutex struct {
	c chan struct{}
}

func New() Mutex {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return &extMutex{c: ch}
}

func (nM *extMutex) LockChannel() <-chan struct{} {
	return nM.c
}

func (nM *extMutex) Lock() {
	<-nM.c
}

func (nM *extMutex) Unlock() {
	nM.c <- struct{}{}
}
