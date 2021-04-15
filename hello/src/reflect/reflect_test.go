package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}

type Employee struct {
	ID   string
	Name string
	Age  int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Consumer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "nike", 30}
	t.Logf("Name:value(%[1]v),Type(%[1]T)",
		reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("failed to get 'Name' field")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//t.Log(a==b)
	fmt.Println(reflect.DeepEqual(a, b))
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1==s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1==s3?", reflect.DeepEqual(s1, s3))

	c1 := Consumer{"1", "Mike", 40}
	c2 := Consumer{"1", "Mike", 40}
	fmt.Println(reflect.DeepEqual(c1, c2))
}
