package hash

type HashTable interface {
	Size() int
	IsEmpty() bool
	Insert(key []byte) bool
	Search(key []byte) bool
	Remove(key []byte) bool
}

//以下算法搜集自网络，基本思路类似，效果相近

func JenkinsHash(str []byte) uint32 {
	var code = uint32(0)
	for i := 0; i < len(str); i++ {
		code += uint32(str[i])
		code += code << 10
		code ^= code >> 6
	}
	code += code << 3
	code ^= code >> 11
	code += code << 15
	return code
}

func MurmurHash(str []byte) uint32 {
	var code = uint32(0)

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

//by Brian Kernighan & Dennis Ritchie
func BKDRhash(str []byte) uint32 {
	const factor uint32 = 131 //31、131、1313、 13131、131313
	var code = uint32(0)
	for i := 0; i < len(str); i++ {
		code = code*factor + uint32(str[i])
	}
	return code
}

//from SDBM
func SDBMhash(str []byte) uint32 {
	var code = uint32(0)
	for i := 0; i < len(str); i++ {
		code = code*65599 + uint32(str[i])
	}
	return code
}

//by Daniel J. Bernstein
func DJBhash(str []byte) uint32 {
	var code = uint32(5381)
	for i := 0; i < len(str); i++ {
		code = code*33 + uint32(str[i])
	}
	return code
}
func DJB2hash(str []byte) uint32 {
	var code = uint32(5381)
	for i := 0; i < len(str); i++ {
		code = code*33 ^ uint32(str[i])
	}
	return code
}

//from Unix
func FNVhash(str []byte) uint32 {
	var code = uint32(2166136261)
	for i := 0; i < len(str); i++ {
		code = code*16777619 ^ uint32(str[i])
	}
	return code
}

//Robert Sedgwicks
func RShash(str []byte) uint32 {
	var magic = uint32(63689)
	var code uint32 = 0
	for i := 0; i < len(str); i++ {
		code = code*magic + uint32(str[i])
		magic *= 378551
	}
	return code
}

//by Justin Sobel
func JShash(str []byte) uint32 {
	var code = uint32(1315423911)
	for i := 0; i < len(str); i++ {
		code ^= (code << 5) + uint32(str[i]) + (code >> 2)
	}
	return code
}

//by Arash Partow
func APhash(str []byte) uint32 {
	var code = uint32(0)
	var tick = true
	for i := 0; i < len(str); i++ {
		if tick {
			code ^= (code << 7) ^ uint32(str[i]) ^ (code >> 3)
		} else {
			code ^= ^((code << 11) ^ uint32(str[i]) ^ (code >> 5))
		}
		tick = !tick
	}
	return code
}

/*
Bible.txt [31102/49157]
JenkinsHash: <6> [16612 10472 3219 688 105 6]
MurmurHash:  <7> [16629 10316 3351 712 75 19]
BKDRhash:    <7> [16428 10614 3201 740 100 19]
SDBMhash:    <6> [16472 10484 3351 684 105 6]
DJBhash:     <5> [16425 10594 3303 660 120 0]
DJB2hash:    <6> [16544 10260 3336 776 180 6]
FNVhash:     <6> [16492 10412 3312 732 130 24]
RShash:      <5> [16409 10628 3264 736 65 0]
JShash:      <6> [16609 10370 3231 716 140 36]
APhash:      <6> [16612 10300 3351 696 125 18]

圣经.txt [31102/49157]
JenkinsHash: <6> [16500 10482 3363 644 95 18]
MurmurHash:  <6> [16460 10416 3402 672 140 12]
BKDRhash:    <6> [16501 10412 3345 712 120 12]
SDBMhash:    <6> [16388 10290 3576 684 140 24]
DJBhash:     <7> [16740 10402 3201 616 130 13]
DJB2hash:    <6> [16529 10552 3147 716 140 18]
FNVhash:     <6> [16392 10476 3366 748 90 30]
RShash:      <6> [16537 10462 3261 756 80 6]
JShash:      <6> [16500 10346 3360 788 90 18]
APhash:      <5> [16527 10502 3228 740 105 0]
*/
