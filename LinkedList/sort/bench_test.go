package sort

import (
	"testing"
)

func Benchmark_MergeSort(b *testing.B) {
	b.StopTimer()
	var head = ramdomLinkList(b.N)
	b.StartTimer()
	MergeSort(head)
}
func Benchmark_QuickSort(b *testing.B) {
	b.StopTimer()
	var head = ramdomLinkList(b.N)
	b.StartTimer()
	QuickSort(head)
}
