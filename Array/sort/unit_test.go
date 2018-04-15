package sort

import (
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
