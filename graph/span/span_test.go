package span

import (
	"DSGO/array"
	"DSGO/graph"
	"DSGO/utils"
	"testing"
)

func Test_Prim(t *testing.T) {
	defer utils.FailInPanic(t)

	roads := [][]graph.Path{
		{{1, 8}, {3, 4}, {6, 11}},          //0
		{{0, 8}, {2, 7}, {4, 2}, {8, 4}},   //1
		{{1, 7}, {5, 9}, {8, 14}},          //2
		{{0, 4}, {6, 8}},                   //3
		{{1, 2}, {6, 7}, {7, 6}},           //4
		{{2, 9}, {8, 10}},                  //5
		{{0, 11}, {3, 8}, {4, 7}, {7, 1}},  //6
		{{4, 6}, {6, 1}, {8, 2}},           //7
		{{1, 4}, {2, 14}, {5, 10}, {7, 2}}, //8
	}

	dist, err := Prim(roads)
	utils.Assert(t, err == nil && dist == 37)

	expected := []graph.SimpleEdge{
		{0, 3}, {0, 1}, {1, 4}, {1, 8}, {8, 7}, {7, 6}, {1, 2}, {2, 5}}
	ret, err := PrimTree(roads)
	utils.Assert(t, err == nil && array.Equal(ret, expected))
}

func Test_PlainPrim(t *testing.T) {
	defer utils.FailInPanic(t)

	matrix := [][]uint{
		{0, 8, 0, 4, 0, 0, 11, 0, 0},
		{8, 0, 7, 0, 2, 0, 0, 0, 4},
		{0, 7, 0, 0, 0, 9, 0, 0, 14},
		{4, 0, 0, 0, 0, 0, 8, 0, 0},
		{0, 2, 0, 0, 0, 0, 7, 6, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 10},
		{11, 0, 0, 8, 7, 0, 0, 1, 0},
		{0, 0, 0, 0, 6, 0, 1, 0, 2},
		{0, 4, 14, 0, 0, 10, 0, 2, 0},
	}

	dist, err := PlainPrim(matrix)
	utils.Assert(t, err == nil && dist == 37)

	expected := []graph.SimpleEdge{
		{0, 3}, {0, 1}, {1, 4}, {1, 8}, {8, 7}, {7, 6}, {1, 2}, {2, 5}}
	ret, err := PlainPrimTree(matrix)
	utils.Assert(t, err == nil && array.Equal(ret, expected))
}

func Test_Kruskal(t *testing.T) {
	defer utils.FailInPanic(t)

	roads := []graph.Edge{
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
		{7, 8, 2},
	}

	dist, err := Kruskal(roads, 9)
	utils.Assert(t, err == nil && dist == 37)

	dist, err = Kruskal_v2(roads, 9)
	utils.Assert(t, err == nil && dist == 37)
}
