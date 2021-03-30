package map_test

import "testing"

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Log("len m1:", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len m2:", len(m2))
	m3 := make(map[int]int, 3)
	t.Log("len m3:", len(m3))
}

//判断key是否存在
func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	if v, ok := m1[1]; ok {
		t.Log("key 1's value is ", v)
	} else {
		t.Log("key 1 is not exist")
	}
}

//遍历map
func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k,v)
	}
}
