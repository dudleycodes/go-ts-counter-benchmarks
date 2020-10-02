package benchmarks

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	_sut     int64
	_mutex   sync.Mutex
	_rwMutex sync.RWMutex
)

func atomicRead() {
	atomic.LoadInt64(&_sut)
}

func atomicUpdate() {
	atomic.StoreInt64(&_sut, time.Now().Unix())
}

func mutexRead() {
	_mutex.Lock()
	defer _mutex.Unlock()

	_ = _sut
}

func mutexUpdate() {
	_mutex.Lock()
	defer _mutex.Unlock()

	_sut = time.Now().Unix()
}

func rwMutextUpdate() {
	_rwMutex.Lock()
	_rwMutex.Unlock()
	return
}

func rwMutextRead() {
	_rwMutex.RLock()
	defer _rwMutex.RUnlock()
	_ = _sut
}

func runBenchmark(b *testing.B, u func(), r func()) {
	_sut = 0
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					u()
				}
			}
		}()
	}

	for n := 0; n < b.N; n++ {
		if n%100 == 0 {
			r()
		}

		time.Sleep(100)
	}

	cancel()
	wg.Wait()
}

func BenchmarkAtomic(b *testing.B) {
	runBenchmark(b, atomicUpdate, atomicRead)
}

func BenchmarkMutext(b *testing.B) {
	runBenchmark(b, mutexUpdate, mutexRead)
}

func BenchmarkRWMutex(b *testing.B) {
	runBenchmark(b, rwMutextUpdate, rwMutextRead)
}
