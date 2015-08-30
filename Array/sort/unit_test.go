package sort

import (
	"math/rand"
	"testing"
	"time"
)

const sz_big = 2000
const sz_small = 300

func Test_BubleSort(t *testing.T) {
	testArraySort(t, sz_small, ramdomArray, BubleSort)
	//testArraySort(t, sz_small, stupidArray, BubleSort)
}
func Test_SelectSort(t *testing.T) {
	testArraySort(t, sz_small, ramdomArray, SelectSort)
	//testArraySort(t, sz_small, stupidArray, SelectSort)
}
func Test_InsertSort(t *testing.T) {
	testArraySort(t, sz_small, ramdomArray, InsertSort)
	//testArraySort(t, sz_small, stupidArray, InsertSort)
}
func Test_HeapSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, HeapSort)
	//testArraySort(t, sz_small, stupidArray, HeapSort)
}
func Test_MergeSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, MergeSort)
	//testArraySort(t, sz_small, stupidArray, MergeSort)
}
func Test_QuickSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, QuickSort)
	//testArraySort(t, sz_small, stupidArray, QuickSort)
}
func Test_IntroSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, IntroSort)
	//testArraySort(t, sz_big, stupidArray, IntroSort)
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

func testArraySort(t *testing.T, size int, create func(int) []int, doit func([]int)) {
	defer guard_ut(t)

	var list = create(size)
	doit(list)
	assert(t, checkArrary(list))
	list = create(5)
	doit(list)
	assert(t, checkArrary(list))

	list = []int{0}
	doit(list)
	list = list[:0]
	doit(list)
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
