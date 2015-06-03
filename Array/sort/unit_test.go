//1234567890
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
	testArraySort(t, sz_small, stupidArrayA, BubleSort)
	testArraySort(t, sz_small, stupidArrayB, BubleSort)
}
func Test_SelectSort(t *testing.T) {
	testArraySort(t, sz_small, ramdomArray, SelectSort)
	testArraySort(t, sz_small, stupidArrayA, SelectSort)
	testArraySort(t, sz_small, stupidArrayB, SelectSort)
}
func Test_InsertSort(t *testing.T) {
	testArraySort(t, sz_small, ramdomArray, InsertSort)
	testArraySort(t, sz_small, stupidArrayA, InsertSort)
	testArraySort(t, sz_small, stupidArrayB, InsertSort)
}
func Test_HeapSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, HeapSort)
	testArraySort(t, sz_small, stupidArrayA, HeapSort)
	testArraySort(t, sz_small, stupidArrayB, HeapSort)
}
func Test_MergeSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, MergeSort)
	testArraySort(t, sz_small, stupidArrayA, MergeSort)
	testArraySort(t, sz_small, stupidArrayB, MergeSort)
}
func Test_QuickSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, QuickSort)
	testArraySort(t, sz_small, stupidArrayA, QuickSort)
	testArraySort(t, sz_small, stupidArrayB, QuickSort)
}
func Test_IntroSort(t *testing.T) {
	testArraySort(t, sz_big, ramdomArray, IntroSort)
	testArraySort(t, sz_big, stupidArrayA, IntroSort)
	testArraySort(t, sz_big, stupidArrayB, IntroSort)
}

func testArraySort(t *testing.T, size int, create func(int) []int, doit func([]int)) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	var list = create(size)
	doit(list)
	if !checkArrary(list) {
		t.Fail()
	}
	list = create(5)
	doit(list)
	if !checkArrary(list) {
		t.Fail()
	}
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
func stupidArrayA(size int) []int {
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i / 2
	}
	return list
}
func stupidArrayB(size int) []int {
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = (size - i) / 2
	}
	return list
}
