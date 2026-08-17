package sort

import (
	"DSGO/utils"
	"DSGO/utils/constraints"
	"math"
	"slices"
	"testing"
)

func testRadixSortInteger[T constraints.Integer](t *testing.T) {
	t.Helper()
	var boundaries []T
	for _, value := range []int64{
		math.MinInt64, math.MaxInt64, math.MinInt32, math.MaxInt32,
		math.MaxUint32, math.MinInt16, math.MaxInt16, math.MaxUint16,
		math.MinInt8, math.MaxInt8, math.MaxUint8, -1, 0, 1,
	} {
		boundaries = append(boundaries, T(value))
	}
	// For int8 and uint8 this includes every possible value.
	allBytes := make([]T, 256)
	for i := range allBytes {
		allBytes[i] = T(255 - i)
	}
	for _, input := range [][]T{
		nil, {}, {1}, {2, 1}, {2, 1, 2, 0}, boundaries, allBytes,
		utils.PseudoRandomArray[T](1024, 42),
	} {
		got, want := slices.Clone(input), slices.Clone(input)
		slices.Sort(want)
		RadixSort(got)
		if !slices.Equal(got, want) {
			t.Fatalf("RadixSort(%v) = %v; want %v", input, got, want)
		}
	}
}

func Test_RadixSortIntegerTypes(t *testing.T) {
	t.Run("int8", testRadixSortInteger[int8])
	t.Run("uint8", testRadixSortInteger[uint8])
	t.Run("int16", testRadixSortInteger[int16])
	t.Run("uint16", testRadixSortInteger[uint16])
	t.Run("int32", testRadixSortInteger[int32])
	t.Run("uint32", testRadixSortInteger[uint32])
	t.Run("int64", testRadixSortInteger[int64])
	t.Run("uint64", testRadixSortInteger[uint64])
	t.Run("int", testRadixSortInteger[int])
	t.Run("uint", testRadixSortInteger[uint])
	t.Run("uintptr", testRadixSortInteger[uintptr])
	type signedByte int8
	type unsignedByte uint8
	t.Run("named_int8", testRadixSortInteger[signedByte])
	t.Run("named_uint8", testRadixSortInteger[unsignedByte])
}
