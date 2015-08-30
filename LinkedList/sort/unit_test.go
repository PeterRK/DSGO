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
	testLinkListSort(t, sz_big, ramdomLinkList, MergeSort)
	//testLinkListSort(t, sz_small, stupidLinkList, MergeSort)
}
func Test_QuickSort(t *testing.T) {
	testLinkListSort(t, sz_big, ramdomLinkList, QuickSort)
	//testLinkListSort(t, sz_small, stupidLinkList, QuickSort)
}
func Test_IntroSort(t *testing.T) {
	testLinkListSort(t, sz_big, ramdomLinkList, IntroSort)
	//testLinkListSort(t, sz_small, stupidLinkList, IntroSort)
}

func testLinkListSort(t *testing.T, size int, create func(int) *list.Node, doit func(*list.Node) *list.Node) {
	defer guard_ut(t)

	var head = create(size)
	head = doit(head)
	assert(t, checkLinkList(head, size))
	head = create(5)
	head = doit(head)
	assert(t, checkLinkList(head, 5))

	head = new(list.Node)
	head.Next = nil
	head = doit(head)
	assert(t, head != nil && head.Next == nil)
	head = nil
	assert(t, doit(head) == nil)
}
func checkLinkList(head *list.Node, size int) bool {
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
		tail.Val = i / 2
	}
	tail.Next = nil
	return head
}
