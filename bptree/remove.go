package bptree

//B+树的删除相当复杂
import (
	"unsafe"
)

//要求peer为leaf后的节点
func (node *leaf) combine(peer *leaf) bool {
	var total = node.cnt + peer.cnt
	if total <= LEAF_HALF+LEAF_QUARTER {
		for i := node.cnt; i < total; i++ {
			node.data[i] = peer.data[i-node.cnt]
		}
		node.cnt, peer.cnt = total, 0
		node.next = peer.next
		return true
	}
	//分流而不合并
	var node_size = total / 2
	if node.cnt == node_size {
		return false
	}
	var peer_size = total - node_size
	if peer.cnt > peer_size { //后填前
		var diff = peer.cnt - peer_size
		for i := node.cnt; i < node_size; i++ {
			node.data[i] = peer.data[i-node.cnt]
		}
		for i := 0; i < peer_size; i++ {
			peer.data[i] = peer.data[i+diff]
		}
	} else { //前填后
		var diff = peer_size - peer.cnt
		for i := peer.cnt - 1; i >= 0; i-- {
			peer.data[i+diff] = peer.data[i]
		}
		for i := node.cnt - 1; i >= node_size; i-- {
			peer.data[i-node_size] = node.data[i]
		}
	}
	node.cnt, peer.cnt = node_size, peer_size
	return false
}

func (node *index) combine(peer *index) bool {
	var total = (-node.cnt) + (-peer.cnt)
	if total <= INDEX_HALF+INDEX_QUARTER {
		for i := (-node.cnt); i < total; i++ {
			node.data[i] = peer.data[i-(-node.cnt)]
			node.kids[i] = peer.kids[i-(-node.cnt)]
		}
		node.cnt, peer.cnt = -total, 0
		return true
	}
	//分流而不合并
	var node_size = total / 2
	if (-node.cnt) == node_size {
		return false
	}
	var peer_size = total - node_size
	if (-peer.cnt) > peer_size { //后填前
		var diff = (-peer.cnt) - peer_size
		for i := (-node.cnt); i < node_size; i++ {
			node.data[i] = peer.data[i-(-node.cnt)]
			node.kids[i] = peer.kids[i-(-node.cnt)]
		}
		for i := 0; i < peer_size; i++ {
			peer.data[i] = peer.data[i+diff]
			peer.kids[i] = peer.kids[i+diff]
		}
	} else { //前填后
		var diff = peer_size - (-peer.cnt)
		for i := (-peer.cnt) - 1; i >= 0; i-- {
			peer.data[i+diff] = peer.data[i]
			peer.kids[i+diff] = peer.kids[i]
		}
		for i := (-node.cnt) - 1; i >= node_size; i-- {
			peer.data[i-node_size] = node.data[i]
			peer.kids[i-node_size] = node.kids[i]
		}
	}
	node.cnt, peer.cnt = -node_size, -peer_size
	return false
}

func (tree *Tree) Remove(key int) { //积极合并
	if tree.root == nil ||
		key > tree.root.ceil() {
		return
	}
	tree.path.clear()

	var target = tree.root
	for target.cnt < 0 { //index节点
		var idx = target.locate(key)
		tree.path.push(target, idx)
		target = target.kids[idx]
	}
	var node = (*leaf)(unsafe.Pointer(target)) //叶节点
	var place = node.locate(key)
	if key != node.data[place] { //不存在
		return
	}
	node.cnt--
	for i := place; i < node.cnt; i++ {
		node.data[i] = node.data[i+1]
	}

	if tree.path.isEmpty() { //孤叶，不寻求合并
		if node.cnt == 0 {
			tree.root, tree.head = nil, nil
		}
		return
	}
	var shrink = false
	var path_ceil int
	if place == node.cnt {
		shrink, path_ceil = true, node.data[node.cnt-1]
	}

	//此后保证node没有上界问题，parent可能有
	var parent *index
	parent, place = tree.path.pop() //parent一定有至少2个child
	if node.cnt > LEAF_QUARTER {
		goto adjust_ceil
	}

	if place == (-parent.cnt)-1 { //只能与前项合并
		var before = (*leaf)(unsafe.Pointer(parent.kids[place-1]))
		var combined = before.combine(node)
		parent.data[place-1] = before.data[before.cnt-1] //调整中界
		if !combined {
			goto adjust_ceil
		}
	} else { //与后项合并
		var after = (*leaf)(unsafe.Pointer(parent.kids[place+1]))
		var combined = node.combine(after)
		parent.data[place] = node.data[node.cnt-1] //调整中界
		if !combined {
			return
		}
		place++
		shrink = false
	}

	for {
		parent.cnt++
		for i := place; i < (-parent.cnt); i++ {
			parent.data[i] = parent.data[i+1]
			parent.kids[i] = parent.kids[i+1]
		} //处理过后没有上界问题
		var node = parent //此后代码与之前类似，但node的类型已经不同

		if tree.path.isEmpty() { //根Index没法合并
			if (-node.cnt) == 1 { //但可以节缩
				tree.root = node.kids[0]
			}
			return
		}
		parent, place = tree.path.pop() //parent一定有至少2个child
		if (-node.cnt) > INDEX_QUARTER {
			goto adjust_ceil
		}
		if place == (-parent.cnt)-1 { //只能与前项合并
			var before = parent.kids[place-1]
			var combined = before.combine(node)
			parent.data[place-1] = before.data[(-before.cnt)-1] //调整中界
			if !combined {
				goto adjust_ceil
			}
		} else { //与后项合并
			var after = parent.kids[place+1]
			var combined = node.combine(after)
			parent.data[place] = node.data[(-node.cnt)-1] //调整中界
			if !combined {
				return
			}
			place++
			shrink = false
		}
	}
	return

adjust_ceil:
	if shrink {
		parent.data[place] = path_ceil
		for place == (-parent.cnt)-1 && //级联
			!tree.path.isEmpty() {
			parent, place = tree.path.pop()
			parent.data[place] = path_ceil
		}
	}
}
