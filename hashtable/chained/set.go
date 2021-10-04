package chained

import (
	"DSGO/array"
	"DSGO/hashtable"
	"DSGO/linkedlist"
	"DSGO/utils"
)

type node = linkedlist.Node[string]

type hashSet struct {
	bucket    []*node
	size      int
	nextLine  int     //标记待处理的旧表行
	oldBucket []*node //旧表（仅在rehash过程中有内容）
}

func (s *hashSet) Size() int {
	return s.size
}

func (s *hashSet) IsEmpty() bool {
	return s.size == 0
}

func (s *hashSet) isCrowded() bool {
	return s.size*2 > len(s.bucket)*3
}

func (s *hashSet) isWasteful() bool {
	return s.size*10 < len(s.bucket)
}

func (s *hashSet) isMoving() bool {
	return len(s.oldBucket) != 0
}

func (s *hashSet) Clear() {
	s.bucket = make([]*node, primes[0])
	s.size = 0
	s.nextLine = 0
	s.oldBucket = nil
}

func NewSet() utils.StrSet {
	s := new(hashSet)
	s.Clear()
	return s
}

func hash(key string) uint32 {
	return hashtable.Hash32(0, key)
}

func (s *hashSet) moveLine() {
	size := uint32(len(s.bucket))
	for head := s.oldBucket[s.nextLine]; head != nil; {
		unit, index := head, hash(head.Val)%size
		head = head.Next
		unit.Next, s.bucket[index] = s.bucket[index], unit
	}
	s.oldBucket[s.nextLine] = nil
	s.nextLine++
	if s.nextLine == len(s.oldBucket) { //rehash完成
		s.nextLine = 0
		s.oldBucket = nil //GC
	}
}

//此数列借鉴自SGI C++STL
var primes = []uint32{
	17, 29, 53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613,
	393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 1610612741}

func (s *hashSet) resize(size uint32) {
	s.oldBucket, s.bucket = s.bucket, make([]*node, size)
}

func (s *hashSet) Search(key string) bool {
	code := hash(key)
	index := code % uint32(len(s.bucket))
	found := search(s.bucket[index], key)
	if s.isMoving() {
		if !found { //尝试从旧表中查找
			index = code % uint32(len(s.oldBucket))
			found = search(s.oldBucket[index], key)
		}
		s.moveLine() //推进rehash过程
	}
	return found
}

func search(head *node, key string) bool {
	for ; head != nil; head = head.Next {
		if key == head.Val {
			return true
		}
	}
	return false
}

func (s *hashSet) Remove(key string) bool {
	code, done := hash(key), false
	index := code % uint32(len(s.bucket))
	s.bucket[index], done = remove(s.bucket[index], key)
	if s.isMoving() {
		if !done { //尝试从旧表中删除
			index = code % uint32(len(s.oldBucket))
			s.oldBucket[index], done = remove(s.oldBucket[index], key)
		}
		s.moveLine()
	}
	if done {
		s.size--
		if !s.isMoving() && s.isWasteful() {
			idx := array.SearchFirstGE(primes, uint32(len(s.bucket))) - 1
			if idx >= 0 {
				s.resize(primes[idx])
			}
		}
	}
	return done
}

func remove(head *node, key string) (*node, bool) {
	for knot := linkedlist.FakeHead(&head); knot.Next != nil; knot = knot.Next {
		if key == knot.Next.Val {
			knot.Next = knot.Next.Next
			return head, true
		}
	}
	return head, false
}

func (s *hashSet) Insert(key string) bool {
	code := hash(key)
	index := code % uint32(len(s.bucket))
	conflict := search(s.bucket[index], key)
	if s.isMoving() {
		if !conflict {
			index := code % uint32(len(s.oldBucket))
			conflict = search(s.oldBucket[index], key)
		}
		s.moveLine()
	}
	if !conflict {
		unit := new(node)
		unit.Val = key
		unit.Next, s.bucket[index] = s.bucket[index], unit
		s.size++
		if !s.isMoving() && s.isCrowded() {
			idx := array.SearchSuccessor(primes, uint32(len(s.bucket)))
			if idx < len(primes) {
				s.resize(primes[idx])
			}
		}
		return true
	}
	return false
}
