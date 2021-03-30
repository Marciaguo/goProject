package ch5

import "testing"

func TestWhileLoop(t *testing.T){
	/*while(n<5)*/
	n:=0
	for n<5{
		t.Log(n)
		n++
	}
}