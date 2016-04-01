package flow

import (
	"DSGO/Graph/graph"
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

//seperate & search ?

func Test_DinicM(t *testing.T) {
	defer guardUT(t)

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
	assert(t, DinicM(matrix, 8, 0) == 12)

	matrix = [][]uint{
		{0, 16, 13, 0, 0, 0},
		{0, 0, 0, 12, 0, 0},
		{0, 4, 0, 0, 14, 0},
		{0, 0, 0, 0, 0, 20},
		{0, 0, 0, 7, 0, 4},
		{0, 0, 0, 0, 0, 0}}
	assert(t, DinicM(matrix, 0, 5) == 23)

	matrix = [][]uint{
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0}}
	assert(t, DinicM(matrix, 0, 5) == 2)
}

func Test_Dinic(t *testing.T) {
	defer guardUT(t)

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
	assert(t, Dinic(roads, 8, 0) == 12)

	roads = make([][]graph.Path, 6)
	roads[0] = []graph.Path{{1, 16}, {2, 13}}
	roads[1] = []graph.Path{{3, 12}}
	roads[2] = []graph.Path{{1, 4}, {4, 14}}
	roads[3] = []graph.Path{{5, 20}}
	roads[4] = []graph.Path{{3, 7}, {5, 4}}
	assert(t, Dinic(roads, 0, 5) == 23)

	roads = make([][]graph.Path, 6)
	roads[0] = []graph.Path{{1, 1}, {2, 1}}
	roads[1] = []graph.Path{{4, 1}}
	roads[2] = []graph.Path{{3, 1}, {4, 1}}
	roads[3] = []graph.Path{{5, 1}}
	roads[4] = []graph.Path{{5, 1}}
	assert(t, Dinic(roads, 0, 5) == 2)
}
