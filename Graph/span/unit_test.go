package span

import (
	"Graph/graph"
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

func isTheSame(vec1 []Edge, vec2 []Edge) bool {
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

	var dist, err = Prim(roads)
	assert(t, err == nil && dist == 37)
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

	var dist, err = PlainPrim(matrix)
	assert(t, err == nil && dist == 37)
}

func Test_Kruskal(t *testing.T) {
	defer guard_ut(t)

	var roads = []graph.Edge{
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

	var dist, err = Kruskal(roads, 9)
	assert(t, err == nil && dist == 37)
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

	var expected = []Edge{{0, 3}, {0, 1}, {1, 4}, {1, 8}, {8, 7}, {7, 6}, {1, 2}, {2, 5}}

	var ret, err = PrimTree(roads)
	assert(t, err == nil && isTheSame(ret, expected))
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

	var expected = []Edge{{0, 3}, {0, 1}, {1, 4}, {1, 8}, {8, 7}, {7, 6}, {1, 2}, {2, 5}}

	var ret, err = PlainPrimTree(matrix)
	assert(t, err == nil && isTheSame(ret, expected))
}
