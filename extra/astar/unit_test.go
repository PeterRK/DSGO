package astar

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

func Test_AStarAsDijkstra(t *testing.T) {
	defer guardUT(t)

	var roads = [][]Path{
		{{1, 1}, {3, 2}},         //0
		{{4, 4}},                 //1
		{{0, 10}, {3, 5}},        //2
		{{0, 3}, {1, 9}, {4, 2}}, //3
		{{1, 6}, {2, 7}},         //4
	}

	var expected = []int{1, 4, 2, 3, 0}
	var vec = AStar(roads, 1, 0, func(int) uint { return 0 })
	assert(t, isTheSame(vec, expected))
}

func Test_AStarNotBest(t *testing.T) {
	defer guardUT(t)

	var roads = [][]Path{
		{{1, 1}, {2, 2}}, //0
		{{3, 1}},         //1
		{{4, 2}},         //2
		{{4, 1}},         //3
		{{0, 6}},         //4
	}

	var book = []uint{100, 100, 1, 1, 0}

	var expected = []int{0, 2, 4}
	var vec = AStar(roads, 0, 4, func(id int) uint { return book[id] })
	assert(t, isTheSame(vec, expected))
}
