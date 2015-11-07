package array

import (
	"errors"
	"math/rand"
	"time"
)

// 选取第k大数
func Pick(list []int, k int) (int, error) {
	if k <= 0 || k > len(list) {
		return 0, errors.New("out of range")
	}
	for begin, end := 0, len(list); begin < end-1; {
		var pivot = list[(begin+end)/2] //一定要选偏后
		var a, b = begin, end - 1
		for { //注意对称性
			for list[a] < pivot {
				a++
			}
			for list[b] > pivot {
				b--
			}
			if a >= b {
				break
			}
			list[a], list[b] = list[b], list[a]
			a++
			b--
		}
		if k <= a {
			end = a
		} else {
			begin = a
		}
	}
	return list[k-1], nil
}

// 随机排列
func Randomize(list []int) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < len(list); i++ {
		var j = rand.Int() % (i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

// 随机组合，前n项为结果
func RandomPart(list []int, n int) {
	var size = len(list)
	if n > 0 && n < size {
		rand.Seed(time.Now().Unix())
		for i := n; i < size; i++ {
			var j = rand.Int() % (i + 1)
			list[i], list[j] = list[j], list[i]
		}
	}
}
