package hashtable

//此数列借鉴自SGI STL
var size_primes = []uint{
	0, 53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317,
	196613, 393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843,
	50331653, 1610612741, 3221225473, 4294967291}

func nextSize(size uint) (newsz uint, ok bool) {
	var start, end = 0, len(size_primes)
	for start < end {
		var mid = (start + end) / 2
		if size > size_primes[mid] {
			start = mid + 1
		} else if size < size_primes[mid] {
			end = mid
		} else {
			return size_primes[mid], true
		}
	}
	return size, false
}

type Node struct {
	name string
	next *Node
}

type HashTable struct {
	hash func(str string) uint
}
