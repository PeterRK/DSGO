package cuckoo

import (
	"bytes"
)

func (tb *hashTable) Search(key []byte) bool {
	for i := 0; i < WAYS; i++ {
		var idx = (tb.idx + i) % WAYS
		var table = &tb.core[idx]
		var code = table.hash(key)
		//var index = code % uint(len(table.bucket))
		var index = code & (uint(len(table.bucket)) - 1)
		var target = table.bucket[index]
		if target != nil &&
			target.code[idx] == code &&
			bytes.Compare(target.key, key) == 0 {
			return true
		}
	}
	return false
}

//成功返回true，没有返回false
func (tb *hashTable) Remove(key []byte) bool {
	for i := 0; i < WAYS; i++ {
		var idx = (tb.idx + i) % WAYS
		var table = &tb.core[idx]
		var code = table.hash(key)
		//var index = code % uint(len(table.bucket))
		var index = code & (uint(len(table.bucket)) - 1)
		var target = table.bucket[index]
		if target != nil &&
			target.code[idx] == code &&
			bytes.Compare(target.key, key) == 0 {
			tb.cnt--
			table.bucket[index] = nil
			return true
		}
	}
	return false
}

//成功返回true，冲突返回false
func (tb *hashTable) Insert(key []byte) bool {
	var code [WAYS]uint
	for i := 0; i < WAYS; i++ {
		var table = &tb.core[i]
		code[i] = table.hash(key)
		//var index = code[i] % uint(len(table.bucket))
		var index = code[i] & (uint(len(table.bucket)) - 1)
		var target = table.bucket[index]
		if target != nil &&
			target.code[i] == code[i] &&
			bytes.Compare(target.key, key) == 0 {
			return false
		}
	}
	tb.cnt++
	var unit = new(node)
	unit.key, unit.code = key, code

	for obj, age := unit, 0; ; age++ {
		for idx, trys := tb.idx, 0; trys < WAYS; idx = (idx + 1) % WAYS {
			var table = &tb.core[idx]
			//var index = obj.code[idx] % uint(len(table.bucket))
			var index = obj.code[idx] & (uint(len(table.bucket)) - 1)
			if table.bucket[index] == nil {
				table.bucket[index] = obj
				return true
			}
			obj, table.bucket[index] = table.bucket[index], obj
			if obj == unit {
				trys++ //回绕计数
			}
		}

		if age != 0 { //这里设定一个阈值，限制扩容次数
			panic("too many conflicts")
		} //实际上不能解决大量hash重码的情况，最坏情况只能报错

		//调整失败(回绕)，扩容
		tb.idx = (tb.idx + (WAYS - 1)) % WAYS
		var table = &tb.core[tb.idx]
		var old_bucket = table.bucket
		table.bucket = make([]*node, len(old_bucket)<<WAYS)
		for _, u := range old_bucket {
			if u != nil {
				//var index = u.code[tb.idx] % uint(len(table.bucket))
				var index = u.code[tb.idx] & (uint(len(table.bucket)) - 1)
				table.bucket[index] = u //倍扩，绝对不会冲突
			}
		}
	}
	return false
}
