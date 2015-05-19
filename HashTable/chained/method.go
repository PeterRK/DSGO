package chained

func (table *HashTable) Search(key string) bool {
	var index = table.hash(key) % uint(len(table.bucket))
	for unit := table.bucket[index]; unit != nil; unit = unit.next {
		if key == unit.key {
			return true
		}
	}
	return false
}

//成功返回true，没有返回false
func (table *HashTable) Remove(key string) bool {
	var index = table.hash(key) % uint(len(table.bucket))
	for knot := fakeHead(&table.bucket[index]); knot.next != nil; knot = knot.next {
		if key == knot.next.key {
			knot.next = knot.next.next
			table.cnt--
			return true
		}
	}
	return false
}

//成功返回true，冲突返回false
func (table *HashTable) Insert(key string) bool {
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
func (table *HashTable) resize(size uint) {
	var old_bucket = table.bucket
	table.bucket = make([]*node, size)
	for _, unit := range old_bucket {
		for ; unit != nil; unit = unit.next {
			var index = table.hash(unit.key) % size
			unit.next, table.bucket[index] = table.bucket[index], unit
		}
	}
}
