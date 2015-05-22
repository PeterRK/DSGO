package prim

import (
	"Graph/graph"
	"testing"
)

func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func isTheSame(mtx1 [][]int, mtx2 [][]int) bool {
	if len(mtx1) != len(mtx2) {
		return false
	}
	for i, size := 0, len(mtx1); i < size; i++ {
		if len(mtx1[i]) != len(mtx2[i]) {
			return false
		}
		for j, sz := 0, len(mtx1[i]); j < sz; j++ {
			if mtx1[i][j] != mtx2[i][j] {
				return false
			}
		}
	}
	return true
}

func Test_Prim(t *testing.T) {
	defer guard_ut(t)
	var roads = make([][]graph.Path, 9)
	roads[0] = []graph.Path{{1, 8}, {3, 4}, {6, 11}}
	roads[1] = []graph.Path{{0, 8}, {2, 7}, {4, 2}, {8, 4}}
	roads[2] = []graph.Path{{1, 7}, {5, 9}, {8, 14}}
	roads[3] = []graph.Path{{0, 4}, {6, 8}}
	roads[4] = []graph.Path{{1, 2}, {6, 7}, {7, 6}}
	roads[5] = []graph.Path{{2, 9}, {8, 10}}
	roads[6] = []graph.Path{{0, 11}, {3, 8}, {4, 7}, {7, 1}}
	roads[7] = []graph.Path{{4, 6}, {6, 1}, {8, 2}}
	roads[8] = []graph.Path{{1, 4}, {2, 14}, {5, 10}, {7, 2}}

	var dist, fail = Prim(roads)
	if fail || dist != 37 {
		t.Fail()
	}
}

func Test_PlainPrim(t *testing.T) {
	defer guard_ut(t)
	var matrix = [][]uint{
		{0, 8, 0, 4, 0, 0, 11, 0, 0},
		{8, 0, 7, 0, 2, 0, 0, 0, 4},
		{0, 7, 0, 0, 0, 9, 0, 0, 14},
		{4, 0, 0, 0, 0, 0, 8, 0, 0},
		{0, 2, 0, 0, 0, 0, 7, 6, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 10},
		{11, 0, 0, 8, 7, 0, 0, 1, 0},
		{0, 0, 0, 0, 6, 0, 1, 0, 2},
		{0, 4, 14, 0, 0, 10, 0, 2, 0}}

	var dist, fail = PlainPrim(matrix)
	if fail || dist != 37 {
		t.Fail()
	}
}

func Test_Kruskal(t *testing.T) {
	defer guard_ut(t)
	var roads = []graph.PathX{
		{0, 1, 8},
		{0, 3, 4},
		{0, 6, 11},
		{1, 2, 7},
		{1, 4, 2},
		{1, 8, 4},
		{2, 5, 9},
		{2, 8, 14},
		{3, 6, 8},
		{4, 6, 7},
		{4, 7, 6},
		{5, 8, 10},
		{6, 7, 1},
		{7, 8, 2}}

	var dist, fail = Kruskal(roads, 9)
	if fail || dist != 37 {
		t.Fail()
	}
}

func Test_PrimTree(t *testing.T) {
	defer guard_ut(t)
	var roads = make([][]graph.Path, 9)
	roads[0] = []graph.Path{{1, 8}, {3, 4}, {6, 11}}
	roads[1] = []graph.Path{{0, 8}, {2, 7}, {4, 2}, {8, 4}}
	roads[2] = []graph.Path{{1, 7}, {5, 9}, {8, 14}}
	roads[3] = []graph.Path{{0, 4}, {6, 8}}
	roads[4] = []graph.Path{{1, 2}, {6, 7}, {7, 6}}
	roads[5] = []graph.Path{{2, 9}, {8, 10}}
	roads[6] = []graph.Path{{0, 11}, {3, 8}, {4, 7}, {7, 1}}
	roads[7] = []graph.Path{{4, 6}, {6, 1}, {8, 2}}
	roads[8] = []graph.Path{{1, 4}, {2, 14}, {5, 10}, {7, 2}}

	var expected = make([][]int, 9)
	expected[0] = []int{3, 1}
	expected[1] = []int{4, 8, 2}
	expected[2] = []int{5}
	expected[7] = []int{6}
	expected[8] = []int{7}

	var ret, fail = PrimTree(roads)
	if fail || !isTheSame(ret, expected) {
		t.Fail()
	}
}

func Test_PlainPrimTree(t *testing.T) {
	defer guard_ut(t)
	var matrix = [][]uint{
		{0, 8, 0, 4, 0, 0, 11, 0, 0},
		{8, 0, 7, 0, 2, 0, 0, 0, 4},
		{0, 7, 0, 0, 0, 9, 0, 0, 14},
		{4, 0, 0, 0, 0, 0, 8, 0, 0},
		{0, 2, 0, 0, 0, 0, 7, 6, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 10},
		{11, 0, 0, 8, 7, 0, 0, 1, 0},
		{0, 0, 0, 0, 6, 0, 1, 0, 2},
		{0, 4, 14, 0, 0, 10, 0, 2, 0}}

	var expected = make([][]int, 9)
	expected[0] = []int{3, 1}
	expected[1] = []int{4, 8, 2}
	expected[2] = []int{5}
	expected[7] = []int{6}
	expected[8] = []int{7}

	var ret, fail = PlainPrimTree(matrix)
	if fail || !isTheSame(ret, expected) {
		t.Fail()
	}
}
