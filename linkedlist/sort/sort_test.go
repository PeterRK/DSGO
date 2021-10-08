package sort

import (
	ll "DSGO/linkedlist"
	"DSGO/utils"
	"math/rand"
	"testing"
	"time"
)

type elem int32

func testSort(t *testing.T,
	doit func(*ll.Node[elem]) *ll.Node[elem], sz1, sz2 int) {
	defer utils.FailInPanic(t)

	head := genRand(sz1)
	trait := getTrait(head)
	head = doit(head)
	utils.Assert(t, IsSorted(head, sz1) && trait == getTrait(head))

	head = genDesc(sz2)
	head = doit(head)
	utils.Assert(t, IsSorted(head, sz2))

	for i := 0; i < 6; i++ {
		head = genRand(i)
		head = doit(head)
		utils.Assert(t, IsSorted(head, i))
	}
	doit(nil)
}

func IsSorted(head *ll.Node[elem], size int) bool {
	if size == 0 {
		return head == nil
	}
	cnt := 1
	for ; head.Next != nil; head = head.Next {
		if head.Val > head.Next.Val {
			return false
		}
		cnt++
	}
	return cnt == size
}

func getTrait(head *ll.Node[elem]) elem {
	trait := elem(0)
	for ; head != nil; head = head.Next {
		trait ^= head.Val
	}
	return trait
}

func genRand(size int) *ll.Node[elem] {
	rand.Seed(time.Now().UnixNano())
	var head *ll.Node[elem]
	tail := ll.FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.Next = new(ll.Node[elem])
		tail = tail.Next
		tail.Val = elem(rand.Uint64())
	}
	tail.Next = nil
	return head
}

func genPseudo(size int) *ll.Node[elem] {
	rand.Seed(999)
	var head *ll.Node[elem]
	tail := ll.FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.Next = new(ll.Node[elem])
		tail = tail.Next
		tail.Val = elem(rand.Uint64())
	}
	tail.Next = nil
	return head
}

func genDesc(size int) *ll.Node[elem] {
	var head *ll.Node[elem]
	tail := ll.FakeHead(&head)
	for i := 0; i < size; i++ {
		tail.Next = new(ll.Node[elem])
		tail = tail.Next
		tail.Val = elem(size - i)
	}
	tail.Next = nil
	return head
}

func benchSort(b *testing.B,
	doit func(*ll.Node[elem]) *ll.Node[elem], genList func(int) *ll.Node[elem]) {
	b.StopTimer()
	head := genList(b.N)
	b.StartTimer()
	doit(head)
}

const bigSize = 500
const smallSize = 100

func Test_MergeSort(t *testing.T) {
	testSort(t, MergeSort[elem], bigSize, smallSize)
}
func Test_QuickSort(t *testing.T) {
	testSort(t, QuickSort[elem], bigSize, smallSize)
}
func Test_IntroSort(t *testing.T) {
	testSort(t, IntroSort[elem], bigSize, bigSize)
}
func Test_RadixSort(t *testing.T) {
	testSort(t, RadixSort[elem], bigSize, smallSize)
}

func Benchmark_MergeSort(b *testing.B) {
	benchSort(b, MergeSort[elem], genPseudo)
}
func Benchmark_QuickSort(b *testing.B) {
	benchSort(b, QuickSort[elem], genPseudo)
}
func Benchmark_IntroSort(b *testing.B) {
	benchSort(b, IntroSort[elem], genPseudo)
}
func Benchmark_RadixSort(b *testing.B) {
	benchSort(b, RadixSort[elem], genPseudo)
}

func Benchmark_DescMergeSort(b *testing.B) {
	benchSort(b, MergeSort[elem], genDesc)
}
func Benchmark_DescQuickSort(b *testing.B) {
	benchSort(b, QuickSort[elem], genDesc)
}
func Benchmark_DescIntroSort(b *testing.B) {
	benchSort(b, IntroSort[elem], genDesc)
}
func Benchmark_DescRadixSort(b *testing.B) {
	benchSort(b, RadixSort[elem], genDesc)
}
