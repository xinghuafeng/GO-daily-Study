package ch14

import (
	"errors"
	"testing"
)

func Fib(n int  )([]int,error) {
	if n<2||n>100 {
		return nil,errors.New("n should be in[2,100]")
	}
	fiblist:=[] int {1,1}
	for i:=2;i<n;i++{
		fiblist=append(fiblist,fiblist[i-2]+fiblist[i-1])
	}
	return  fiblist,nil
	
}
func TestDemo( t *testing.T)  {

	if v,error:= Fib(50);error!=nil{
		t.Error(error)
	} else {
		t.Log(v)
	}
}
