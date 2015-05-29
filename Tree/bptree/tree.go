package bptree

import (
	"unsafe"
)

//当int的位宽和指针都是一个字长时，index和leaf的大小都是2*BASE_SIZE个字长
const BASE_SIZE = 16

type node struct {
	inner bool
	cnt   int
	data  [0]int
}

func (unit *leaf) asIndex() *index {
	return (*index)(unsafe.Pointer(unit))
}
func (unit *index) asLeaf() *leaf {
	return (*leaf)(unsafe.Pointer(unit))
}

/*
//关闭数组越界审查后使用
func (unit *node) ceil() int {
	return unit.data[unit.cnt-1]
}
func (unit *node) locate(key int) int {
	var start, end = 0, unit.cnt - 1
	for start < end {
		var mid = (start + end) / 2
		if key > unit.data[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	} //寻找第一个大于或等于key的位置
	return start
}
*/
//开启数组越界审查时使用
func (unit *node) ceil() int {
	if unit.inner {
		var unitx = (*index)(unsafe.Pointer(unit))
		return unitx.data[unitx.cnt-1]
	} else {
		var unitx = (*leaf)(unsafe.Pointer(unit))
		return unitx.data[unitx.cnt-1]
	}
}
func (unit *node) locate(key int) int {
	var start, end = 0, unit.cnt - 1
	if unit.inner {
		var unitx = (*index)(unsafe.Pointer(unit))
		for start < end {
			var mid = (start + end) / 2
			if key > unitx.data[mid] {
				start = mid + 1
			} else {
				end = mid
			}
		}
	} else {
		var unitx = (*leaf)(unsafe.Pointer(unit))
		for start < end {
			var mid = (start + end) / 2
			if key > unitx.data[mid] {
				start = mid + 1
			} else {
				end = mid
			}
		}
	}
	return start
}

type Tree struct {
	root *index
	head *leaf
	path stack
}

func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *Tree) Search(key int) bool {
	if tree.root == nil ||
		key > tree.root.ceil() {
		return false
	}
	var target = tree.root
	for target.inner {
		var idx = target.locate(key)
		if key == target.data[idx] {
			return true
		}
		target = target.kids[idx]
	}
	var unit = target.asLeaf()
	return key == unit.data[unit.locate(key)]
}

func (tree *Tree) Travel(doit func(int)) {
	for unit := tree.head; unit != nil; unit = unit.next {
		for i := 0; i < unit.cnt; i++ {
			doit(unit.data[i])
		}
	}
}
