package binary

func (hp *Heap) adjustDown(pos int) {
	var key = hp.core[pos]
	var kid, last = pos*2 + 1, len(hp.core) - 1
	for kid < last {
		if hp.core[kid+1] < hp.core[kid] {
			kid++
		}
		if key <= hp.core[kid] {
			break
		}
		hp.core[pos] = hp.core[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && key > hp.core[kid] {
		hp.core[pos], pos = hp.core[kid], kid
	}
	hp.core[pos] = key
}

func (hp *Heap) adjustUp(pos int) {
	var key = hp.core[pos]
	for pos > 0 {
		var parent = (pos - 1) / 2
		if hp.core[parent] <= key {
			break
		}
		hp.core[pos], pos = hp.core[parent], parent
	}
	hp.core[pos] = key
}
