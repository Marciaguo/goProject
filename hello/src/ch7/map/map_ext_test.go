package _map

import "testing"

func TestMapWithFunValue(t *testing.T){
	m:=map[int]func(op int) int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](1),m[2](2),m[3](3))
}
func TestMapForSet(t *testing.T){
	mySet:=map[int]bool{}
	mySet[1]=true
	n:=1
	//元素唯一
	if mySet[n]{
		t.Logf("%d is existing",n)
	}else{
		t.Logf("%d is not existing",n)
	}
	mySet[3]=true
	//集合的长度
	t.Log("len is:",len(mySet))
	//删除指定元素
	delete(mySet,3)
}