package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues()(int,int){
	return rand.Intn(10),rand.Intn(20)
}
func TestFunc(t *testing.T){
	//a,_:=returnMultiValues()
	//t.Log(a)
	tsSF:=timeSpent(slowFun)
	t.Log(tsSF(10))
}

func timeSpent(inner func(op int) int) func(op int) int{
	return func(n int) int {
		start:=time.Now()
		ret:=inner(n)
		fmt.Println("time spend:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int{
	time.Sleep(time.Second*1)
	return op
}

func Sum(ops ...int) int{
	ret:=0
	for _,op:=range ops{
		ret+=op
	}
	return ret
}
func TestVarParam(t *testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5,6))
}

func TestDefer(t *testing.T){
	defer func(){
		t.Log("clear resources")
	}()
	t.Log("started")
	panic("fatal error")
}