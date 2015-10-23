package hash

import (
	"testing"
)

var bench_str = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Benchmark_BKDRhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BKDRhash(bench_str)
	}
}
func Benchmark_SDBMhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SDBMhash(bench_str)
	}
}
func Benchmark_DJBhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DJBhash(bench_str)
	}
}
func Benchmark_DJB2hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DJB2hash(bench_str)
	}
}
func Benchmark_FNVhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNVhash(bench_str)
	}
}
func Benchmark_RShash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RShash(bench_str)
	}
}
func Benchmark_JShash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JShash(bench_str)
	}
}
func Benchmark_APhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		APhash(bench_str)
	}
}
