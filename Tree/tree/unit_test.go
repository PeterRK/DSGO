package tree

import (
	"testing"
)

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guardUT(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func Test_Iterator(t *testing.T) {
	guardUT(t)

	const size = 100
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}

	var root = BuildBalanceTree(list, nil)

	var node = root
	for node.left != nil {
		node = node.left
	}
	for i := 0; i < size; i++ {
		assert(t, node != nil && node.key == i)
		node = MoveForward(node)
	}
	assert(t, MoveForward(node) == nil)

	node = root
	for node.right != nil {
		node = node.right
	}
	for i := size - 1; i >= 0; i-- {
		assert(t, node != nil && node.key == i)
		node = MoveBackward(node)
	}
	assert(t, MoveBackward(node) == nil)
}
