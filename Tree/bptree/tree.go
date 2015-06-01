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

func (u *leaf) asIndex() *index {
	return (*index)(unsafe.Pointer(u))
}
func (u *index) asLeaf() *leaf {
	return (*leaf)(unsafe.Pointer(u))
}

/*
//关闭数组越界审查后使用
func (u *node) ceil() int {
	return u.data[u.cnt-1]
}
func (u *node) locate(key int) int {
	var start, end = 0, u.cnt - 1
	for start < end {
		var mid = (start + end) / 2
		if key > u.data[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	} //寻找第一个大于或等于key的位置
	return start
}
*/
//开启数组越界审查时使用
func (u *node) ceil() int {
	if u.inner {
		var unit = (*index)(unsafe.Pointer(u))
		return unit.data[unit.cnt-1]
	} else {
		var unit = (*leaf)(unsafe.Pointer(u))
		return unit.data[unit.cnt-1]
	}
}
func (u *node) locate(key int) int {
	var start, end = 0, u.cnt - 1
	if u.inner {
		var unitx = (*index)(unsafe.Pointer(u))
		for start < end {
			var mid = (start + end) / 2
			if key > unitx.data[mid] {
				start = mid + 1
			} else {
				end = mid
			}
		}
	} else {
		var unitx = (*leaf)(unsafe.Pointer(u))
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

func (tr *Tree) IsEmpty() bool {
	return tr.root == nil
}

func (tr *Tree) Search(key int) bool {
	if tr.root == nil ||
		key > tr.root.ceil() {
		return false
	}
	var target = tr.root
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

func (tr *Tree) Travel(doit func(int)) {
	for unit := tr.head; unit != nil; unit = unit.next {
		for i := 0; i < unit.cnt; i++ {
			doit(unit.data[i])
		}
	}
}
