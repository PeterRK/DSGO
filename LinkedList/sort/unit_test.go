package sort

import (
	"LinkedList/list"
	"math/rand"
	"testing"
	"time"
)

const sz_tiny = 5
const sz_small = 300
const sz_big = 10000

func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
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
func testLinkListSort(t *testing.T, doit func(*list.Node) *list.Node) {
	defer guard_ut(t)
	var head = ramdomLinkList(sz_small)
	head = doit(head)
	if !checkLinkList(head, sz_small) {
		t.Fail()
	}
	head = ramdomLinkList(sz_tiny)
	head = doit(head)
	if !checkLinkList(head, sz_tiny) {
		t.Fail()
	}
	head = new(list.Node)
	head.Next = nil
	head = doit(head)
	head = nil
	head = doit(head)
}

func Test_MergeSort(t *testing.T) {
	testLinkListSort(t, MergeSort)
}
func Test_QuickSort(t *testing.T) {
	testLinkListSort(t, QuickSort)
}
func Test_IntroSort(t *testing.T) {
	testLinkListSort(t, IntroSort)
}
