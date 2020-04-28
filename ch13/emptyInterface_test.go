package ch13

import (
	"fmt"
	"testing"
)
//空接口 可以代替所有类型；
//通过断言 来将空接口转化为指定类型
func  doSomething( p interface{})  {
	if i,ok:=p.(int);ok{
		fmt.Println("Int",i)
		return
	}
	if s,ok:=p.(string);ok{
		fmt.Println("string",s)
		return
	}
	fmt.Println("Unknow!")

	
}
func  TestEmptyeAssertion( t *testing.T)  {
	doSomething(10)
	doSomething("abc")
}
