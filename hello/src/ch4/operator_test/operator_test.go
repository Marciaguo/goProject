package operator_test

import "testing"
const(
	Readable = 1<<iota
	Writable
	Executable
)
func TestBitClear(t *testing.T){
	a :=7
	a = a&^Readable
	t.Log(a)
	t.Log(a&Readable==Readable,a&Writable==Writable,a&Executable==Executable)
}
func TestOperator(t *testing.T) {
	a:=[...]int{1,2,3,4}
	b:=[...]int{1,3,4,5}
	c:=[...]int{1,3,4,5,6}
	d:=[...]int{1,3,4,5,6}

	t.Log(a==b)
	t.Log(c==d)
}
