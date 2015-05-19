package hashtable

//成功返回true，冲突返回false
func (table *hashTable) Insert(key string) bool {
	var tail = table.bucket[table.hash(key)%uint(len(table.bucket))]
	for ; tail != nil; tail = tail.next {
		if key == tail.key {
			return false
		}
	}
	var unit = new(node)
	unit.key = key
	tail.next = unit

	table.cnt++
	if table.isCrowded() {
		if newsz, ok := nextSize(uint(len(table.bucket))); ok {
			table.resize(newsz)
		}
	}
	return true
}

func (table *hashTable) resize(size uint) {
	var old_bucket = table.bucket
	table.bucket = make([]*node, size)
	for _, unit := range old_bucket {
		for ; unit != nil; unit = unit.next {
			var index = table.hash(unit.key) % size
			unit.next, table.bucket[index] = table.bucket[index], unit
		}
	}
}

//此数列借鉴自SGI STL
var size_primes = []uint{
	97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613,
	393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 1610612741}

func initSize() uint {
	return 53
}

func nextSize(size uint) (newsz uint, ok bool) {
	var start, end = 0, len(size_primes)
	for start < end {
		var mid = (start + end) / 2
		if size < size_primes[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(size_primes) {
		return size, false
	}
	return size_primes[start], true
}
