package ch15

import (
	"fmt"
	"testing"
)

func TestDemo( t *testing.T) {
defer func() {//相当于java中的finnlly
	fmt.Println("demo")
}()
	f()
	defer func() {
		fmt.Println("FUCK——44")
	}()

}
func  TestDEMO1(t *testing.T)  {
	defer func() {//相当于java中的finnlly
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f()
	defer func() {
		fmt.Println("FUCK——3")
	}()
}
func f() {
	defer func() {
		fmt.Println("inner func defer")
		if err := recover(); err != nil {
			fmt.Println(err, " fuck")
		}
	}()
	fmt.Println("A")
	panic(3)
	defer func() {
		fmt.Println("inner func defer 1")
	}()
	//以下2个不能执行
	fmt.Println("b")
	fmt.Println("c")
}
