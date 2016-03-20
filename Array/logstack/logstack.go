package logstack

type LogStack interface {
	Insert(key int)
	Delete(key int)
	Search(key int) bool
}

type logStack struct {
	buffer layer
	limit  uint
	stack  []layer
}

func NewLogStack(limit uint) LogStack {
	if limit < 4 {
		limit = 4
	}
	var obj = new(logStack)
	obj.limit = limit
	return obj
}

func (s *logStack) change(key int, mark bool) {
	s.buffer.change(key, mark)
	if s.buffer.size() == s.limit {
		s.stack = append(s.stack, s.buffer)
		s.buffer.reset()
	}
}
func (s *logStack) Insert(key int) {
	s.change(key, true)
}
func (s *logStack) Delete(key int) {
	s.change(key, false)
}

func (s *logStack) Search(key int) bool {
	var found = s.buffer.search(key)
	if found == 0 && len(s.stack) != 0 {
		found = s.stack[len(s.stack)-1].search(key)
	}
	for found == 0 && len(s.stack) > 1 {
		found = s.stack[len(s.stack)-2].search(key)
		s.stack[len(s.stack)-2].merge(&s.stack[len(s.stack)-1])
		s.stack = s.stack[:len(s.stack)-1]
	}
	if len(s.stack) == 1 {
		s.stack[0].compact()
	}
	return found == 1
}

func (s *logStack) debug() {
	s.buffer.print()
	for i := len(s.stack) - 1; i >= 0; i-- {
		s.stack[i].print()
	}
}
