package diysort

import (
	"math/rand"
	"testing"
	"time"
)

func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func ramdomArray(size int) []int {
	rand.Seed(time.Now().Unix())
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}
func checkArrary(list []int) bool {
	for i, size := 1, len(list); i < size; i++ {
		if list[i] < list[i-1] {
			return false
		}
	}
	return true
}

func Test_InsertSort(t *testing.T) {
	defer guard_ut(t)
	var list = ramdomArray(100)
	InsertSort(list)
	if !checkArrary(list) {
		t.Fail()
	}
}
func Test_SelectSort(t *testing.T) {
	defer guard_ut(t)
	var list = ramdomArray(100)
	SelectSort(list)
	if !checkArrary(list) {
		t.Fail()
	}
}
func Test_HeapSort(t *testing.T) {
	defer guard_ut(t)
	var list = ramdomArray(1000)
	HeapSort(list)
	if !checkArrary(list) {
		t.Fail()
	}
}
func Test_MergeSort(t *testing.T) {
	defer guard_ut(t)
	var list = ramdomArray(1000)
	MergeSort(list)
	if !checkArrary(list) {
		t.Fail()
	}
}
func Test_QuickSort(t *testing.T) {
	defer guard_ut(t)
	var list = ramdomArray(1000)
	QuickSort(list)
	if !checkArrary(list) {
		t.Fail()
	}
}
func Test_Introsort(t *testing.T) {
	defer guard_ut(t)
	var list = ramdomArray(1000)
	Introsort(list)
	if !checkArrary(list) {
		t.Fail()
	}
}

func ramdomLinkList(size int) *Node {
	rand.Seed(time.Now().Unix())
	var head *Node
	var tail = FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.next = new(Node)
		tail = tail.next
		tail.key = rand.Int()
	}
	tail.next = nil
	return head
}
func checkLinkList(head *Node, size int) bool {
	var cnt = 1
	for ; head.next != nil; head = head.next {
		if head.next.key < head.key {
			return false
		}
		cnt++
	}
	return cnt == size
}

func Test_LinkSort(t *testing.T) {
	defer guard_ut(t)

	const size = 1000
	rand.Seed(time.Now().Unix())

	var head = ramdomLinkList(size)
	head = LinkSort(head)
	if !checkLinkList(head, size) {
		t.Fail()
	}
}
