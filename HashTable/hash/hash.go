package hash

type HashTable interface {
	Size() int
	IsEmpty() bool
	Insert(key []byte) bool
	Search(key []byte) bool
	Remove(key []byte) bool
}

//以下算法搜集自网络，基本思路类似，效果相近

//by Brian Kernighan & Dennis Ritchie
func BKDRhash(str []byte) uint {
	const factor uint = 131 //31、131、1313、 13131、131313
	var code uint = 0
	for _, ch := range str {
		code = code*factor + uint(ch)
	}
	return code
}

//from SDBM
func SDBMhash(str []byte) uint {
	var code uint = 0
	for _, ch := range str {
		code = code*65599 + uint(ch)
	}
	return code
}

//by Daniel J. Bernstein
func DJBhash(str []byte) uint {
	var code uint = 5381
	for _, ch := range str {
		code = code*33 + uint(ch)
	}
	return code
}
func DJB2hash(str []byte) uint {
	var code uint = 5381
	for _, ch := range str {
		code = code*33 ^ uint(ch)
	}
	return code
}

//from Unix
func FNVhash(str []byte) uint {
	var code uint = 2166136261
	for _, ch := range str {
		code = code*16777619 ^ uint(ch)
	}
	return code
}

//Robert Sedgwicks
func RShash(str []byte) uint {
	var magic uint = 63689
	var code uint = 0
	for _, ch := range str {
		code = code*magic + uint(ch)
		magic *= 378551
	}
	return code
}

//by Justin Sobel
func JShash(str []byte) uint {
	var code uint = 1315423911
	for _, ch := range str {
		code ^= (code << 5) + uint(ch) + (code >> 2)
	}
	return code
}

//by Arash Partow
func APhash(str []byte) uint {
	var code uint = 0
	var tick = true
	for _, ch := range str {
		if tick {
			code ^= (code << 7) ^ uint(ch) ^ (code >> 3)
		} else {
			code ^= ^((code << 11) ^ uint(ch) ^ (code >> 5))
		}
		tick = !tick
	}
	return code
}

/*
Bible.txt [31102/49157]
BKDRhash: <6> [16619 10410 3195 732 140 6]
SDBMhash: <6> [16405 10684 3237 640 130 6]
DJBhash:  <6> [16473 10536 3267 740 80 6]
DJB2hash: <7> [16448 10470 3312 708 145 19]
FNVhash:  <5> [16583 10368 3342 704 105 0]
RShash:   <6> [16527 10502 3231 704 120 18]
JShash:   <6> [16404 10494 3378 664 150 12]
APhash:   <6> [16613 10492 3252 636 85 24]

圣经.txt [31102/49157]
BKDRhash: <5> [16350 10592 3291 744 125 0]
SDBMhash: <6> [16401 10618 3231 708 120 24]
DJBhash:  <6> [16606 10296 3327 752 115 6]
DJB2hash: <5> [16390 10468 3354 800 90 0]
FNVhash:  <8> [16373 10500 3381 732 90 26]
RShash:   <7> [16581 10396 3264 656 180 25]
JShash:   <6> [16542 10368 3372 692 110 18]
APhash:   <6> [16527 10364 3330 732 125 24]

圣经UTF8.txt [31102/49157]
BKDRhash: <7> [16534 10454 3369 612 120 13]
SDBMhash: <6> [16610 10616 3069 664 125 18]
DJBhash:  <6> [16551 10510 3357 592 80 12]
DJB2hash: <7> [16666 10192 3411 680 115 38]
FNVhash:  <6> [16459 10562 3249 688 120 24]
RShash:   <6> [16466 10440 3381 696 95 24]
JShash:   <6> [16641 10478 3279 608 90 6]
APhash:   <6> [16521 10590 3174 700 105 12]
*/
