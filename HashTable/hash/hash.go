package hash

type HashTable interface {
	Size() int
	IsEmpty() bool
	Insert(key string) bool
	Search(key string) bool
	Remove(key string) bool
}

//by Brian Kernighan & Dennis Ritchie
func BKDRhash(str string) uint {
	const factor uint = 131 //31、131、1313、 13131、131313
	var data = []byte(str)
	var code uint = 0
	for _, ch := range data {
		code = code*factor + uint(ch)
	}
	return code
}

//from SDBM
func SDBMhash(str string) uint {
	var data = []byte(str)
	var code uint = 0
	for _, ch := range data {
		code = code*65599 + uint(ch)
	}
	return code
}

//by Daniel J. Bernstein
func DJBhash(str string) uint {
	var data = []byte(str)
	var code uint = 5381
	for _, ch := range data {
		code = code*33 + uint(ch)
	}
	return code
}
func DJB2hash(str string) uint {
	var data = []byte(str)
	var code uint = 5381
	for _, ch := range data {
		code = code*33 ^ uint(ch)
	}
	return code
}

//from Unix
func FNVhash(str string) uint {
	var data = []byte(str)

	var code uint = 2166136261
	for _, ch := range data {
		code = code*16777619 ^ uint(ch)
	}
	return code
}

//Robert Sedgwicks
func RShash(str string) uint {
	var data = []byte(str)

	var magic uint = 63689
	var code uint = 0
	for _, ch := range data {
		code = code*magic + uint(ch)
		magic *= 378551
	}
	return code
}

//by Justin Sobel
func JShash(str string) uint {
	var data = []byte(str)

	var code uint = 1315423911
	for _, ch := range data {
		code ^= (code << 5) + uint(ch) + (code >> 2)
	}
	return code
}

//by Arash Partow
func APhash(str string) uint {
	var data = []byte(str)

	var code uint = 0
	var tick = true
	for _, ch := range data {
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
Bible.txt
BKDRhash: [max=6] [16429 5204 1170 158 21 3]	[31102/49157]
SDBMhash: [max=7] [16407 5248 1100 188 22 6]	[31102/49157]
DJBhash:  [max=6] [16396 5230 1129 184 21 3]	[31102/49157]
DJB2hash: [max=6] [16422 5211 1127 181 27 3]	[31102/49157]
FNVhash:  [max=7] [16609 5155 1115 168 27 5]	[31102/49157]
RShash:   [max=6] [16727 5100 1127 158 30 2]	[31102/49157]
JShash:   [max=6] [16602 5216 1061 191 23 1]	[31102/49157]
APhash:   [max=6] [16641 5239 1055 171 22 4]	[31102/49157]

圣经.txt
BKDRhash: [max=6] [16416 5215 1120 191 24 2]	[31102/49157]
SDBMhash: [max=6] [16607 5252 1068 164 25 1]	[31102/49157]
DJBhash:  [max=6] [16551 5273 1080 167 17 2]	[31102/49157]
DJB2hash: [max=5] [16353 5291 1107 189 18 0]	[31102/49157]
FNVhash:  [max=6] [16491 5221 1120 168 25 2]	[31102/49157]
RShash:   [max=7] [16403 5308 1092 167 24 3]	[31102/49157]
JShash:   [max=7] [16575 5274 1069 165 21 1]	[31102/49157]
APhash:   [max=6] [16365 5305 1104 171 25 1]	[31102/49157]

圣经UTF8.txt
BKDRhash: [max=6] [16562 5326 1049 154 19 5]	[31102/49157]
SDBMhash: [max=6] [16521 5312 1056 172 19 1]	[31102/49157]
DJBhash:  [max=6] [16688 5219 1071 163 21 1]	[31102/49157]
DJB2hash: [max=6] [16525 5180 1129 167 30 2]	[31102/49157]
FNVhash:  [max=6] [16425 5270 1091 177 24 6]	[31102/49157]
RShash:   [max=6] [16487 5195 1133 173 22 4]	[31102/49157]
JShash:   [max=6] [16543 5264 1098 164 15 1]	[31102/49157]
APhash:   [max=6] [16515 5187 1113 177 32 1]	[31102/49157]
*/
