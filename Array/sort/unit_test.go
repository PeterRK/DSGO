package sort

import (
	//"math"
	"math/rand"
	"testing"
	"time"
)

const bigSize = 2000
const smallSize = 300

func Test_BubleSort(t *testing.T) {
	testArraySort(t, BubleSort, smallSize, smallSize)
}
func Test_SelectSort(t *testing.T) {
	testArraySort(t, SelectSort, smallSize, smallSize)
}
func Test_InsertSort(t *testing.T) {
	testArraySort(t, InsertSort, smallSize, smallSize)
}
func Test_SimpleSort(t *testing.T) {
	testArraySort(t, SimpleSort, smallSize, smallSize)
	testArraySort(t, SimpleSortX, smallSize, smallSize)
}
func Test_HeapSort(t *testing.T) {
	testArraySort(t, HeapSort, bigSize, smallSize)
}
func Test_MergeSort(t *testing.T) {
	testArraySort(t, MergeSort, bigSize, smallSize)
}
func Test_QuickSort(t *testing.T) {
	testArraySort(t, QuickSort, bigSize, smallSize)
}
func Test_QuickSortY(t *testing.T) {
	testArraySort(t, QuickSortY, bigSize, smallSize)
}

//func Test_QuickSortS(t *testing.T) {
//	testArraySort(t, QuickSortS, bigSize, smallSize)
//}

func Test_IntroSort(t *testing.T) {
	testArraySort(t, IntroSort, bigSize, bigSize)
}
func Test_IntroSortY(t *testing.T) {
	testArraySort(t, IntroSortY, bigSize, bigSize)
}
func Test_RadixSort(t *testing.T) {
	testArraySort(t, RadixSort, bigSize, smallSize)
}

func Test_reorder(t *testing.T) {
	defer guardUT(t)

	const size = smallSize
	list := make([]Unit, size)
	index := make([]uint32, size)
	for i := 0; i < size; i++ {
		list[i].val = -i
		index[i] = uint32(i)
	}
	rand.Seed(time.Now().Unix())
	for i := 1; i < size; i++ {
		j := rand.Int() % (i + 1)
		index[i], index[j] = index[j], index[i]
	}

	memo := make([]uint32, size)
	for i := 1; i < size; i++ {
		memo[i] = index[i]
	}

	reorder(list, index)
	for i := 1; i < size; i++ {
		assert(t, list[i].val == -int(memo[i]))
	}

	reorder(nil, nil)
}

func Test_IndexIntroSort(t *testing.T) {
	testArraySort(t, IndexIntroSort, smallSize, smallSize)

	/*
		defer guardUT(t)
		list := randArray(bigSize)
		injectConst(list, 10, math.MinInt32)
		injectConst(list, 10, math.MaxInt32)
		injectConst(list, 10, -1)
		injectConst(list, 10, 1)
		IndexIntroSort(list)
		assert(t, checkArraryX(list))
	*/
}

/*
func injectConst(list []Unit, cnt int, val int) {
	for i := 0; i < cnt; i++ {
		j := rand.Int() % len(list)
		list[j].val = val
		list[j].pad[0] = uint32(j)
	}
}
func checkArraryX(list []Unit) bool {
	for i, size := 1, len(list); i < size; i++ {
		if list[i].val < list[i-1].val {
			return false
		} else if list[i].val == list[i-1].val &&
			list[i].pad[0] < list[i-1].pad[0] {
			return false
		}
	}
	return true
}
*/

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

func testArraySort(t *testing.T, doit func([]Unit), sz1 int, sz2 int) {
	defer guardUT(t)

	list := randArray(sz1)
	tips := figureOutTips(list)
	doit(list)
	assert(t, checkArrary(list) && tips == figureOutTips(list))

	list = desArray(sz2)
	doit(list)
	assert(t, checkArrary(list))
	for i := 0; i < len(list); i++ {
		list[i].val = 99
	}
	doit(list)
	for i := 0; i < len(list); i++ {
		assert(t, list[i].val == 99)
	}

	for i := 0; i < 6; i++ {
		list = randArray(i)
		doit(list)
		assert(t, checkArrary(list))
	}
}
func checkArrary(list []Unit) bool {
	for i, size := 1, len(list); i < size; i++ {
		if list[i].val < list[i-1].val {
			return false
		}
	}
	return true
}
func figureOutTips(list []Unit) int {
	tips := 0
	for i := 0; i < len(list); i++ {
		tips ^= list[i].val
	}
	return tips
}

func randArray(size int) []Unit {
	rand.Seed(time.Now().Unix())
	list := make([]Unit, size)
	for i := 0; i < size; i++ {
		list[i].val = rand.Int()
	}
	return list
}
func desArray(size int) []Unit {
	list := make([]Unit, size)
	for i := 0; i < size; i++ {
		list[i].val = size - i
	}
	return list
}
func constArray(size int) []Unit {
	val := rand.Int()
	list := make([]Unit, size)
	for i := 0; i < size; i++ {
		list[i].val = val
	}
	return list
}
