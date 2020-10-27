package array

// 最长递增子序列
func LongestIncreasingSubsequence(src []int) int {
	if len(src) < 2 {
		return len(src)
	} //memo[i]记录某个长度为(i+1)的递增串的可能的最小尾数
	memo := []int{src[0]}
	for i := 1; i < len(src); i++ {
		place := SearchFirst(memo, src[i])
		if place == len(memo) {
			memo = append(memo, src[i])
		} else {
			memo[place] = src[i]
		}
	}
	return len(memo) //memo不是真正的候选序列
}

func LongestIncreasingSubsequenceX(src []int) []int {
	if len(src) < 2 {
		return src
	}
	type Mark struct {
		prev   int
		length uint
		best   bool
	}
	memo := make([]Mark, 1, len(src))
	memo[0] = Mark{-1, 1, true}
	tail := 0
	for i := 1; i < len(src); i++ {
		mark := Mark{-1, 1, false}
		for j := i - 1; j >= 0; j-- {
			if src[i] > src[j] && memo[j].length >= mark.length {
				mark.prev = j
				mark.length = memo[j].length + 1
			}
			if memo[j].best {
				break
			}
		}
		if mark.length >= memo[tail].length {
			mark.best = true
			tail = i
		}
		memo = append(memo, mark)
	}
	result := make([]int, memo[tail].length)
	for i := len(result); tail >= 0; tail = memo[tail].prev {
		i--
		result[i] = src[tail]
	}
	return result
}

// 最大子段和，全为负数时，最佳为空子段
func MaximumIntervalSum(list []int) int {
	best, sum := 0, 0
	for _, num := range list {
		sum += num
		if sum < 0 {
			sum = 0
		} else if sum > best {
			best = sum
		}
	}
	return best
}

// 最大子段和及对应区间
func MaximumIntervalSumX(list []int) (value int, start, end int) {
	value, start, end = 0, 0, 0
	sum, mark := -1, -1
	for i, num := range list {
		if sum < 0 {
			sum, mark = num, i
		} else {
			sum += num
		}
		if sum > value {
			value, start, end = sum, mark, i+1
		}
	}
	return value, start, end
}
