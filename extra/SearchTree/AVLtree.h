#pragma once

class AVLtree : NonCopyable {
private:
	struct Node {
		int		key;
		int8_t	state;
		Node*	parent;
		Node*	left;
		Node*	right;

		Node* hook(Node* child) {
			child->parent = this;
			return child;
		}
		Node* tryHook(Node* child) {
			if (child != NULL) {
				child->parent = this;
			}
			return child;
		}
	};
	Node* m_root;
	Allocator<Node> m_pool;
	Node* newNode(Node* parent, int key);
	static Node* Rotate(Node* G, bool& stop);

public:
	AVLtree(void) : m_root(NULL) {}

	bool isEmpty(void) const { return m_root == NULL; }
	bool search(int key) const;
	bool insert(int key);
	bool remove(int key);
};

