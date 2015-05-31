package bptree

type stackNode struct {
	pt  *index
	idx int
}
type stack struct {
	core []stackNode
}

func (s *stack) clear() {
	s.core = s.core[:0]
}
func (s *stack) isEmpty() bool {
	return len(s.core) == 0
}
func (s *stack) push(pt *index, idx int) {
	s.core = append(s.core, stackNode{pt, idx})
}
func (s *stack) pop() (pt *index, idx int) {
	var sz = len(s.core) - 1
	var unit = s.core[sz]
	s.core = s.core[:sz]
	return unit.pt, unit.idx
}
