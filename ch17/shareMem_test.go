package ch17

import (
	"sync"
	"testing"
)

//锁共享内存并發機制
func TestCountThreadSafeByWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wt sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wt.Add(1)
		go func() {
			defer func() {
				mut.Unlock()

			}()
			mut.Lock()
			counter++
			wt.Done()

		}()

	}
	wt.Wait()
	t.Logf("countValue= %d", counter)
}
