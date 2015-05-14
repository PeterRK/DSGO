package sort

import (
	"math/rand"
	"testing"
	"time"
)

const sz_tiny = 5
const sz_small = 300
const sz_big = 10000

func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func ramdomArray(size int) []int {
	rand.Seed(time.Now().Unix())
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}
func checkArrary(list []int) bool {
	for i, size := 1, len(list); i < size; i++ {
		if list[i] < list[i-1] {
			return false
		}
	}
	return true
}
func testArraySort(t *testing.T, size int, doit func([]int)) {
	defer guard_ut(t)
	var list = ramdomArray(size)
	doit(list)
	if !checkArrary(list) {
		t.Fail()
	}
	list = ramdomArray(sz_tiny)
	doit(list)
	if !checkArrary(list) {
		t.Fail()
	}
	list = []int{0}
	doit(list)
	list = list[:0]
	doit(list)
}

func Test_BubleSort(t *testing.T) {
	testArraySort(t, sz_small, BubleSort)
}
func Test_SelectSort(t *testing.T) {
	testArraySort(t, sz_small, SelectSort)
}
func Test_InsertSort(t *testing.T) {
	testArraySort(t, sz_small, InsertSort)
}
func Test_HeapSort(t *testing.T) {
	testArraySort(t, sz_big, HeapSort)
}
func Test_MergeSort(t *testing.T) {
	testArraySort(t, sz_big, MergeSort)
}
func Test_QuickSort(t *testing.T) {
	testArraySort(t, sz_big, QuickSort)
}
func Test_Introsort(t *testing.T) {
	testArraySort(t, sz_big, Introsort)
}
