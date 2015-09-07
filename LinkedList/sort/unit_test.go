package sort

import (
	"LinkedList/list"
	"math/rand"
	"testing"
	"time"
)

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

const sz_big = 500
const sz_small = 100

func Test_MergeSort(t *testing.T) {
	testLinkListSort(t, MergeSort, sz_big, sz_small)
}
func Test_QuickSort(t *testing.T) {
	testLinkListSort(t, QuickSort, sz_big, sz_small)
}
func Test_IntroSort(t *testing.T) {
	testLinkListSort(t, IntroSort, sz_big, sz_big)
}
func Test_RadixSort(t *testing.T) {
	testLinkListSort(t, RadixSort, sz_big, sz_small)
}

func testLinkListSort(t *testing.T, doit func(*list.Node) *list.Node, sz1 int, sz2 int) {
	defer guard_ut(t)

	var head = ramdomLinkList(sz1)
	head = doit(head)
	assert(t, checkLinkList(head, sz1))

	head = stupidLinkList(sz2)
	head = doit(head)
	assert(t, checkLinkList(head, sz2))

	for i := 0; i < 6; i++ {
		head = ramdomLinkList(i)
		head = doit(head)
		assert(t, checkLinkList(head, i))
	}
}
func checkLinkList(head *list.Node, size int) bool {
	if size == 0 {
		return head == nil
	}
	var cnt = 1
	for ; head.Next != nil; head = head.Next {
		if head.Next.Val < head.Val {
			return false
		}
		cnt++
	}
	return cnt == size
}

func ramdomLinkList(size int) *list.Node {
	rand.Seed(time.Now().Unix())
	var head *list.Node
	var tail = list.FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.Next = new(list.Node)
		tail = tail.Next
		tail.Val = rand.Int()
	}
	tail.Next = nil
	return head
}
func stupidLinkList(size int) *list.Node {
	var head *list.Node
	var tail = list.FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.Next = new(list.Node)
		tail = tail.Next
		tail.Val = (size - i) / 2
	}
	tail.Next = nil
	return head
}
