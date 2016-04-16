package binary

func (hp *Heap) adjustDown(root int) {
	var key = hp.core[root]
	var kid, last = root*2 + 1, len(hp.core) - 1
	for kid < last {
		if hp.core[kid+1] < hp.core[kid] {
			kid++
		}
		if key <= hp.core[kid] {
			break
		}
		hp.core[root] = hp.core[kid]
		root, kid = kid, kid*2+1
	}
	if kid == last && key > hp.core[kid] {
		hp.core[root], root = hp.core[kid], kid
	}
	hp.core[root] = key
}

func (hp *Heap) adjustUp(root int) {
	var key = hp.core[root]
	for root > 0 {
		var parent = (root - 1) / 2
		if hp.core[parent] <= key {
			break
		}
		hp.core[root], root = hp.core[parent], parent
	}
	hp.core[root] = key
}
