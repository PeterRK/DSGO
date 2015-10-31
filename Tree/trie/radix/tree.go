package radix

import (
	"unsafe"
)

const UINT_LEN = uint(unsafe.Sizeof(uint(0))) * 8
const STEP = uint(2) //2 or 4
const KIDS_LIMIT = uint(1) << STEP
const DEPTH = UINT_LEN / STEP
const MASK = KIDS_LIMIT - 1

type node struct {
	kids [KIDS_LIMIT]*node
}
type Tree struct {
	root node
}

func cut(key uint, i uint) uint {
	return (key >> ((UINT_LEN - STEP) - i*STEP)) & MASK
}

func (tr *Tree) Search(key uint) unsafe.Pointer {
	var root = &tr.root
	for i := uint(0); i < DEPTH && root != nil; i++ {
		root = root.kids[cut(key, i)]
	}
	return unsafe.Pointer(root)
}

func (tr *Tree) Insert(key uint, ptr unsafe.Pointer) bool {
	var root = &tr.root
	for i := uint(0); i < DEPTH-1; i++ {
		var idx = cut(key, i)
		if root.kids[idx] == nil {
			root.kids[idx] = new(node)
		}
		root = root.kids[idx]
	}
	var idx = key & MASK
	if root.kids[idx] != nil {
		return false
	}
	root.kids[idx] = (*node)(ptr)
	return true
}

func (tr *Tree) Remove(key uint) bool {
	var path [DEPTH]*node
	path[0] = &tr.root
	for i := uint(0); i < DEPTH-1; i++ {
		path[i+1] = path[i].kids[cut(key, i)]
		if path[i+1] == nil {
			return false
		}
	}
	var idx = key & MASK
	if path[DEPTH-1].kids[idx] == nil {
		return false
	}
	path[DEPTH-1].kids[idx] = nil
	for i := DEPTH - 1; i != 0; i-- {
		var j = uint(0)
		for j < KIDS_LIMIT && path[i].kids[j] == nil {
			j++
		}
		if j == KIDS_LIMIT { //全空
			path[i-1].kids[cut(key, i-1)] = nil
		}
	}
	return true
}
