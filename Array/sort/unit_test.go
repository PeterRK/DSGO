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
}
func Test_SimpleSortX(t *testing.T) {
	testArraySort(t, SimpleSort, smallSize, smallSize)
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

func testArraySort(t *testing.T, doit func([]int), sz1 int, sz2 int) {
	defer guardUT(t)

	var list = randArray(sz1)
	var tips = figureOutTips(list)
	doit(list)
	assert(t, checkArrary(list) && tips == figureOutTips(list))

	list = desArray(sz2)
	doit(list)
	assert(t, checkArrary(list))
	for i := 0; i < len(list); i++ {
		list[i] = 99
	}
	doit(list)
	for _, num := range list {
		assert(t, num == 99)
	}

	for i := 0; i < 6; i++ {
		list = randArray(i)
		doit(list)
		assert(t, checkArrary(list))
	}
}
func checkArrary(list []int) bool {
	for i, size := 1, len(list); i < size; i++ {
		if list[i] < list[i-1] {
			return false
		}
	}
	return true
}
func figureOutTips(list []int) int {
	var tips = 0
	for _, num := range list {
		tips ^= num
	}
	return tips
}

func randArray(size int) []int {
	rand.Seed(time.Now().Unix())
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}
func desArray(size int) []int {
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = size - i
	}
	return list
}
func constArray(size int) []int {
	var val = rand.Int()
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = val
	}
	return list
}
