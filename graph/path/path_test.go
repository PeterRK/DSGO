package path

import (
	"DSGO/array"
	"DSGO/graph"
	"DSGO/utils"
	"math"
	"testing"
)

func Test_Dijkstra(t *testing.T) {
	defer utils.FailInPanic(t)

	roads := [][]graph.Path{
		{{1, 1}, {3, 2}},         //0
		{{4, 4}},                 //1
		{{0, 10}, {3, 5}},        //2
		{{0, 3}, {1, 9}, {4, 2}}, //3
		{{1, 6}, {2, 7}},         //4
	}

	expected := []int{19, 0, 11, 16, 4}
	ret := Dijkstra(roads, 1)
	utils.Assert(t, array.Equal(ret, expected))

	expected = []int{1, 4, 2, 3, 0}
	vec := DijkstraPath(roads, 1, 0)
	utils.Assert(t, array.Equal(vec, expected))
}

func Test_PlainDijkstra(t *testing.T) {
	defer utils.FailInPanic(t)

	matrix := [][]uint{
		{0, 1, 0, 2, 0},
		{0, 0, 0, 0, 4},
		{10, 0, 0, 5, 0},
		{3, 9, 0, 0, 2},
		{0, 6, 7, 0, 0},
	}

	expected := []int{19, 0, 11, 16, 4}
	vec := PlainDijkstra(matrix, 1)
	utils.Assert(t, array.Equal(vec, expected))

	expected = []int{1, 4, 2, 3, 0}
	vec = PlainDijkstraPath(matrix, 1, 0)
	utils.Assert(t, array.Equal(vec, expected))
}

func Test_SPFA(t *testing.T) {
	defer utils.FailInPanic(t)

	roads := [][]graph.PathS{
		{{1, 1}, {3, 2}},         //0
		{{4, 4}},                 //1
		{{0, 10}, {3, 5}},        //2
		{{0, 3}, {1, 9}, {4, 2}}, //3
		{{1, 6}, {2, 7}},         //4
	}

	expected := []int{19, 0, 11, 16, 4}
	dists, err := SPFA(roads, 1)
	utils.Assert(t, err == nil && array.Equal(dists, expected))
}

func Test_FloydWarshall(t *testing.T) {
	defer utils.FailInPanic(t)

	matrix := [][]int{
		{0, 1, math.MaxInt, 2, math.MaxInt},
		{math.MaxInt, 0, math.MaxInt, math.MaxInt, 4},
		{10, math.MaxInt, 0, 5, math.MaxInt},
		{3, 9, math.MaxInt, 0, 2},
		{math.MaxInt, 6, 7, math.MaxInt, 0}}

	expected := [][]int{
		{0, 1, 11, 2, 4},
		{19, 0, 11, 16, 4},
		{8, 9, 0, 5, 7},
		{3, 4, 9, 0, 2},
		{15, 6, 7, 12, 0},
	}

	FloydWarshall(matrix)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			utils.Assert(t, matrix[i][j] == expected[i][j])
		}
	}
}
