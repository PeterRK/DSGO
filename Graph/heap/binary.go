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
	var hp = new(binaryHeap)
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
	var unit = hp.core[pos]
	for pos > 0 {
		var parent = (pos - 1) / 2
		if hp.core[parent].Dist <= unit.Dist {
			break
		}
		hp.core[pos] = hp.core[parent]
		hp.core[pos].off, pos = pos, parent
	}
	hp.core[pos], unit.off = unit, pos
}

func (hp *binaryHeap) Push(unit *memo) {
	var pos = len(hp.core)
	hp.core = append(hp.core, unit)
	//	unit.off = pos
	hp.shiftUp(pos)
}

func (hp *binaryHeap) Pop() *memo {
	var size = len(hp.core)
	//	if size == 0 {
	//		return nil
	//	}
	var result = hp.core[0]
	if size == 1 {
		hp.core = hp.core[:0]
		return result
	}

	var unit = hp.core[size-1]
	hp.core = hp.core[:size-1]

	var pos, kid, last = 0, 1, size - 2
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
