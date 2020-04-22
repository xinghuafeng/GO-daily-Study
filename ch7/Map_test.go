package ch7

import "testing"

func TestMap(t *testing.T)  {
m1:=map[int]	int{1:1,2:4,3:9}
t.Log(len(m1))//输出长度
t.Log(m1[1])//输出元素
}
//判断值是否不存在
func  TestExitMapValue(t *testing.T)  {
	m1:=map[int] int{}
	m1[0]=1
	m1[2]=1
	if  _,ok:=m1[3];ok{
	t.Log("存在k3")
	}else {
		t.Log("不存在k3")
	}

}
//遍历map
func  TestTrallMap(t *testing.T)  {
	
	m1:=map[int]	int{1:1,2:4,3:9}
	for k,v:=range m1{
t.Log(k,v)
	}

}