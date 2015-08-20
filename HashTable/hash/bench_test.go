package hash

import (
	"testing"
)

func Test_Nothing(t *testing.T) {
	//Do Nothing
}

var str = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Benchmark_BKDRhash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		BKDRhash(data)
	}
}
func Benchmark_SDBMhash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		SDBMhash(data)
	}
}
func Benchmark_DJBhash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		DJBhash(data)
	}
}
func Benchmark_DJB2hash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		DJB2hash(data)
	}
}
func Benchmark_FNVhash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		FNVhash(data)
	}
}
func Benchmark_RShash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		RShash(data)
	}
}
func Benchmark_JShash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		JShash(data)
	}
}
func Benchmark_APhash(b *testing.B) {
	var data = []byte(str)
	for i := 0; i < b.N; i++ {
		APhash(data)
	}
}
