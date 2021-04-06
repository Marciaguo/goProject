package ch18

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
func service() string {
	time.Sleep(time.Microsecond * 50)
	return "Done"
}
func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Task is done.")
}
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}