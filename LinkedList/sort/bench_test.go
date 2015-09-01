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
func Benchmark_IntroSort(b *testing.B) {
	b.StopTimer()
	var head = ramdomLinkList(b.N)
	b.StartTimer()
	IntroSort(head)
}
func Benchmark_RadixSort(b *testing.B) {
	b.StopTimer()
	var head = ramdomLinkList(b.N)
	b.StartTimer()
	RadixSort(head)
}
func Benchmark_ExtremeMergeSort(b *testing.B) {
	b.StopTimer()
	var head = stupidLinkList(b.N)
	b.StartTimer()
	MergeSort(head)
}
func Benchmark_ExtremeQuickSort(b *testing.B) {
	b.StopTimer()
	var head = stupidLinkList(b.N)
	b.StartTimer()
	QuickSort(head)
}
func Benchmark_ExtremeIntroSort(b *testing.B) {
	b.StopTimer()
	var head = stupidLinkList(b.N)
	b.StartTimer()
	IntroSort(head)
}
func Benchmark_ExtremeRadixSort(b *testing.B) {
	b.StopTimer()
	var head = stupidLinkList(b.N)
	b.StartTimer()
	RadixSort(head)
}
