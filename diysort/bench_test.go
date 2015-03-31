package diysort

import (
	"testing"
)

/*	这两个要跑很久
func Benchmark_InsertSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	InsertSort(list)
}
func Benchmark_SelectSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	SelectSort(list)
}
*/

func Benchmark_HeapSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	HeapSort(list)
}
func Benchmark_MergeSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	MergeSort(list)
}
func Benchmark_QuickSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	QuickSort(list)
}
func Benchmark_Introsort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	Introsort(list)
}

func Benchmark_LinkMergeSort(b *testing.B) {
	b.StopTimer()
	var head = ramdomLinkList(b.N)
	b.StartTimer()
	LinkMergeSort(head)
}

func Benchmark_LinkQuickSort(b *testing.B) {
	b.StopTimer()
	var head = ramdomLinkList(b.N)
	b.StartTimer()
	LinkQuickSort(head)
}
