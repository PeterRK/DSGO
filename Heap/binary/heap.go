package binary

import (
	"errors"
)

//二叉堆，底层采用数组。
//Build的复杂度为O(N)，Top的复杂度为O(1)，其余核心操作复杂度为O(logN)。
type Heap struct {
	core []int
}

func (hp *Heap) Size() int {
	return len(hp.core)
}
func (hp *Heap) IsEmpty() bool {
	return len(hp.core) == 0
}
func (hp *Heap) Clear() {
	hp.core = hp.core[:0]
}

func (hp *Heap) Top() (int, error) {
	if hp.IsEmpty() {
		return 0, errors.New("empty")
	}
	return hp.core[0], nil
}

func (hp *Heap) Build(list []int) {
	var size = len(list)
	hp.core = list
	for idx := size/2 - 1; idx >= 0; idx-- {
		hp.adjustDown(idx)
	}
}
func (hp *Heap) Push(key int) {
	var place = len(hp.core)
	hp.core = append(hp.core, key)
	hp.adjustUp(place)
}
func (hp *Heap) Pop() (int, error) {
	var size = hp.Size()
	if size == 0 {
		return 0, errors.New("empty")
	}

	var key = hp.core[0]
	if size == 1 {
		hp.core = hp.core[:0]
	} else {
		hp.core[0] = hp.core[size-1]
		hp.core = hp.core[:size-1]
		hp.adjustDown(0)
	}
	return key, nil
}
