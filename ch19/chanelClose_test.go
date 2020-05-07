package ch19

import (
	"fmt"
	"sync"
	"testing"
)

//channel 的关闭
/*
1.向关闭的channel 发送数据，会导致painc
2.v,ok<-ch;ok为Bool ,ture正常接收，false 表示通道关闭
3.所有channe接受者会在channle关闭时，立刻从阻塞等待中返回ok false
这种广播机制被利用，可以向多个订阅者发送信号，如退出信号


*/
//消费者
func Recive(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}

		}
		wg.Done()
	}()

}

//生产者
func Product(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}
func TestChanel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	Product(ch, &wg)
	wg.Add(1)
	Recive(ch, &wg)
	wg.Add(1)
	Recive(ch, &wg)
	wg.Wait()

}
