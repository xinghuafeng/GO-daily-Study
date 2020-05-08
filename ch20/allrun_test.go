package ch20

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
仅需任意任务完成
chanel 与buffer chanel区别。buffer有接受者就会拿走不需要等待.chane 阻塞的
buffer channel为防止协程的泄露

*/
func runTash(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)

}
func FirstResponse() string {

	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTash(i)
			ch <- ret
		}(i)
	}
	return <-ch
}
func TestFirstRespone(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
