package path

import (
	"math"
)

//处理距离矩阵(MaxInt指不通)。
//复杂度为O(V^3)。
//可以处理有向图和负权边，但不能判定负回路。
func FloydWarshall(matrix [][]int) {
	size := len(matrix)
	for k := 0; k < size; k++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if matrix[i][k] != math.MaxInt && matrix[k][j] != math.MaxInt {
					distance := matrix[i][k] + matrix[k][j]
					if distance < matrix[i][j] {
						matrix[i][j] = distance
					}
				}
			}
		}
	}
}
