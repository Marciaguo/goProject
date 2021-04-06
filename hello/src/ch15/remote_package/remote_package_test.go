package remote_package

import (
	"fmt"
	cm "github.com/easierway/concurrent_map"
	"runtime"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("test")))
}
func TestGetVersion(t *testing.T) {
	fmt.Println(runtime.Version())
}
