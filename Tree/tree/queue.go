package tree

const piece_sz = 30

type piece struct {
	fw, bw *piece
	space  [piece_sz]int
}
type index struct {
	pt  *piece
	idx int
}
type queue struct {
	front, back index
	cnt         int
}
