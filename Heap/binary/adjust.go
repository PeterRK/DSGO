package binary

func (hp *Heap) adjustDown(spot int) {
	var size = len(hp.core)
	var key = hp.core[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < size {
		var kid = 0
		if hp.core[left] < hp.core[right] {
			kid = left
		} else {
			kid = right
		}
		if key <= hp.core[kid] {
			goto LabelOver
		}
		hp.core[spot] = hp.core[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == size && key > hp.core[left] {
		hp.core[spot], hp.core[left] = hp.core[left], key
		return
	}
LabelOver:
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
