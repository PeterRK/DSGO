package radix

import (
	"unsafe"
)

const lenOfUint = uint(unsafe.Sizeof(uint(0))) * 8
const step = uint(2) //2 or 4
const kidsLimit = uint(1) << step
const depth = lenOfUint / step
const mask = kidsLimit - 1

type node struct {
	kids [kidsLimit]*node
}
type Tree struct {
	root node
}

func cut(key uint, i uint) uint {
	return (key >> ((lenOfUint - step) - i*step)) & mask
}

func (tr *Tree) Search(key uint) unsafe.Pointer {
	var root = &tr.root
	for i := uint(0); i < depth && root != nil; i++ {
		root = root.kids[cut(key, i)]
	}
	return unsafe.Pointer(root)
}

func (tr *Tree) Insert(key uint, ptr unsafe.Pointer) bool {
	var root = &tr.root
	for i := uint(0); i < depth-1; i++ {
		var idx = cut(key, i)
		if root.kids[idx] == nil {
			root.kids[idx] = new(node)
		}
		root = root.kids[idx]
	}
	var idx = key & mask
	if root.kids[idx] != nil {
		return false
	}
	root.kids[idx] = (*node)(ptr)
	return true
}

func (tr *Tree) Remove(key uint) bool {
	var path [depth]*node
	path[0] = &tr.root
	for i := uint(0); i < depth-1; i++ {
		path[i+1] = path[i].kids[cut(key, i)]
		if path[i+1] == nil {
			return false
		}
	}
	var idx = key & mask
	if path[depth-1].kids[idx] == nil {
		return false
	}
	path[depth-1].kids[idx] = nil
	for i := depth - 1; i != 0; i-- {
		var j = uint(0)
		for j < kidsLimit && path[i].kids[j] == nil {
			j++
		}
		if j == kidsLimit { //全空
			path[i-1].kids[cut(key, i-1)] = nil
		}
	}
	return true
}
