package ch3

import "testing"

/*

 不支持类型转换
不支持指针运算
string 是值类型，默认值为空，非nil

*/
func  TestPoint(t *testing.T)  {
	 a := 1
	aptr :=&a

	t.Log(a,aptr)

}
func  TestString(t *testing.T)  {
	var s string
	t.Log(len(s))

}

