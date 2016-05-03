package graph

import (
	"testing"
)

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
	defer guardUT(t)

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

func Test_SplitDirectedGraph(t *testing.T) {
	defer guardUT(t)

	var roads = [][]int{
		{1},
		{2, 3, 4},
		{5},
		{},
		{1, 5, 6},
		{2, 7},
		{7, 9},
		{10},
		{6},
		{8},
		{11},
		{9}}

	var expected = [][]int{{0}, {4, 1}, {3}, {5, 2}, {10, 11, 9, 8, 6, 7}}
	var parts = SplitDirectedGraph(roads)
	assert(t, len(parts) == len(expected))
	for i := 0; i < len(parts); i++ {
		assert(t, isTheSame(parts[i], expected[i]))
	}
}
