package sort

import (
	"testing"
)

/*	这几个要跑很久
func Benchmark_BubleSort(b *testing.B) {
	benchArraySort(b, BubleSort, randArray)
}
func Benchmark_SelectSort(b *testing.B) {
	benchArraySort(b, SelectSort, randArray)
}
func Benchmark_InsertSort(b *testing.B) {
	benchArraySort(b, InsertSort, randArray)
}
*/
func Benchmark_HeapSort(b *testing.B) {
	benchArraySort(b, HeapSort, randArray)
}
func Benchmark_MergeSort(b *testing.B) {
	benchArraySort(b, MergeSort, randArray)
}
func Benchmark_QuickSort(b *testing.B) {
	benchArraySort(b, QuickSort, randArray)
}
func Benchmark_QuickSortY(b *testing.B) {
	benchArraySort(b, QuickSortY, randArray)
}
func Benchmark_IntroSort(b *testing.B) {
	benchArraySort(b, IntroSort, randArray)
}
func Benchmark_IntroSortY(b *testing.B) {
	benchArraySort(b, IntroSortY, randArray)
}
func Benchmark_RadixSort(b *testing.B) {
	benchArraySort(b, RadixSort, randArray)
}

func Benchmark_DesHeapSort(b *testing.B) {
	benchArraySort(b, HeapSort, desArray)
}
func Benchmark_DesMergeSort(b *testing.B) {
	benchArraySort(b, MergeSort, desArray)
}
func Benchmark_DesQuickSort(b *testing.B) {
	benchArraySort(b, QuickSort, desArray)
}
func Benchmark_DesQuickSortY(b *testing.B) {
	benchArraySort(b, QuickSortY, desArray)
}
func Benchmark_DesIntroSort(b *testing.B) {
	benchArraySort(b, IntroSort, desArray)
}
func Benchmark_DesIntroSortY(b *testing.B) {
	benchArraySort(b, IntroSortY, desArray)
}
func Benchmark_DesRadixSort(b *testing.B) {
	benchArraySort(b, RadixSort, desArray)
}

func Benchmark_ConstHeapSort(b *testing.B) {
	benchArraySort(b, HeapSort, constArray)
}
func Benchmark_ConstMergeSort(b *testing.B) {
	benchArraySort(b, MergeSort, constArray)
}
func Benchmark_ConstQuickSort(b *testing.B) {
	benchArraySort(b, QuickSort, constArray)
}
func Benchmark_ConstIntroSort(b *testing.B) {
	benchArraySort(b, IntroSort, constArray)
}
func Benchmark_ConstRadixSort(b *testing.B) {
	benchArraySort(b, RadixSort, constArray)
}

// 太快，Benchmark框架崩溃
//func Benchmark_ConstQuickSortY(b *testing.B) {
//	benchArraySort(b, QuickSortY, constArray)
//}
//func Benchmark_ConstIntroSortY(b *testing.B) {
//	benchArraySort(b, IntroSortY, constArray)
//}

func benchArraySort(b *testing.B,
	doit func([]int), make_array func(int) []int) {
	b.StopTimer()
	var list = make_array(b.N)
	b.StartTimer()
	doit(list)
}
