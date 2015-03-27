package cuckoo

func (table *HashTable) Search(key string) bool {
	return table.first.search(key) || table.second.search(key)
}
func (table *coreTable) search(key string) bool {
	var code = table.hash(key)
	var pt = table.bucket[code%uint(len(table.bucket))]
	return pt != nil && pt.code[table.id] == code && pt.key == key
}

//成功返回true，没有返回false
func (table *HashTable) Remove(key string) bool {
	if table.first.search(key) || table.second.search(key) {
		table.cnt--
		return true
	}
	return false
}
func (table *coreTable) remove(key string) bool {
	var code = table.hash(key)
	var index = code & table.mask
	var pt = table.bucket[index]
	if pt != nil && pt.code[table.id] == code && pt.key == key {
		table.bucket[index] = nil
		return true
	}
	return false
}

//成功返回true，冲突返回false
func (table *HashTable) Insert(key string) bool {
	if table.Search(key) {
		return false
	}
	var obj = new(node)
	obj.key = key
	obj.code[table.first.id], obj.code[table.second.id] = table.first.hash(key), table.second.hash(key)

	for age := 0; ; age++ {
		var pt = obj
		for { //震荡调整
			var index = obj.code[table.first.id] & table.first.mask
			if table.first.bucket[index] == nil {
				table.first.bucket[index] = pt
				return true
			}
			pt, table.first.bucket[index] = table.first.bucket[index], pt
			if pt == obj {
				break
			}
			index = obj.code[table.second.id] & table.second.mask
			if table.second.bucket[index] == nil {
				table.second.bucket[index] = pt
				return true
			}
			pt, table.second.bucket[index] = table.second.bucket[index], pt
		}

		if age == 2 {
			panic("hash fail!")
		} //实际上不能解决大量hash重码的情况

		//调整失败(回绕)，扩容
		table.first, table.second = table.second, table.first
		var old_bucket = table.first.bucket
		table.first.bucket = make([]*node, len(old_bucket)<<2)
		table.first.mask = (table.first.mask << 2) | 0x3
		for _, unit := range old_bucket {
			var index = unit.code[table.first.id] & table.first.mask
			table.first.bucket[index] = unit //倍扩，绝对不会冲突
		}
	}
	return false
}
