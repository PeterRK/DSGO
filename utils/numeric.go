package utils

import (
	"DSGO/utils/constraints"
	"math/rand/v2"
	"time"
)

func PseudoRandomArray[T constraints.Integer](size int, seed int64) []T {
	base := uint64(seed)
	rng := rand.New(rand.NewPCG(base, ^base))
	list := make([]T, size)
	for i := 0; i < size; i++ {
		list[i] = T(rng.Uint64())
	}
	return list
}

func RandomArray[T constraints.Integer](size int) []T {
	return PseudoRandomArray[T](size, time.Now().UnixNano())
}
