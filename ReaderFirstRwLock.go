package rwlock

import "sync"

type ReaderFirstRwLock struct {
	rCount     int
	rCountLock sync.Mutex
	wDataLock  sync.Mutex
}

func (l *ReaderFirstRwLock) RLock() {
	l.rCountLock.Lock()
	if l.rCount == 0 {
		l.wDataLock.Lock()
	}
	l.rCount++
	l.rCountLock.Unlock()
}

func (l *ReaderFirstRwLock) RUnlock() {
	l.rCountLock.Lock()
	l.rCount--
	if l.rCount == 0 {
		l.wDataLock.Unlock()
	}
	l.rCountLock.Unlock()
}

func (l *ReaderFirstRwLock) Lock() {
	l.wDataLock.Lock()
}

func (l *ReaderFirstRwLock) Unlock() {
	l.wDataLock.Unlock()
}
