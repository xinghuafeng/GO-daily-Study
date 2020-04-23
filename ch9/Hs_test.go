package ch9

import (
	"fmt"
	"testing"
)

func TestHs(t *testing.T) {
	t.Log(Add(2, 3))
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */

	Swap1(&a, &b)
	t.Log(a, b)
}
func Add(x, y int) int {
	/* 定义局部变量 */
	var result int
	result = x + y
	return result
}

//go 默认按照值传递 值传递不会改变变量值
//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
/* 定义相互交换值的函数 */
func Swap1(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

//GO 是函数编程，函数可以作为另外一个函数的实参传递
type cb func(int) int

func TestByHS(t *testing.T) {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，xx：%d\n", x)
		return x
	})
}
func testCallBack(x int, f cb) {
	f(x)
}
func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

//Go 语言函数闭包 Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
func TestBB(t *testing.T) {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

}
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，
//接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
//所有给定类型的方法属于该类型的方法集。
type Circle struct {
	radius float64
}

func TestMethod(t *testing.T) {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}
func (c Circle) getArea() float64 {

	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clean")
	}()

	t.Log("start")
	panic("error")
}
