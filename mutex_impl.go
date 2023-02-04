package contest

import (
	"sync"
	"sync/atomic"
)

const mutexLocked = 1

type extMutex struct {
	locked int32
	stMu   sync.Mutex
	c      chan struct{}
}

func New() Mutex {
	return &extMutex{c: make(chan struct{}, 1)}
}

func (nM *extMutex) LockChannel() <-chan struct{} {
	if atomic.CompareAndSwapInt32(&nM.locked, 0, mutexLocked) {
		nM.stMu.Lock()
		nM.c <- struct{}{}
	}
	return nM.c
}

func (nM *extMutex) Lock() {
	nM.stMu.Lock()
	atomic.CompareAndSwapInt32(&nM.locked, 0, mutexLocked)
}

func (nM *extMutex) Unlock() {
	nM.stMu.Unlock()
	atomic.CompareAndSwapInt32(&nM.locked, mutexLocked, 0)
}
