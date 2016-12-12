package path

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

func Test_Dijkstra(t *testing.T) {
	defer guardUT(t)

	var roads = [][]graph.Path{
		{{1, 1}, {3, 2}},         //0
		{{4, 4}},                 //1
		{{0, 10}, {3, 5}},        //2
		{{0, 3}, {1, 9}, {4, 2}}, //3
		{{1, 6}, {2, 7}},         //4
	}

	var expected = []int{19, 0, 11, 16, 4}
	var ret = Dijkstra(roads, 1)
	assert(t, isTheSame(ret, expected))
	ret = DijkstraX(roads, 1)
	assert(t, isTheSame(ret, expected))

	expected = []int{1, 4, 2, 3, 0}
	var vec = DijkstraPath(roads, 1, 0)
	assert(t, isTheSame(vec, expected))
	vec = DijkstraPathX(roads, 1, 0)
	assert(t, isTheSame(vec, expected))
}

func Test_PlainDijkstra(t *testing.T) {
	defer guardUT(t)

	var matrix = [][]uint{
		{0, 1, 0, 2, 0},
		{0, 0, 0, 0, 4},
		{10, 0, 0, 5, 0},
		{3, 9, 0, 0, 2},
		{0, 6, 7, 0, 0},
	}

	var expected = []int{19, 0, 11, 16, 4}
	var ret = PlainDijkstra(matrix, 1)
	if !isTheSame(ret, expected) {
		t.Fail()
	}

	expected = []int{1, 4, 2, 3, 0}
	var vec = PlainDijkstraPath(matrix, 1, 0)
	assert(t, isTheSame(vec, expected))

}

func Test_SPFA(t *testing.T) {
	defer guardUT(t)

	var roads = [][]Path{
		{{1, 1}, {3, 2}},         //0
		{{4, 4}},                 //1
		{{0, 10}, {3, 5}},        //2
		{{0, 3}, {1, 9}, {4, 2}}, //3
		{{1, 6}, {2, 7}},         //4
	}

	var expected = []int{19, 0, 11, 16, 4}
	var dists, err = SPFA(roads, 1)
	assert(t, err == nil && isTheSame(dists, expected))
}

func Test_FloydWarshall(t *testing.T) {
	defer guardUT(t)

	var matrix = [][]int{
		{0, 1, MAX_DIST, 2, MAX_DIST},
		{MAX_DIST, 0, MAX_DIST, MAX_DIST, 4},
		{10, MAX_DIST, 0, 5, MAX_DIST},
		{3, 9, MAX_DIST, 0, 2},
		{MAX_DIST, 6, 7, MAX_DIST, 0}}

	var expected = [][]int{
		{0, 1, 11, 2, 4},
		{19, 0, 11, 16, 4},
		{8, 9, 0, 5, 7},
		{3, 4, 9, 0, 2},
		{15, 6, 7, 12, 0},
	}

	FloydWarshall(matrix)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			assert(t, matrix[i][j] == expected[i][j])
		}
	}
}
