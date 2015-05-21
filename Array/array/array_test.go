package array

import (
	"testing"
)

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

func Test_Delete(t *testing.T) {
	var source, expected []int

	source = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4, 5}
	source = Delete(source, -1)
	if !isTheSame(source, expected) {
		t.Fail()
	}
	source = Delete(source, 9)
	if !isTheSame(source, expected) {
		t.Fail()
	}

	source = []int{1, 2, 3, 4, 5}
	expected = []int{2, 3, 4, 5}
	source = Delete(source, 0)
	if !isTheSame(source, expected) {
		t.Fail()
	}

	source = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4}
	source = Delete(source, 4)
	if !isTheSame(source, expected) {
		t.Fail()
	}

	source = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 4, 5}
	source = Delete(source, 2)
	if !isTheSame(source, expected) {
		t.Fail()
	}
}

func Test_Erease(t *testing.T) {
	var source, expected []int

	source = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4, 5}
	source = Erease(source, -1)
	if !isTheSame(source, expected) {
		t.Fail()
	}
	source = Erease(source, 9)
	if !isTheSame(source, expected) {
		t.Fail()
	}

	source = []int{1, 2, 3, 4, 5}
	expected = []int{5, 2, 3, 4}
	source = Erease(source, 0)
	if !isTheSame(source, expected) {
		t.Fail()
	}

	source = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4}
	source = Erease(source, 4)
	if !isTheSame(source, expected) {
		t.Fail()
	}

	source = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 5, 4}
	source = Erease(source, 2)
	if !isTheSame(source, expected) {
		t.Fail()
	}
}
