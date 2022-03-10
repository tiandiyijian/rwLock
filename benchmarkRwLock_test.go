package rwlock

import (
	"sync"
	"testing"
)

var cs1 = 0 // 模拟临界区要保护的数据
var mu1 sync.Mutex

var cs2 = 0 // 模拟临界区要保护的数据
var mu2 sync.RWMutex

var cs3 = 0
var mu3 FairRwLock

var cs4 = 0
var mu4 ReaderFirstRwLock

var cs5 = 0
var mu5 WriterFirstRwLock

func BenchmarkWriteSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu1.Lock()
			cs1++
			mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu1.Lock()
			_ = cs1
			mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.RLock()
			_ = cs2
			mu2.RUnlock()
		}
	})
}

func BenchmarkWriteSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.Lock()
			cs2++
			mu2.Unlock()
		}
	})
}

func BenchmarkWriteSyncByFairRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu3.Lock()
			cs3++
			mu3.Unlock()
		}
	})
}

func BenchmarkReadSyncByFairRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu3.RLock()
			_ = cs3
			mu3.RUnlock()
		}
	})
}

func BenchmarkSyncByRWMutex(b *testing.B) {

	for i := 0; i < 100000; i++ {
		go func() {
			mu2.Lock()
			cs2++
			mu2.Unlock()
		}()
	}

	for i := 0; i < 1000000; i++ {
		go func() {
			mu2.RLock()
			_ = cs2
			mu2.RUnlock()
		}()
	}
	// fmt.Println(cs3)
}

func BenchmarkSyncByFairRWMutex(b *testing.B) {

	for i := 0; i < 100000; i++ {
		go func() {
			mu3.Lock()
			cs3++
			mu3.Unlock()
		}()
	}

	for i := 0; i < 1000000; i++ {
		go func() {
			mu3.RLock()
			_ = cs3
			mu3.RUnlock()
		}()
	}
	// fmt.Println(cs3)
}

func BenchmarkSyncByReaderFirstRWMutex(b *testing.B) {
	for i := 0; i < 100000; i++ {
		go func() {
			mu4.Lock()
			cs4++
			mu4.Unlock()
		}()
	}

	for i := 0; i < 1000000; i++ {
		go func() {
			mu4.RLock()
			_ = cs4
			mu4.RUnlock()
		}()
	}
	// fmt.Println(cs3)
}

func BenchmarkSyncByWriterFirstRWMutex(b *testing.B) {
	for i := 0; i < 100000; i++ {
		go func() {
			mu5.Lock()
			cs5++
			mu5.Unlock()
		}()
	}

	for i := 0; i < 1000000; i++ {
		go func() {
			mu5.RLock()
			_ = cs5
			mu5.RUnlock()
		}()
	}
	// fmt.Println(cs3)
}
