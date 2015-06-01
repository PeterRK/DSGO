package tree

type BinaryNode struct {
	key         int
	left, right *BinaryNode
}

type Node struct {
	key   int
	child *Node
	peer  *Node
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

func BreadthFirstSearch(root *Node, doit func(int)) {
	if root != nil {
		var q = newQ()
		for fail := false; !fail; root, fail = q.pop() {
			for kid := root.child; kid != nil; kid = kid.peer {
				q.push(kid)
			}
		}
	}
}
