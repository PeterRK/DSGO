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

func DepthFirstSearch(root *Node, doit func(int)) {

}

func BreadthFirstSearch(root *Node, doit func(int)) {

}
