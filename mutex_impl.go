package contest

import (
	"sync"
)

const mutexLocked = 1

type extMutex struct {
	stMu sync.Mutex
	c    chan struct{}
}

func New() Mutex {
	return &extMutex{c: make(chan struct{}, 1)}
}

func (nM *extMutex) LockChannel() <-chan struct{} {
	if nM.stMu.TryLock() {
		nM.c <- struct{}{}
	}
	return nM.c
}

func (nM *extMutex) Lock() {
	nM.stMu.Lock()
}

func (nM *extMutex) Unlock() {
	nM.stMu.Unlock()
}
