package binary

func (hp *Heap) adjustDown(spot int) {
	var key = hp.core[spot]
	var kid, last = spot*2 + 1, len(hp.core) - 1
	for kid < last {
		if hp.core[kid+1] < hp.core[kid] {
			kid++
		}
		if key <= hp.core[kid] {
			goto Label_OVER
		}
		hp.core[spot] = hp.core[kid]
		spot, kid = kid, kid*2+1
	}
	if kid == last && key > hp.core[kid] {
		hp.core[spot], spot = hp.core[kid], kid
	}
Label_OVER:
	hp.core[spot] = key
}

func (hp *Heap) adjustUp(spot int) {
	var key = hp.core[spot]
	for spot > 0 {
		var parent = (spot - 1) / 2
		if hp.core[parent] <= key {
			break
		}
		hp.core[spot], spot = hp.core[parent], parent
	}
	hp.core[spot] = key
}
