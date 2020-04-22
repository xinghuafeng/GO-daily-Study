package ch1

import "testing"
func TestFib(t *testing.T)  {
	//要求输出10个以内 fib数字：比如： 1 1 2 3 5 8
a:=1
b:=1
t.Log(a)
for i:=0;i<10;i++{
	t.Log(b)
	temp :=a
	a = b
	b =temp +a
}

}
//go與其他語言變量賦值差別 1：GO可以自動判斷賦值；2.GO可以同時對多個變量賦值
func TestExangeValue(t *testing.T){
	a:=1
	b:=2
	a,b = b,a
	t.Log(a,b)
}
//中間變量賦值法
func  TestExange( t *testing.T)  {
	a:=1
	b:=2
	tem :=a
	a = b
	b = tem
	t.Log(a,b)
}

