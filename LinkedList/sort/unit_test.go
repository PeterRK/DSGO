package sort

import (
	"linkedlist"
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

func ramdomLinkList(size int) *linkedlist.Node {
	rand.Seed(time.Now().Unix())
	var head *linkedlist.Node
	var tail = linkedlist.FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.Next = new(linkedlist.Node)
		tail = tail.Next
		tail.Val = rand.Int()
	}
	tail.Next = nil
	return head
}
func checkLinkList(head *linkedlist.Node, size int) bool {
	var cnt = 1
	for ; head.Next != nil; head = head.Next {
		if head.Next.Val < head.Val {
			return false
		}
		cnt++
	}
	return cnt == size
}
func testLinkListSort(t *testing.T, doit func(*linkedlist.Node) *linkedlist.Node) {
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
	head = new(linkedlist.Node)
	head.Next = nil
	head = doit(head)
	head = nil
	head = doit(head)
}

func Test_LinkMergeSort(t *testing.T) {
	testLinkListSort(t, MergeSort)
}
func Test_LinkQuickSort(t *testing.T) {
	testLinkListSort(t, QuickSort)
}
