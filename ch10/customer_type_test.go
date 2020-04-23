package ch10

import (
	"fmt"
	"time"
)
//用于计算时间的方法 ，自定义函数？如何优化，闭包？
type ConvertMethod func(op int) int
func  TimeSpent(inner func(op int) int)  func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent:", time.Since(start).Seconds())
		return ret

	}
}
func TimeSpent2(inner ConvertMethod) ConvertMethod {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent:", time.Since(start).Seconds())
		return ret

	}
}



