package hash

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ConflictRate(data []string, size uint, fn func([]byte) uint) float64 {
	var cnt = uint(0)
	var book = make([]bool, size)
	for i := uint(0); i < size; i++ {
		book[i] = false
	}
	for _, str := range data {
		var index = fn([]byte(str)) % size
		if book[index] {
			cnt++
		} else {
			book[index] = true
		}
	}
	return float64(cnt) / float64(len(data))
}

func TryByFile(filenme string) {
	if data, err := FetchLines(filenme); err != nil {
		fmt.Println("fail to open file:", filenme)
	} else {
		var size = bestSize(uint(len(data)))
		fmt.Printf("%s [%d/%d]\n", filenme, len(data), size)
		ShowBucketCounts("BKDRhash:", data, size, BKDRhash)
		ShowBucketCounts("SDBMhash:", data, size, SDBMhash)
		ShowBucketCounts("DJBhash: ", data, size, DJBhash)
		ShowBucketCounts("DJB2hash:", data, size, DJB2hash)
		ShowBucketCounts("FNVhash: ", data, size, FNVhash)
		ShowBucketCounts("RShash:  ", data, size, RShash)
		ShowBucketCounts("JShash:  ", data, size, JShash)
		ShowBucketCounts("APhash:  ", data, size, APhash)
		fmt.Println()
	}
}
func ShowBucketCounts(msg string, data []string, size uint, fn func([]byte) uint) {
	var vec, top = BucketCounts(data, size, fn)
	fmt.Printf("%s <%d> %v\n", msg, top, vec)
}
func BucketCounts(data []string, size uint, fn func([]byte) uint) (vec [6]uint, top uint) {
	var book = make([]uint, size)
	for i := uint(0); i < size; i++ {
		book[i] = 0
	}
	for _, str := range data {
		var index = fn([]byte(str)) % size
		book[index]++
	}
	top = 1
	for i := 0; i < len(vec); i++ {
		vec[i] = 0
	}
	for _, num := range book {
		if num != 0 {
			if num > top {
				top = num
			}
			if num > uint(len(vec)) {
				vec[len(vec)-1] += num
			} else {
				vec[num-1] += num
			}
		}
	}
	return vec, top
}

func FetchLines(name string) (data []string, err error) {
	var file *os.File
	if file, err = os.Open(name); err != nil {
		return data, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var str string
	for {
		if str, err = reader.ReadString('\n'); err != nil {
			break
		} else {
			data = append(data, str)
		}
	}
	if err == io.EOF {
		err = nil
	}
	return data, err
}

//此数列借鉴自SGI STL
var sz_primes = []uint{
	17, 29, 53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613,
	393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 1610612741}

func bestSize(hint uint) uint {
	var start, end = 0, len(sz_primes)
	for start < end {
		var mid = (start + end) / 2
		if hint < sz_primes[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(sz_primes) {
		start = len(sz_primes) - 1
	}
	return sz_primes[start]
}
