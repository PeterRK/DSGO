package skiparray

type SkipArray struct {
	space []interface{}
	index [][]uint64
}

func NewSkipArray(size uint) *SkipArray {
	var sa = new(SkipArray)
	sa.space = make([]interface{}, size)
	for size >= 128 {
		var r = size % 64
		size = (size + 63) / 64
		var level = make([]uint64, size)
		level[size-1] = ^uint64(0) >> r
		sa.index = append(sa.index, level)
	}
	return sa
}

func (sa *SkipArray) Capacity() int {
	return len(sa.space)
}

func (sa *SkipArray) alloc() int {
	var idx = 0
	for i := len(sa.index) - 1; i >= 0; i-- {
		var level = sa.index[i]
		var j = idx * 64
		for ; j < len(level); j++ {
			if level[j] != ^uint64(0) {
				idx = j
			}
		}
		if j == len(level) {
			return -1
		}
	}
	for idx *= 64; idx < len(sa.space); idx++ {
		if sa.space[idx] == nil {
			return idx //分配编号最小的空位
		}
	}
	return -1
}

func (sa *SkipArray) Insert(obj interface{}) int {
	var idx = sa.alloc()
	if idx != -1 {
		sa.space[idx] = obj
	}
	return idx
}

func (sa *SkipArray) Search(idx int) interface{} {
	if idx >= len(sa.space) {
		return nil
	}
	return sa.space[idx]
}

func (sa *SkipArray) Remove(idx int) {
	if idx >= len(sa.space) ||
		sa.space[idx] == nil {
		return
	}
	sa.space[idx] = nil
	for i := 0; i < len(sa.index); i++ {
		var level = sa.index[i]
		var r = uint(idx) % 64
		idx /= 64
		var old = level[idx]
		level[idx] &= ^(uint64(1) << (63 - r))
		if old != ^uint64(0) {
			break
		}
	}
}
