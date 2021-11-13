// пакет Atomic (атомарные операции)

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var status int32 = 0

func main() {
	var wg sync.WaitGroup // без WG будет давать нестабильное колличество итераций 999 или 998
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&status, 1)
		}()
	}

	wg.Wait()
	fmt.Println(atomic.LoadInt32(&status)) // обращаться напрямую к status вызовет гонку
}
