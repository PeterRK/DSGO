package skiplist

//本实现参考维基百科

type mt19937 struct {
	index int
	array [624]uint32
}

func (mt *mt19937) initialize(seed uint32) {
	mt.array[0] = seed
	for i := 1; i < 624; i++ {
		mt.array[i] = uint32(0x6c078965)*(mt.array[i-1]^(mt.array[i-1]>>30)) + uint32(i)
	}
}

func (mt *mt19937) next() uint32 {
	if mt.index == 0 {
		for i := 0; i < 624; i++ {
			var num = (mt.array[i] & uint32(0x80000000)) + (mt.array[(i+1)%624] & uint32(0x7fffffff))
			mt.array[i] = mt.array[(i+397)%624] ^ (num >> 1)
			if num%2 != 0 {
				mt.array[i] ^= uint32(0x9908b0df)
			}
		}
	}

	var num = mt.array[mt.index]
	mt.index = (mt.index + 1) % 624
	num ^= num >> 11
	num ^= (num << 7) & uint32(0x9d2c5680)
	num ^= (num << 15) & uint32(0xefc60000)
	num ^= num >> 18
	return num
}
