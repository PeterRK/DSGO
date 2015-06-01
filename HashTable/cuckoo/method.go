package cuckoo

func (tb *hashTable) Search(key string) bool {
	return tb.a.search(key) || tb.b.search(key)
}
func (tb *table) search(key string) bool {
	var code = tb.hash(key)
	var unit = tb.bucket[code&tb.mask]
	return unit != nil && unit.code[tb.id] == code && unit.key == key
}

//成功返回true，没有返回false
func (tb *hashTable) Remove(key string) bool {
	if tb.a.search(key) || tb.b.search(key) {
		tb.cnt--
		return true
	}
	return false
}
func (tb *table) remove(key string) bool {
	var code = tb.hash(key)
	var index = code & tb.mask
	var unit = tb.bucket[index]
	if unit != nil && unit.code[tb.id] == code && unit.key == key {
		tb.bucket[index] = nil
		return true
	}
	return false
}

//成功返回true，冲突返回false
func (tb *hashTable) Insert(key string) bool {
	if tb.Search(key) {
		return false
	}
	var unit = new(node)
	unit.key = key
	unit.code[tb.a.id], unit.code[tb.b.id] = tb.a.hash(key), tb.b.hash(key)
	tb.cnt++

	for obj, age := unit, 0; ; age++ {
		for { //震荡调整
			var index = obj.code[tb.a.id] & tb.a.mask
			if tb.a.bucket[index] == nil {
				tb.a.bucket[index] = obj
				return true
			}
			obj, tb.a.bucket[index] = tb.a.bucket[index], obj
			if obj == unit {
				break
			}
			index = obj.code[tb.b.id] & tb.b.mask
			if tb.b.bucket[index] == nil {
				tb.b.bucket[index] = obj
				return true
			}
			obj, tb.b.bucket[index] = tb.b.bucket[index], obj
		}

		if age == 2 {
			panic("hash fail!")
		} //实际上不能解决大量hash重码的情况

		//调整失败(回绕)，扩容
		tb.a, tb.b = tb.b, tb.a
		var old_bucket = tb.a.bucket
		tb.a.bucket = make([]*node, len(old_bucket)<<2)
		tb.a.mask = (tb.a.mask << 2) | 0x3
		for _, pt := range old_bucket {
			if pt != nil {
				var index = pt.code[tb.a.id] & tb.a.mask
				tb.a.bucket[index] = pt //倍扩，绝对不会冲突
			}
		}
	}
	return false
}
