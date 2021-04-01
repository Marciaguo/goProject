package empty_interface

import (
	"fmt"
	"testing"
)

func doSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
	//if i,ok:=p.(int);ok{
	//	fmt.Println("Integer",i)
	//	return
	//}
	//if i,ok:=p.(string);ok{
	//	fmt.Println("String",i)
	//	return
	//}
	//fmt.Println("Unknown Type")
}

func TestDoSomething(t *testing.T) {
	doSomething(10)
	doSomething("10")
}
