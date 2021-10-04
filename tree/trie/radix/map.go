package radix

import (
	"unsafe"
)

const bitWidth = uint(unsafe.Sizeof(uint(0))) * 8
const Step = uint(2) //2 or 4
const Radix = uint(1) << Step
const Depth = bitWidth / Step
const Mask = Radix - 1

type node struct {
	kids [Radix]unsafe.Pointer
}

type Map struct {
	root node
}

func branch(key uint, depth uint) uint {
	return (key >> (bitWidth - (depth+1)*Step)) & Mask
}

func (m *Map) Search(key uint) unsafe.Pointer {
	root := &m.root
	for i := uint(0); i < Depth && root != nil; i++ {
		root = (*node)(root.kids[branch(key, i)])
	}
	return unsafe.Pointer(root)
}

func (m *Map) Insert(key uint, val unsafe.Pointer) bool {
	root := &m.root
	for i := uint(0); i < Depth-1; i++ {
		idx := branch(key, i)
		if root.kids[idx] == nil {
			root.kids[idx] = unsafe.Pointer(new(node))
		}
		root = (*node)(root.kids[idx])
	}
	idx := key & Mask
	if root.kids[idx] != nil {
		return false
	}
	root.kids[idx] = val
	return true
}

func (m *Map) Remove(key uint) bool {
	var path [Depth]*node
	path[0] = &m.root
	for i := uint(0); i < Depth-1; i++ {
		path[i+1] = (*node)(path[i].kids[branch(key, i)])
		if path[i+1] == nil {
			return false
		}
	}
	idx := key & Mask
	if path[Depth-1].kids[idx] == nil {
		return false
	}
	path[Depth-1].kids[idx] = nil
	for i := Depth - 1; i != 0; i-- {
		j := uint(0)
		for j < Radix && path[i].kids[j] == nil {
			j++
		}
		if j == Radix { //全空
			path[i-1].kids[branch(key, i-1)] = nil
		}
	}
	return true
}
