package sort

import (
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
func testLinkListSort(t *testing.T, doit func(*Node) *Node) {
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
	head = new(Node)
	head.next = nil
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
