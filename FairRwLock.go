package rwlock

import "sync"

type FairRwLock struct {
	rCount     int
	rCountLock sync.Mutex
	wDataLock  sync.Mutex
	flag       sync.Mutex
}

func (l *FairRwLock) RLock() {
	l.flag.Lock()
	l.rCountLock.Lock()
	if l.rCount == 0 {
		l.wDataLock.Lock()
	}
	l.rCount++
	l.rCountLock.Unlock()
	l.flag.Unlock()
}

func (l *FairRwLock) RUnlock() {
	l.rCountLock.Lock()
	l.rCount--
	if l.rCount == 0 {
		l.wDataLock.Unlock()
	}
	l.rCountLock.Unlock()
}

func (l *FairRwLock) Lock() {
	l.flag.Lock()
	l.wDataLock.Lock()
}

func (l *FairRwLock) Unlock() {
	l.wDataLock.Unlock()
	l.flag.Unlock()
}
