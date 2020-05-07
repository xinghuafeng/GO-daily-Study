package ch17

import (
	"sync"
	"testing"
	"time"
)

func TestCountThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++

		}()

	}
	time.Sleep(1 * time.Second)
	t.Logf("countValue= %d", counter)
}
