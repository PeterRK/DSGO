package hashtable

func rot(x uint64, k int) uint64 {
	return (x << k) | (x >> (64 - k))
}

type state struct {
	a, b, c, d uint64
}

func (s *state) mix() {
	s.c = rot(s.c, 50)
	s.c += s.d
	s.a ^= s.c
	s.d = rot(s.d, 52)
	s.d += s.a
	s.b ^= s.d
	s.a = rot(s.a, 30)
	s.a += s.b
	s.c ^= s.a
	s.b = rot(s.b, 41)
	s.b += s.c
	s.d ^= s.b
	s.c = rot(s.c, 54)
	s.c += s.d
	s.a ^= s.c
	s.d = rot(s.d, 48)
	s.d += s.a
	s.b ^= s.d
	s.a = rot(s.a, 38)
	s.a += s.b
	s.c ^= s.a
	s.b = rot(s.b, 37)
	s.b += s.c
	s.d ^= s.b
	s.c = rot(s.c, 62)
	s.c += s.d
	s.a ^= s.c
	s.d = rot(s.d, 34)
	s.d += s.a
	s.b ^= s.d
	s.a = rot(s.a, 5)
	s.a += s.b
	s.c ^= s.a
	s.b = rot(s.b, 36)
	s.b += s.c
	s.d ^= s.b
}

func (s *state) end() {
	s.d ^= s.c
	s.c = rot(s.c, 15)
	s.d += s.c
	s.a ^= s.d
	s.d = rot(s.d, 52)
	s.a += s.d
	s.b ^= s.a
	s.a = rot(s.a, 26)
	s.b += s.a
	s.c ^= s.b
	s.b = rot(s.b, 51)
	s.c += s.b
	s.d ^= s.c
	s.c = rot(s.c, 28)
	s.d += s.c
	s.a ^= s.d
	s.d = rot(s.d, 9)
	s.a += s.d
	s.b ^= s.a
	s.a = rot(s.a, 47)
	s.b += s.a
	s.c ^= s.b
	s.b = rot(s.b, 54)
	s.c += s.b
	s.d ^= s.c
	s.c = rot(s.c, 32)
	s.d += s.c
	s.a ^= s.d
	s.d = rot(s.d, 25)
	s.a += s.d
	s.b ^= s.a
	s.a = rot(s.a, 63)
	s.b += s.a
}

func getU64(str string) uint64 {
	return uint64(str[0]) | (uint64(str[1]) << 8) |
		(uint64(str[2]) << 16) | (uint64(str[3]) << 24) |
		(uint64(str[4]) << 32) | (uint64(str[5]) << 40) |
		(uint64(str[6]) << 48) | (uint64(str[7]) << 56)
}

func getU32(str string) uint32 {
	return uint32(str[0]) | (uint32(str[1]) << 8) |
		(uint32(str[2]) << 16) | (uint32(str[3]) << 24)
}

func getU16(str string) uint16 {
	return uint16(str[0]) | (uint16(str[1]) << 8)
}

func Hash160(seed uint64, str string) (uint64, uint64, uint32) {
	const magic uint64 = 0xdeadbeefdeadbeef
	s := state{seed, seed, magic, magic}

	for ; len(str) >= 32; str = str[32:] {
		s.c += getU64(str)
		s.d += getU64(str[8:])
		s.mix()
		s.a += getU64(str[16:])
		s.b += getU64(str[24:])
	}
	if len(str) >= 16 {
		s.c += getU64(str)
		s.d += getU64(str[8:])
		s.mix()
		str = str[16:]
	}

	s.d += uint64(len(str)) << 56
	switch len(str) {
	case 15:
		s.d += (uint64(str[14]) << 48) |
			(uint64(getU16(str[12:])) << 32) |
			uint64(getU32(str[8:]))
		s.c += getU64(str)
	case 14:
		s.d += (uint64(getU16(str[12:])) << 32) |
			uint64(getU32(str[8:]))
		s.c += getU64(str)
	case 13:
		s.d += (uint64(str[12]) << 32) | uint64(getU32(str[8:]))
		s.c += getU64(str)
	case 12:
		s.d += uint64(getU32(str[8:]))
		s.c += getU64(str)
	case 11:
		s.d += (uint64(str[10]) << 16) | uint64(getU16(str[8:]))
		s.c += getU64(str)
	case 10:
		s.d += uint64(getU16(str[8:]))
		s.c += getU64(str)
	case 9:
		s.d += uint64(str[8])
		s.c += getU64(str)
	case 8:
		s.c += getU64(str)
	case 7:
		s.c += (uint64(str[6]) << 48) |
			(uint64(getU16(str[4:])) << 32) |
			uint64(getU32(str))
	case 6:
		s.c += (uint64(getU16(str[4:])) << 32) |
			uint64(getU32(str))
	case 5:
		s.c += (uint64(str[4]) << 32) | uint64(getU32(str))
	case 4:
		s.c += uint64(getU32(str))
	case 3:
		s.c += (uint64(str[2]) << 16) | uint64(getU16(str))
	case 2:
		s.c += uint64(getU16(str))
	case 1:
		s.c += uint64(str[0])
	case 0:
		s.c += magic
		s.d += magic

	}
	s.end()
	return s.a, s.b, uint32(s.c)
}

func Hash128(seed uint64, str string) (uint64, uint64) {
	a, b, _ := Hash160(seed, str)
	return a, b
}

func Hash64(seed uint64, str string) uint64 {
	c, _, _ := Hash160(seed, str)
	return c
}
