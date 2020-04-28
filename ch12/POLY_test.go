package ch12

import (
	"fmt"
	"testing"
)

//实现多态
type Code string
type ProgramDemo interface {
	WriteHellWolrd() Code
}
type JavaTypeWrite struct {

}
type GoTypeWrite struct {

}

func (p *JavaTypeWrite)WriteHellWolrd() Code  {

	return  "来源java hello wolrd"
	//fmt.Printf("%T %v",p,p.WriteJavaHellWolrd("来源java hello wolrd"))
}
func (p *GoTypeWrite)WriteHellWolrd() Code  {

	return  "来源Go hello wolrd"
	//fmt.Printf("%T %v",p,p.WriteJavaHellWolrd("来源java hello wolrd"))
}
func  writeHelloWorld( p ProgramDemo)   {
	fmt.Printf("%T %v \n",p,p.WriteHellWolrd())
}
func  TestDemo( t *testing.T)  {
	javaHello := new(JavaTypeWrite)
	GOHello := new(GoTypeWrite)
	writeHelloWorld(javaHello)
	writeHelloWorld(GOHello)
}

	

