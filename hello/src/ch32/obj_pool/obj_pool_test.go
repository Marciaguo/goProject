package obj_pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReusableObj struct {
}
type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用对象
}

func newObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	default:
		return nil, errors.New("overflow")
	}

}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := newObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Println("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
