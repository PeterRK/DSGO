package graph

import (
	"testing"
)

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

func isTheSame(vec1 []int, vec2 []int) bool {
	if len(vec1) != len(vec2) {
		return false
	}
	for i, size := 0, len(vec1); i < size; i++ {
		if vec1[i] != vec2[i] {
			return false
		}
	}
	return true
}

func Test_TopologicalSort(t *testing.T) {
	defer guard_ut(t)

	var roads = make([][]int, 13)
	roads[0] = []int{1, 5, 6}
	roads[1] = []int{}
	roads[2] = []int{0, 3}
	roads[3] = []int{5}
	roads[4] = []int{}
	roads[5] = []int{4}
	roads[6] = []int{4, 9}
	roads[7] = []int{6}
	roads[8] = []int{7}
	roads[9] = []int{10, 11, 12}
	roads[10] = []int{}
	roads[11] = []int{12}
	roads[12] = []int{}

	var expected = []int{8, 7, 2, 3, 0, 6, 9, 11, 12, 10, 5, 4, 1}
	var vec, err = TopologicalSort(roads)
	assert(t, err == nil && isTheSame(vec, expected))
}
