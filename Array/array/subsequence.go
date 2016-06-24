package array

// 最长递增子序列
func LongestIncreasingSubsequence(src []int) []int {
	var size = len(src)
	if size == 0 {
		return nil
	} //result[i]记录某个长度为(i+1)的递增串的可能的最小尾数
	var result = []int{src[0]}
	for i := 1; i < size; i++ {
		var place = SearchFirst(result, src[i])
		if place == len(result) {
			result = append(result, src[i])
		} else {
			result[place] = src[i]
		}
	}
	return result
}

// 最大子段和，全为负数时，最佳为空子段
func MaximumIntervalSum(list []int) int {
	var best, sum = 0, 0
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
	var sum, mark = -1, -1
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
