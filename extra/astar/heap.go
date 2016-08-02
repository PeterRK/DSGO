package astar

type heap struct {
	core []*memo
}

func (hp *heap) isEmpty() bool {
	return len(hp.core) == 0
}

func (hp *heap) floatUp(unit *memo, delta uint) {
	unit.weight -= delta
	hp.shiftUp(unit.off)
}
func (hp *heap) shiftUp(pos int) {
	var unit = hp.core[pos]
	for pos > 0 {
		var parent = (pos - 1) / 2
		if hp.core[parent].weight <= unit.weight {
			break
		}
		hp.core[pos] = hp.core[parent]
		hp.core[pos].off, pos = pos, parent
	}
	hp.core[pos], unit.off = unit, pos
}

func (hp *heap) push(unit *memo) {
	var pos = len(hp.core)
	hp.core = append(hp.core, unit)
	//	unit.off = pos
	hp.shiftUp(pos)
}

func (hp *heap) pop() *memo {
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
		if hp.core[kid+1].weight < hp.core[kid].weight {
			kid++
		}
		if unit.weight <= hp.core[kid].weight {
			break
		}
		hp.core[pos] = hp.core[kid]
		hp.core[pos].off = pos
		pos, kid = kid, kid*2+1
	}
	if kid == last && unit.weight > hp.core[kid].weight {
		hp.core[pos] = hp.core[kid]
		hp.core[pos].off, pos = pos, kid
	}
	hp.core[pos], unit.off = unit, pos

	return result
}
