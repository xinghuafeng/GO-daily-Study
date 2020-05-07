package ch18

import (
	"fmt"
	"testing"
	"time"
)

// 多路选择与超时
func service() string {
	time.Sleep(time.Millisecond * 180)
	return "Done"
}
func AsyncService() chan string {
	rech := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("Returne result.")
		rech <- ret
		fmt.Println("Service exited.")
	}()
	return rech

}
func TestAsyncService(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 200):
		t.Error("time out")

	}

}
