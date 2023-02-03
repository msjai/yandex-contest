package contest

import (
	"sync/atomic"
)

const mutexLocked = 1

type extMutex struct {
	locked int32
	c      chan struct{}
}

func New() Mutex {
	return &extMutex{c: make(chan struct{}, 1)}
}

func (nM *extMutex) LockChannel() <-chan struct{} {
	if atomic.CompareAndSwapInt32(&nM.locked, 0, mutexLocked) {
		nM.c <- struct{}{}
	}
	return nM.c
}

func (nM *extMutex) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&nM.locked, 0, mutexLocked) {
			return
		}
	}
}

func (nM *extMutex) Unlock() {
	atomic.CompareAndSwapInt32(&nM.locked, mutexLocked, 0)
}
