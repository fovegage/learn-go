package lock

import (
	"sync"
	"time"
)

type KeyLock struct {
	l sync.Mutex
	m map[interface{}]struct{}
}

func NewKeyLock() *KeyLock {
	return &KeyLock{
		l: sync.Mutex{},
		m: make(map[interface{}]struct{}),
	}
}

func (kl *KeyLock) Lock(key interface{}) {
	if kl == nil {
		return
	}
	for {
		kl.l.Lock()
		if _, exist := kl.m[key]; !exist {
			break
		}
		kl.l.Unlock()
		time.Sleep(50 * time.Millisecond)
	}
	kl.m[key] = struct{}{}
	kl.l.Unlock()
	return
}

func (kl *KeyLock) ULock(key interface{}) {
	if kl == nil {
		return
	}
	kl.l.Lock()
	delete(kl.m, key)
	kl.l.Unlock()
}
