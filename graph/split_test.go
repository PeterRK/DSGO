package graph

import (
	"DSGO/array"
	"DSGO/utils"
	"testing"
)

func Test_SplitDirectedGraph(t *testing.T) {
	defer utils.FailInPanic(t)

	roads := [][]int{
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

	expected := [][]int{{0}, {4, 1}, {3}, {5, 2}, {10, 11, 9, 8, 6, 7}}
	parts := SplitDirectedGraph(roads)
	utils.Assert(t, len(parts) == len(expected))
	for i := 0; i < len(parts); i++ {
		utils.Assert(t, array.Equal(parts[i], expected[i]))
	}
}
