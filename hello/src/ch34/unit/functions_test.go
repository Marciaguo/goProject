package unit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expects := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expects[i] {
			t.Errorf("input is %d,the actual is %d,the expected is %d", inputs[i], ret, expects[i])
		}
	}
}

func TestError(t *testing.T) {
	fmt.Println("start")
	t.Error("error")
	fmt.Println("end")
}

func TestFatal(t *testing.T) {
	fmt.Println("start")
	t.Fatal("error")
	fmt.Println("end")
}

func TestWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expects := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expects[i], ret)
	}
}
