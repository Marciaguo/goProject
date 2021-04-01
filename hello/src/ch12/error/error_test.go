package error

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New(" n should be not less than 2")
var LargerThanHundredError = errors.New(" n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, errors.New(LessThanTwoError.Error())
	}
	if n > 100 {
		return nil, errors.New(LargerThanHundredError.Error())
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-1); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
