package flow

//筛分层次，生成分层残图，复杂度为(V^2)。
func separate(shadow [][]edge, matrix [][]uint, q *queue, memo []uint) (ok bool) {
	var size = len(matrix)
	const FAKE_LEVEL = ^uint(0)
	for i := 1; i < size; i++ {
		memo[i] = FAKE_LEVEL
	}
	memo[0] = 0

	q.clear()
	q.push(0)
	for !q.isEmpty() {
		var current = q.pop()
		if matrix[current][size-1] != 0 {
			memo[size-1] = memo[current] + 1
			goto Label_REACH //到终点层
		}
		for i := 1; i < size-1; i++ {
			if memo[i] == FAKE_LEVEL && matrix[current][i] != 0 {
				memo[i] = memo[current] + 1
				q.push(i)
			}
		}
	}
	return false

Label_REACH:
	for {
		var current, err = q.traceBack()
		if err != nil {
			break
		}
		//shadow[current] = shadow[current][:0]
		for i := 1; i < size; i++ {
			if memo[i] == memo[current]+1 && matrix[current][i] != 0 {
				shadow[current] = append(shadow[current], edge{next: i, val: matrix[current][i]})
				matrix[current][i] = 0
			}
		}
		if len(shadow[current]) == 0 {
			memo[current] = FAKE_LEVEL
		}
	}
	return true
}

func flushBack(shadow [][]edge, matrix [][]uint) {
	var size = len(shadow) //len(shadow) == len(matrix)-1
	for i := 1; i < size; i++ {
		if len(shadow[i]) != 0 {
			for _, path := range shadow[i] {
				matrix[i][path.next] += path.val
			}
			shadow[i] = shadow[i][:0]
		}
	}
}
