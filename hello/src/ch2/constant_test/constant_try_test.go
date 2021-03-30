package constant_test

import "testing"

const(
	Monday=iota+1
	TuesDay
	Wednesday
)
const(
	Readable = 1<<iota
	Writable
	Executable
)
func TestConstant(t *testing.T){
	a:=7//0111
	t.Log(a&Readable==Readable,a&Writable==Writable,a&Executable==Executable)
	//t.Log(Monday,TuesDay)
}