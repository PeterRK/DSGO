package sort

import (
	"DSGO/LinkedList/list"
	"math/rand"
	"testing"
	"time"
)

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guardUT(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

const bigSize = 500
const smallSize = 100

func Test_MergeSort(t *testing.T) {
	testLinkListSort(t, MergeSort, bigSize, smallSize)
}
func Test_QuickSort(t *testing.T) {
	testLinkListSort(t, QuickSort, bigSize, smallSize)
}
func Test_IntroSort(t *testing.T) {
	testLinkListSort(t, IntroSort, bigSize, bigSize)
}
func Test_RadixSort(t *testing.T) {
	testLinkListSort(t, RadixSort, bigSize, smallSize)
}

func testLinkListSort(t *testing.T,
	doit func(*list.Node) *list.Node, sz1 int, sz2 int) {
	defer guardUT(t)

	var head = randLinkedList(sz1)
	var tips = figureOutTips(head)
	head = doit(head)
	assert(t, checkLinkList(head, sz1) && tips == figureOutTips(head))

	head = desLinkList(sz2)
	head = doit(head)
	assert(t, checkLinkList(head, sz2))

	for i := 0; i < 6; i++ {
		head = randLinkedList(i)
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
func figureOutTips(head *list.Node) int {
	var tips = 0
	for ; head != nil; head = head.Next {
		tips ^= head.Val
	}
	return tips
}

func randLinkedList(size int) *list.Node {
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
func desLinkList(size int) *list.Node {
	var head *list.Node
	var tail = list.FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.Next = new(list.Node)
		tail = tail.Next
		tail.Val = size - i
	}
	tail.Next = nil
	return head
}
