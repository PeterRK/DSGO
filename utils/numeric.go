package utils

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"time"
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

func PseudoRandomArray[T constraints.Integer](size int, seed int64) []T {
	rand.Seed(seed)
	list := make([]T, size)
	for i := 0; i < size; i++ {
		list[i] = T(rand.Uint64())
	}
	return list
}

func RandomArray[T constraints.Integer](size int) []T {
	return PseudoRandomArray[T](size, time.Now().UnixNano())
}
