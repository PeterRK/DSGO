package hash

import (
	"crypto/md5"
)

type BloomFliter struct {
	core [1 << 13]byte
}

//实际使用中应调整容积率和hash段数以达较优效果
//建议hash段数为ln2*(m/n)，容积率(n/m)低于1/16

var masks = [8]byte{1, 2, 4, 8, 16, 32, 64, 128}

func (tb *BloomFliter) Insert(key []byte) {
	var hash = md5.Sum(key)
	for i := 1; i < len(hash); i++ {
		var code = (uint16(hash[i]) << 8) | uint16(hash[i-1])
		tb.core[code>>3] |= masks[code&7]
	}
}

func (tb *BloomFliter) Search(key []byte) bool {
	var hash = md5.Sum(key)
	for i := 1; i < len(hash); i++ {
		var code = (uint16(hash[i]) << 8) | uint16(hash[i-1])
		if tb.core[code>>3]&masks[code&7] == byte(0) {
			return false
		}
	}
	return true
}
