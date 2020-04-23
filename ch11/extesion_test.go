package ch11

import (
	"fmt"
	"testing"
)

//go里没有继承，但是可以扩展与复用，提供了一种匿名嵌套类型
//扩展与复用
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Printf("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host)
}

//匿名嵌套类型
type Dog struct {
	//p *Pet
	Pet
}

//func (d *Dog) speak(){
//	fmt.Println("wang")
//	//d.p.Speak()
//}
//func (d *Dog)SpeakTo(host string) {
//	d.p.SpeakTo(host)
//}
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("chao")
}
