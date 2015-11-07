package tree

type BinaryNode struct {
	key         int
	left, right *BinaryNode
}

func DepthFirstSearch(root *BinaryNode, doit func(int)) {
	if root != nil {
		doit(root.key) //前序
		DepthFirstSearch(root.left, doit)
		//doit(root.key) //中序
		DepthFirstSearch(root.right, doit)
		//doit(root.key) //后序
	}
}

type Node struct {
	key   int
	child *Node
	peer  *Node
}

func BreadthFirstSearch(root *Node, doit func(int)) {
	if root != nil {
		var q = newQ()
		for err := error(nil); err == nil; root, err = q.pop() {
			for kid := root.child; kid != nil; kid = kid.peer {
				q.push(kid)
			}
		}
	}
}

type TreeNode struct {
	key         int
	parent      *TreeNode
	left, right *TreeNode
}

func BuildBalanceTree(list []int, parent *TreeNode) *TreeNode {
	var size = len(list)
	if size == 0 {
		return nil
	}
	var node = new(TreeNode)
	node.parent = parent
	//node.left, node.right = nil, nil
	var m = size / 2
	node.key = list[m]
	if size != 1 {
		node.left = BuildBalanceTree(list[:m], node)
		node.right = BuildBalanceTree(list[m+1:], node)
	}
	return node
}

func MoveForward(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		var kid = node.right
		for kid.left != nil {
			kid = kid.left
		}
		return kid
	}
	var parent = node.parent
	for parent != nil && node == parent.right {
		node, parent = parent, parent.parent
	}
	return parent
}

func MoveBackward(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.left != nil {
		var kid = node.left
		for kid.right != nil {
			kid = kid.right
		}
		return kid
	}
	var parent = node.parent
	for parent != nil && node == parent.left {
		node, parent = parent, parent.parent
	}
	return parent
}
