package contest

import (
	"sync/atomic"
)

const mutexLocked = 1

type newMutex struct {
	locked int32
	c      chan struct{}
}

func New() Mutex {
	return &newMutex{c: make(chan struct{}, 1)}
}

func (nM *newMutex) LockChannel() <-chan struct{} {
	if atomic.CompareAndSwapInt32(&nM.locked, 0, mutexLocked) {
		nM.c <- struct{}{}
	}
	return nM.c
}

func (nM *newMutex) Lock() {
	if atomic.CompareAndSwapInt32(&nM.locked, 0, mutexLocked) {
		return
	}

	for {
		if atomic.CompareAndSwapInt32(&nM.locked, 0, mutexLocked) {
			break
		}
	}
}

func (nM *newMutex) Unlock() {
	atomic.CompareAndSwapInt32(&nM.locked, mutexLocked, 0)
}
