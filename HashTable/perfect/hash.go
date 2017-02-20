package perfect

func MurmurHash(seed uint32, str []byte) uint32 {
	var code = seed

	var m = len(str) % 4
	for i := 0; i < len(str)-m; i += 4 {
		var w = uint32(str[i]) | (uint32(str[i+1]) << 8) |
			(uint32(str[i+2]) << 16) | (uint32(str[i+3]) << 24)
		w *= 0xcc9e2d51
		w = (w << 15) | (w >> 17)
		w *= 0x1b873593
		code ^= w
		code = (code << 13) | (code >> 19)
		code += (code << 2) + 0xe6546b64
	}
	if m != 0 {
		var w = uint32(0)
		for i := len(str) - 1; i >= len(str)-m; i-- {
			w = (w << 8) | uint32(str[i])
		}
		w *= 0xcc9e2d51
		w = (w << 15) | (w >> 17)
		w *= 0x1b873593
		code ^= w
	}
	code ^= uint32(len(str))
	code ^= code >> 16
	code *= 0x85ebca6b
	code ^= code >> 13
	code *= 0xc2b2ae35
	code ^= code >> 16
	return code
}
