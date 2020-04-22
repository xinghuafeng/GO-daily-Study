package ch6

import (
	"testing"
)
/*
数组与切片区别
1.容量是否可以伸缩
2.是否可以比较
 */
//切片结构定义
func TestSliceInit(t *testing.T)  {
	var s0 [] int
	t.Log(len(s0),cap(s0))
	s0=append(s0,1)
	t.Log(len(s0),cap(s0))
	s1:=[] int{1,2,3,4,5}
	t.Log(len(s1),cap(s1))
	s2:=make([] int,3,5)
	t.Log(len(s2),cap(s2))
	s2=append(s2,1)
	t.Log(s2[0],s2[1],s2[2],s2[3])



}
//切片如何实现变长.放入元素的超出原有的cap的长度时候，扩容*2
func  TestSliceGrowing(t *testing.T)  {
	var s0 [] int
 for i:=1;i<10;i++{
	 s0=append(s0,i)
	 t.Log(len(s0),cap(s0))
 }
}
//切片共享存储空间,如果一个被改其他也会被修改
func TestSliceShareMemory(t *testing.T){

	a:=[]string{"one","two","three","foure","five","six"}
	t.Log(len(a),cap(a))
	q1 :=a[1:4]
	t.Log(q1,len(q1),cap(q1))
	q2 :=a[3:5]
	t.Log(q2,len(q2),cap(q2))
	q2[0]="unKnow"
	t.Log(q1,len(q1),cap(q1))
	t.Log(q2,len(q2),cap(q2))


}
