package ch10

import (
	"testing"
)
//go接口与其他编程语言差异
//1/非侵入式，实现不依赖接口定义
//2.可以定义在使用者包里
type Programmer interface {
 WriteHellowolrd() string
}
type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHellowolrd() string {
	return "fmt.Print(\"Hello,wolrd\")"
}
func TestGoProgrammer(t *testing.T)  {
	var p Programmer
	p=new(GoProgrammer)
	t.Log(p.WriteHellowolrd())
}
