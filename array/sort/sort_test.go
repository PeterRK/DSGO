package sort

import (
	"DSGO/utils"
	"testing"
)

type elem uint

func testSort(t *testing.T, doit func([]elem), sz1, sz2 int) {
	defer utils.FailInPanic(t)

	list := genRand(sz1)
	trait := getTrait(list)
	doit(list)
	utils.Assert(t, IsSorted(list) && trait == getTrait(list))

	list = genDesc(sz2)
	doit(list)
	utils.Assert(t, IsSorted(list))
	for i := 0; i < len(list); i++ {
		list[i] = 99
	}
	doit(list)
	for i := 0; i < len(list); i++ {
		utils.Assert(t, list[i] == 99)
	}

	for i := 0; i < 6; i++ {
		list = genRand(i)
		doit(list)
		utils.Assert(t, IsSorted(list))
	}
}

func getTrait(list []elem) elem {
	trait := elem(0)
	for i := 0; i < len(list); i++ {
		trait ^= list[i]
	}
	return trait
}

func genRand(size int) []elem {
	return utils.RandomArray[elem](size)
}

func genPseudo(size int) []elem {
	return utils.PseudoRandomArray[elem](size, 999)
}

func genDesc(size int) []elem {
	list := make([]elem, size)
	for i := 0; i < size; i++ {
		list[i] = elem(size - i)
	}
	return list
}

func genConst(size int) []elem {
	list := make([]elem, size)
	for i := 0; i < size; i++ {
		list[i] = 12345
	}
	return list
}

func benchSort(b *testing.B,
	doit func([]elem), genArray func(int) []elem) {
	b.StopTimer()
	list := genArray(b.N)
	b.StartTimer()
	doit(list)
}

const bigSize = 2000
const smallSize = 300

func Test_BubleSort(t *testing.T) {
	testSort(t, BubleSort[elem], smallSize, smallSize)
}
func Test_SelectSort(t *testing.T) {
	testSort(t, SelectSort[elem], smallSize, smallSize)
}
func Test_InsertSort(t *testing.T) {
	testSort(t, InsertSort[elem], smallSize, smallSize)
}
func Test_SimpleSort(t *testing.T) {
	testSort(t, SimpleSort[elem], smallSize, smallSize)
}
func Test_HeapSort(t *testing.T) {
	testSort(t, HeapSort[elem], bigSize, smallSize)
}
func Test_MergeSort(t *testing.T) {
	testSort(t, MergeSort[elem], bigSize, smallSize)
}
func Test_SymMergeSort(t *testing.T) {
	testSort(t, SymMergeSort[elem], bigSize, smallSize)
}
func Test_QuickSort(t *testing.T) {
	testSort(t, QuickSort[elem], bigSize, smallSize)
}
func Test_QuickSortY(t *testing.T) {
	testSort(t, QuickSortY[elem], bigSize, smallSize)
}
func Test_BlockQuickSort(t *testing.T) {
	testSort(t, BlockQuickSort[elem], bigSize, smallSize)
}
func Test_IntroSort(t *testing.T) {
	testSort(t, IntroSort[elem], bigSize, bigSize)
}
func Test_IntroSortY(t *testing.T) {
	testSort(t, IntroSortY[elem], bigSize, bigSize)
}
func Test_BlockIntroSort(t *testing.T) {
	testSort(t, BlockIntroSort[elem], bigSize, bigSize)
}
func Test_RadixSort(t *testing.T) {
	testSort(t, RadixSort[elem], bigSize, smallSize)
}

/*	这几个要跑很久
func Benchmark_BubleSort(b *testing.B) {
	benchSort(b, BubleSort[elem], genPseudo)
}
func Benchmark_SelectSort(b *testing.B) {
	benchSort(b, SelectSort[elem], genPseudo)
}
func Benchmark_InsertSort(b *testing.B) {
	benchSort(b, InsertSort[elem], genPseudo)
}
*/
func Benchmark_HeapSort(b *testing.B) {
	benchSort(b, HeapSort[elem], genPseudo)
}
func Benchmark_MergeSort(b *testing.B) {
	benchSort(b, MergeSort[elem], genPseudo)
}
func Benchmark_SymMergeSort(b *testing.B) {
	benchSort(b, SymMergeSort[elem], genPseudo)
}
func Benchmark_QuickSort(b *testing.B) {
	benchSort(b, QuickSort[elem], genPseudo)
}
func Benchmark_QuickSortY(b *testing.B) {
	benchSort(b, QuickSortY[elem], genPseudo)
}
func Benchmark_BlockQuickSort(b *testing.B) {
	benchSort(b, BlockQuickSort[elem], genPseudo)
}
func Benchmark_IntroSort(b *testing.B) {
	benchSort(b, IntroSort[elem], genPseudo)
}
func Benchmark_IntroSortY(b *testing.B) {
	benchSort(b, IntroSortY[elem], genPseudo)
}
func Benchmark_BlockIntroSort(b *testing.B) {
	benchSort(b, BlockIntroSort[elem], genPseudo)
}
func Benchmark_RadixSort(b *testing.B) {
	benchSort(b, RadixSort[elem], genPseudo)
}

func Benchmark_DescHeapSort(b *testing.B) {
	benchSort(b, HeapSort[elem], genDesc)
}
func Benchmark_DescMergeSort(b *testing.B) {
	benchSort(b, MergeSort[elem], genDesc)
}
func Benchmark_DescQuickSort(b *testing.B) {
	benchSort(b, QuickSort[elem], genDesc)
}
func Benchmark_DescQuickSortY(b *testing.B) {
	benchSort(b, QuickSortY[elem], genDesc)
}
func Benchmark_DescBlockQuickSort(b *testing.B) {
	benchSort(b, BlockQuickSort[elem], genDesc)
}
func Benchmark_DescIntroSort(b *testing.B) {
	benchSort(b, IntroSort[elem], genDesc)
}
func Benchmark_DescIntroSortY(b *testing.B) {
	benchSort(b, IntroSortY[elem], genDesc)
}
func Benchmark_DescBlockIntroSort(b *testing.B) {
	benchSort(b, BlockIntroSort[elem], genDesc)
}
func Benchmark_DescRadixSort(b *testing.B) {
	benchSort(b, RadixSort[elem], genDesc)
}

func Benchmark_ConstHeapSort(b *testing.B) {
	benchSort(b, HeapSort[elem], genConst)
}
func Benchmark_ConstMergeSort(b *testing.B) {
	benchSort(b, MergeSort[elem], genConst)
}
func Benchmark_ConstQuickSort(b *testing.B) {
	benchSort(b, QuickSort[elem], genConst)
}
func Benchmark_ConstQuickSortY(b *testing.B) {
	benchSort(b, QuickSortY[elem], genConst)
}
func Benchmark_ConstBlockQuickSort(b *testing.B) {
	benchSort(b, BlockQuickSort[elem], genConst)
}
func Benchmark_ConstIntroSort(b *testing.B) {
	benchSort(b, IntroSort[elem], genConst)
}
func Benchmark_ConstIntroSortY(b *testing.B) {
	benchSort(b, IntroSortY[elem], genConst)
}
func Benchmark_ConstBlockIntroSort(b *testing.B) {
	benchSort(b, BlockIntroSort[elem], genConst)
}
func Benchmark_ConstRadixSort(b *testing.B) {
	benchSort(b, RadixSort[elem], genConst)
}
