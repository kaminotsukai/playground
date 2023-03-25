package goroutine

import (
	"sync"
	"sync/atomic"
)

// 1から10000までの和を求めるプログラムを並行で実行する
func SyncAtmicFunc() int32 {
	var result int32
	var wg sync.WaitGroup // sync.WaitGroup を利用してゴルーチンの実行が完了するのを待機する
	for i := 1; i <= 10000; i++ {
		wg.Add(1)

		// 各ゴルーチンが参照する値が互いに干渉しないように実行時に変数iを渡す
		// 変数iをループ内で閉じた変数にすることで同じ挙動を実現できる
		go func(i int) {
			defer wg.Done()
			atomic.AddInt32(&result, int32(i))
		}(i)
	}
	wg.Wait()
	return result // expect: 50005000
}

// 1から10000までの和を求めるプログラムを並行で実行する
func SyncMutexFunc() int32 {
	var result int32
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			result += int32(i)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	return result // expect: 50005000
}
