package sort

import (
	"math/rand"
	"testing"
	"time"
)

const sz_big = 2000
const sz_small = 300

func Test_BubleSort(t *testing.T) {
	testArraySort(t, BubleSort, sz_small, sz_small)
}
func Test_SelectSort(t *testing.T) {
	testArraySort(t, SelectSort, sz_small, sz_small)
}
func Test_InsertSort(t *testing.T) {
	testArraySort(t, InsertSort, sz_small, sz_small)
}
func Test_HeapSort(t *testing.T) {
	testArraySort(t, HeapSort, sz_big, sz_small)
}
func Test_MergeSort(t *testing.T) {
	testArraySort(t, MergeSort, sz_big, sz_small)
}
func Test_QuickSort(t *testing.T) {
	testArraySort(t, QuickSort, sz_big, sz_small)
}
func Test_IntroSort(t *testing.T) {
	testArraySort(t, IntroSort, sz_big, sz_big)
}
func Test_RadixSort(t *testing.T) {
	testArraySort(t, RadixSort, sz_big, sz_small)
}

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func testArraySort(t *testing.T, doit func([]int), sz1 int, sz2 int) {
	defer guard_ut(t)

	var list = ramdomArray(sz1)
	doit(list)
	assert(t, checkArrary(list))

	list = stupidArray(sz2)
	doit(list)
	assert(t, checkArrary(list))

	for i := 0; i < 6; i++ {
		list = ramdomArray(i)
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

func ramdomArray(size int) []int {
	rand.Seed(time.Now().Unix())
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}
func stupidArray(size int) []int {
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = (size - i) / 2
	}
	return list
}
