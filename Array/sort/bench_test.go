package sort

import (
	"testing"
)

/*	这几个要跑很久
func Benchmark_BubleSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	BubleSort(list)
}
func Benchmark_SelectSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	SelectSort(list)
}
func Benchmark_InsertSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	InsertSort(list)
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
func Benchmark_IntroSort(b *testing.B) {
	b.StopTimer()
	var list = ramdomArray(b.N)
	b.StartTimer()
	IntroSort(list)
}
