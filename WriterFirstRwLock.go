package rwlock

import "sync"

type WriterFirstRwLock struct {
	rCount, wCount int
	rCountLock     sync.Mutex
	wCountLock     sync.Mutex
	wDataLock      sync.Mutex
	rLock          sync.Mutex
}

func (l *WriterFirstRwLock) RLock() {
	l.rLock.Lock()
	l.rCountLock.Lock()
	if l.rCount == 0 {
		l.wDataLock.Lock()
	}
	l.rCount++
	l.rCountLock.Unlock()
	l.rLock.Unlock()
}

func (l *WriterFirstRwLock) RUnlock() {
	l.rCountLock.Lock()
	l.rCount--
	if l.rCount == 0 {
		l.wDataLock.Unlock()
	}
	l.rCountLock.Unlock()
}

func (l *WriterFirstRwLock) Lock() {
	l.wCountLock.Lock()
	if l.wCount == 0 {
		l.rLock.Lock()
	}
	l.wCount++
	l.wCountLock.Unlock()
	l.wDataLock.Lock()
}

func (l *WriterFirstRwLock) Unlock() {
	l.wDataLock.Unlock()
	l.wCountLock.Lock()
	l.wCount--
	if l.wCount == 0 {
		l.rLock.Unlock()
	}
	l.wCountLock.Unlock()
}
