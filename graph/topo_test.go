package graph

import (
	"DSGO/array"
	"DSGO/utils"
	"testing"
)

func Test_TopologicalSort(t *testing.T) {
	defer utils.FailInPanic(t)

	roads := [][]int{
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

	expected := []int{2, 8, 0, 3, 7, 1, 5, 6, 4, 9, 10, 11, 12}
	vec, err := TopologicalSort(roads)
	utils.Assert(t, err == nil && array.Equal(vec, expected))
}
