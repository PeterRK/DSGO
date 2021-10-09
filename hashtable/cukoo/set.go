package cukoo

import (
	"DSGO/hashtable"
	"DSGO/utils"
)

type node struct {
	code [4]uint32
	key  string
}

type hashSet struct {
	buckets [4][]*node
	master  int //当前主表号
	size    int
	full    int
}

func (s *hashSet) Size() int {
	return s.size
}

func (s *hashSet) IsEmpty() bool {
	return s.size == 0
}

func (s *hashSet) Clear() {
	s.master, s.size = 0, 0
	size := 2 //2^n
	for i := 3; i >= 0; i-- {
		size *= 2
		s.buckets[i] = make([]*node, size)
	}
	s.full = size * 3 / 2 //(size * 15 / 8) * 4 / 5
}

func NewSet() utils.StrSet {
	s := new(hashSet)
	s.Clear()
	return s
}

func mod(code uint32, bucketSize int) uint32 {
	//return code % uint(bucketSize)
	return code & (uint32(bucketSize) - 1) //bucketSize == 2^n
}

func hash(key string) [4]uint32 {
	a, b := hashtable.Hash128(0, key)
	return [4]uint32{uint32(a), uint32(a >> 32), uint32(b), uint32(b >> 32)}
}

func (s *hashSet) find(key string, erase bool) bool {
	idx := s.master
	code := hash(key)
	for i := 0; i < 4; i++ {
		idx = (s.master + i) % 4
		bucket := s.buckets[idx]
		pos := mod(code[idx], len(bucket))
		target := bucket[pos]
		if target != nil && target.code == code && target.key == key {
			if erase {
				bucket[pos] = nil
			}
			return true
		}
	}
	return false
}

func (s *hashSet) Search(key string) bool {
	return s.find(key, false)
}

func (s *hashSet) Remove(key string) bool {
	if s.find(key, true) {
		s.size--
		return true
	}
	return false
}

func (s *hashSet) Insert(key string) bool {
	if s.find(key, false) {
		return false
	}
	s.size++

	unit := new(node)
	unit.key = key
	unit.code = hash(key)

	for obj, age := unit, 0; ; age++ {
		for idx, trys := s.master, 0; trys < 4; idx = (idx + 1) % 4 {
			bucket := s.buckets[idx]
			pos := mod(obj.code[idx], len(bucket))
			if bucket[pos] == nil {
				bucket[pos] = obj
				return true
			}
			obj, bucket[pos] = bucket[pos], obj
			if obj == unit {
				if s.size > s.full {
					break
				}
				trys++ //回绕计数
			}
		}

		if age > 0 { //这里设定一个阈值，限制扩容次数
			panic("too many conflicts")
		} //实际上不能解决大量hash重码的情况，最坏情况只能报错

		s.expand() //调整失败(回绕)，扩容
	}
	return false
}

func (s *hashSet) expand() {
	s.master = (s.master + 3) % 4
	oldBucket := s.buckets[s.master]
	bucket := make([]*node, len(oldBucket)<<4)
	for _, unit := range oldBucket {
		if unit != nil {
			pos := mod(unit.code[s.master], len(bucket))
			bucket[pos] = unit //倍扩，绝对不会冲突
		}
	}
	s.buckets[s.master] = bucket
	s.full = len(bucket) * 3 / 2
}
