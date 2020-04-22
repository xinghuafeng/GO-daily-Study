package ch5

import "testing"

func  TestArrayDemo(t *testing.T)  {

	var arr [3] int //默认值为0
	t.Log(arr[1],arr[2])
	 arr1:=[3] int {1,2,3}
	 arr2:=[] int {1,2,3,4,5,6}
	t.Log(arr1)
	t.Log(arr2)
}
//对数组遍历
func TestPrintArray(t *testing.T)  {
	//传统语言写法遍历输出数组
	arr1:=[] int {1,2,3,4,5,6}
	for i:=0;i<len(arr1) ;i++  {
		t.Log(arr1[i])
	}
	//GO语言有类似java 那种foreach，ep: id,value
	for idx,v:= range arr1  {
		t.Log(idx,v)

	}
	//GO 語言省略id寫法
	for _,v:= range arr1  {
		t.Log(v)

	}

}
//数组截取
func  TestSliceArray(t *testing.T)  {
	 arr:=[4] int{1,2,3,4}
	 t.Log(arr[0])//截取第一个元素
	t.Log(arr[0:3])//截取第1个到第4个元素之间，不包括第4个原始
	t.Log(arr[1:len(arr)])//截取从第2个元素到最后一个元素

}


