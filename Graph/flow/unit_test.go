package flow

import (
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

//seperate & search ?

func Test_Dinic(t *testing.T) {
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
	assert(t, Dinic(matrix) == 12)

	matrix = [][]uint{
		{0, 16, 13, 0, 0, 0},
		{0, 0, 0, 12, 0, 0},
		{0, 4, 0, 0, 14, 0},
		{0, 0, 0, 0, 0, 20},
		{0, 0, 0, 7, 0, 4},
		{0, 0, 0, 0, 0, 0}}
	assert(t, Dinic(matrix) == 23)

	matrix = [][]uint{
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0}}
	assert(t, Dinic(matrix) == 2)
}
