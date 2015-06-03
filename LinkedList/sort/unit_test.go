package sort

import (
	"LinkedList/list"
	"math/rand"
	"testing"
	"time"
)

const sz_big = 500
const sz_small = 100

func Test_MergeSort(t *testing.T) {
	testLinkListSort(t, sz_big, ramdomLinkList, MergeSort)
	testLinkListSort(t, sz_small, stupidLinkListA, MergeSort)
	testLinkListSort(t, sz_small, stupidLinkListB, MergeSort)
}
func Test_QuickSort(t *testing.T) {
	testLinkListSort(t, sz_big, ramdomLinkList, QuickSort)
	testLinkListSort(t, sz_small, stupidLinkListA, QuickSort)
	testLinkListSort(t, sz_small, stupidLinkListB, QuickSort)
}
func Test_IntroSort(t *testing.T) {
	testLinkListSort(t, sz_big, ramdomLinkList, IntroSort)
	testLinkListSort(t, sz_small, stupidLinkListA, IntroSort)
	testLinkListSort(t, sz_small, stupidLinkListB, IntroSort)
}

func testLinkListSort(t *testing.T, size int, create func(int) *list.Node, doit func(*list.Node) *list.Node) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	var head = create(size)
	head = doit(head)
	if !checkLinkList(head, size) {
		t.Fail()
	}
	head = create(5)
	head = doit(head)
	if !checkLinkList(head, 5) {
		t.Fail()
	}
	head = new(list.Node)
	head.Next = nil
	head = doit(head)
	head = nil
	head = doit(head)
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
func stupidLinkListA(size int) *list.Node {
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
func stupidLinkListB(size int) *list.Node {
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
