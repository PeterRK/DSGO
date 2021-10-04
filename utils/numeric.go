package utils

import (
	"constraints"
)

func Log2Ceil(num uint) uint {
	ceil := uint(0)
	for ; num != 0; ceil++ {
		num /= 2
	}
	return ceil
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}