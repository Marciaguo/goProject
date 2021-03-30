package type_test

import "testing"
type myInt int64
func TestImplication(t *testing.T){
	var a int32 = 1
	var b int64
	b=int64(a)
	var c myInt
	c = myInt(b)

	t.Log(a,b,c)
}

func TestPoint(t *testing.T){
	a:=1
	aPtr:=&a
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)
}

func TestString(t *testing.T){
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))
}