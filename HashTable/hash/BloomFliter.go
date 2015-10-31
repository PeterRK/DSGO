package hash

import (
	"crypto/md5"
)

const VEC_SZIE = 1 << 13

func hashToIndex(high byte, low byte) uint {
	return (uint(high) << 8) | uint(low)
}

type BloomFliter struct {
	core [VEC_SZIE]byte
}

func (tb *BloomFliter) setBit(index uint) {
	tb.core[index>>3] |= (byte(1) << (index & 7))
}
func (tb *BloomFliter) testBit(index uint) bool {
	var bit = tb.core[index>>3] & (byte(1) << (index & 7))
	return bit != byte(0)
}

//实际使用中应调整容积率和hash段数以达较优效果
//建议hash段数为ln2*(m/n)，容积率(n/m)低于1/16

func (tb *BloomFliter) Insert(key []byte) {
	var hash = md5.Sum(key)
	for i := 1; i < len(hash); i++ {
		var index = hashToIndex(hash[i], hash[i-1])
		tb.setBit(index)
	}
}

func (tb *BloomFliter) Search(key []byte) bool {
	var hash = md5.Sum(key)
	for i := 1; i < len(hash); i++ {
		var index = hashToIndex(hash[i], hash[i-1])
		if !tb.testBit(index) {
			return false
		}
	}
	return true
}
