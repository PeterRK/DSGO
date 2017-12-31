package heap

type memo struct {
	Vertex
	off int
}

type binaryHeap struct {
	core []*memo
}

func NewVectorBH(size int) []memo {
	return make([]memo, size)
}

func NewBinaryHeap(size int) *binaryHeap {
	hp := new(binaryHeap)
	hp.core = make([]*memo, 0, size)
	return hp
}

func (hp *binaryHeap) IsEmpty() bool {
	return len(hp.core) == 0
}

func (hp *binaryHeap) FloatUp(unit *memo, dist uint) {
	unit.Dist = dist
	hp.shiftUp(unit.off)
}
func (hp *binaryHeap) shiftUp(pos int) {
	unit := hp.core[pos]
	for pos > 0 {
		parent := (pos - 1) / 2
		if hp.core[parent].Dist <= unit.Dist {
			break
		}
		hp.core[pos] = hp.core[parent]
		hp.core[pos].off, pos = pos, parent
	}
	hp.core[pos], unit.off = unit, pos
}

func (hp *binaryHeap) Push(unit *memo) {
	pos := len(hp.core)
	hp.core = append(hp.core, unit)
	//	unit.off = pos
	hp.shiftUp(pos)
}

func (hp *binaryHeap) Pop() *memo {
	size := len(hp.core)
	//	if size == 0 {
	//		return nil
	//	}
	result := hp.core[0]
	if size == 1 {
		hp.core = hp.core[:0]
		return result
	}

	unit := hp.core[size-1]
	hp.core = hp.core[:size-1]

	pos, kid, last := 0, 1, size-2
	for kid < last {
		if hp.core[kid+1].Dist < hp.core[kid].Dist {
			kid++
		}
		if unit.Dist <= hp.core[kid].Dist {
			break
		}
		hp.core[pos] = hp.core[kid]
		hp.core[pos].off = pos
		pos, kid = kid, kid*2+1
	}
	if kid == last && unit.Dist > hp.core[kid].Dist {
		hp.core[pos] = hp.core[kid]
		hp.core[pos].off, pos = pos, kid
	}
	hp.core[pos], unit.off = unit, pos

	return result
}
