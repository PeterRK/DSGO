package hash

import (
	"testing"
)

func Benchmark_JenkinsHash(b *testing.B) {
	hashBenchmark(b, bench_data, JenkinsHash)
}
func Benchmark_MurmurHash(b *testing.B) {
	hashBenchmark(b, bench_data, MurmurHash)
}
func Benchmark_BKDRhash(b *testing.B) {
	hashBenchmark(b, bench_data, BKDRhash)
}
func Benchmark_SDBMhash(b *testing.B) {
	hashBenchmark(b, bench_data, SDBMhash)
}
func Benchmark_DJBhash(b *testing.B) {
	hashBenchmark(b, bench_data, DJBhash)
}
func Benchmark_DJB2hash(b *testing.B) {
	hashBenchmark(b, bench_data, DJB2hash)
}
func Benchmark_FNVhash(b *testing.B) {
	hashBenchmark(b, bench_data, FNVhash)
}
func Benchmark_RShash(b *testing.B) {
	hashBenchmark(b, bench_data, RShash)
}
func Benchmark_JShash(b *testing.B) {
	hashBenchmark(b, bench_data, JShash)
}
func Benchmark_APhash(b *testing.B) {
	hashBenchmark(b, bench_data, APhash)
}

var bench_data = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func hashBenchmark(b *testing.B, str []byte, hash func(str []byte) uint32) {
	b.SetBytes(int64(len(str)))
	for i := 0; i < b.N; i++ {
		hash(str)
	}
}
