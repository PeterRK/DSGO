package sort

import (
	"LinkedList/list"
	"math/rand"
	"testing"
	"time"
)

func Test_MergeSort(t *testing.T) {
	testLinkListSort(t, MergeSort)
}
func Test_QuickSort(t *testing.T) {
	testLinkListSort(t, QuickSort)
}
func Test_IntroSort(t *testing.T) {
	testLinkListSort(t, IntroSort)
}

func testLinkListSort(t *testing.T, doit func(*list.Node) *list.Node) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	var head = ramdomLinkList(200)
	head = doit(head)
	if !checkLinkList(head, 200) {
		t.Fail()
	}
	head = ramdomLinkList(5)
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
