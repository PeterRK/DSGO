package skiplist

type Random interface {
	Next() uint32
}

type easyRand struct {
	magic uint
}

func NewEasyRand(seed uint) Random {
	var rand = new(easyRand)
	rand.magic = seed
	return rand
}
func (r *easyRand) Next() uint32 {
	var num = uint32(r.magic)
	r.magic = r.magic*1103515245 + 12345
	return num
}

//MT19937的实现参考维基百科
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
func NewMT19937(seed uint) Random {
	var mt = new(mt19937)
	mt.initialize(uint32(seed))
	return mt
}

func (mt *mt19937) Next() uint32 {
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

//Xorshift参考维基百科
type xorshift struct {
	x, y, z, w uint32
}

func (xs *xorshift) initialize(seed uint32) {
	xs.x, xs.y, xs.z = 0x6c078965, 0x9908b0df, 0x9d2c5680
	xs.w = seed
}

func NewXorshift(seed uint) Random {
	var xs = new(xorshift)
	xs.initialize(uint32(seed))
	return xs
}

func (xs *xorshift) Next() uint32 {
	var t = xs.x ^ (xs.x << 11)
	xs.x, xs.y, xs.z = xs.y, xs.z, xs.w
	xs.w ^= (xs.w >> 19) ^ t ^ (t >> 8)
	return xs.w
}
