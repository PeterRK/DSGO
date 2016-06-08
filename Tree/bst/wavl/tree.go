package wavl

//为了方便查询排行，加入了数量记录

type node struct {
	state  int8 //(2), 1, 0, -1, (-2)
	weight int32
	key    int
	parent *node
	left   *node
	right  *node
}
type Tree struct {
	root *node
}

func (tr *Tree) IsEmpty() bool {
	return tr.root == nil
}

//找到返回序号（从1开始），没有返回-1
func (tr *Tree) Search(key int) int {
	var target, base = tr.root, int32(0)
	for target != nil {
		if key == target.key {
			return int(base + target.subRank())
		}
		if key < target.key {
			target = target.left
		} else {
			base += target.subRank()
			target = target.right
		}
	}
	return -1
}

func (parent *node) tryHook(child *node) *node {
	if child != nil {
		child.parent = parent
	}
	return child
}
func (parent *node) hook(child *node) *node {
	child.parent = parent
	return child
}
func (unit *node) realWeight() int32 {
	if unit == nil {
		return 0
	}
	return unit.weight
}
func (unit *node) subRank() int32 {
	return unit.left.realWeight() + 1
}
