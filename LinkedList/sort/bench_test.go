package sort

import (
	"DSGO/LinkedList/list"
	"testing"
)

func Benchmark_MergeSort(b *testing.B) {
	benchLinkedListSort(b, MergeSort, randLinkedList)
}
func Benchmark_QuickSort(b *testing.B) {
	benchLinkedListSort(b, QuickSort, randLinkedList)
}
func Benchmark_IntroSort(b *testing.B) {
	benchLinkedListSort(b, IntroSort, randLinkedList)
}
func Benchmark_RadixSort(b *testing.B) {
	benchLinkedListSort(b, RadixSort, randLinkedList)
}

func Benchmark_DesMergeSort(b *testing.B) {
	benchLinkedListSort(b, MergeSort, desLinkList)
}
func Benchmark_DesQuickSort(b *testing.B) {
	benchLinkedListSort(b, QuickSort, desLinkList)
}
func Benchmark_DesIntroSort(b *testing.B) {
	benchLinkedListSort(b, IntroSort, desLinkList)
}
func Benchmark_DesRadixSort(b *testing.B) {
	benchLinkedListSort(b, RadixSort, desLinkList)
}

func benchLinkedListSort(b *testing.B,
	doit func(*list.Node) *list.Node, make_array func(int) *list.Node) {
	b.StopTimer()
	var head = make_array(b.N)
	b.StartTimer()
	doit(head)
}
