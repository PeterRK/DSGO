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

	var roads = [][]int{
		{1, 5, 6},    //0
		nil,          //1
		{0, 3},       //2
		{5},          //3
		nil,          //4
		{4},          //5
		{4, 9},       //6
		{6},          //7
		{7},          //8
		{10, 11, 12}, //9
		nil,          //10
		{12},         //11
		nil,          //12
	}

	var expected = []int{8, 7, 2, 3, 0, 6, 9, 11, 12, 10, 5, 4, 1}
	var vec, err = TopologicalSort(roads)
	assert(t, err == nil && isTheSame(vec, expected))
}

func Test_SplitDirectedGraph(t *testing.T) {
	defer guardUT(t)

	var roads = [][]int{
		{1},       //0
		{2, 3, 4}, //1
		{5},       //2
		nil,       //3
		{1, 5, 6}, //4
		{2, 7},    //5
		{7, 9},    //6
		{10},      //7
		{6},       //8
		{8},       //9
		{11},      //10
		{9},       //11
	}

	var expected = [][]int{{0}, {4, 1}, {3}, {5, 2}, {10, 11, 9, 8, 6, 7}}
	var parts = SplitDirectedGraph(roads)
	assert(t, len(parts) == len(expected))
	for i := 0; i < len(parts); i++ {
		assert(t, isTheSame(parts[i], expected[i]))
	}
}
