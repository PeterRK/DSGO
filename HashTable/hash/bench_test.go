package hash

import (
	"testing"
)

func Test_Nothing(t *testing.T) {
	//Do Nothing
}

var str = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Benchmark_BKDRhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BKDRhash(str)
	}
}
func Benchmark_SDBMhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SDBMhash(str)
	}
}
func Benchmark_DJBhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DJBhash(str)
	}
}
func Benchmark_DJB2hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DJB2hash(str)
	}
}
func Benchmark_FNVhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNVhash(str)
	}
}
func Benchmark_RShash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RShash(str)
	}
}
func Benchmark_JShash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JShash(str)
	}
}
func Benchmark_APhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		APhash(str)
	}
}
