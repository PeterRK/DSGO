package avlt

type stackNode struct {
	pt *node
	lf bool
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
func (s *stack) push(pt *node, lf bool) {
	s.core = append(s.core, stackNode{pt, lf})
}
func (s *stack) pop() (pt *node, lf bool) {
	var sz = len(s.core) - 1
	var unit = s.core[sz]
	s.core = s.core[:sz]
	return unit.pt, unit.lf
}
