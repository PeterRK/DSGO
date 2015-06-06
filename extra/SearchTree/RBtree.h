#pragma once

class RBtree : NonCopyable {
private:
	struct Node {
		int			key;
		uintptr_t	black : 1;
		uintptr_t	parent : sizeof(uintptr_t)*8 - 1;
		Node*		left;
		Node*		right;

		Node* hook(Node* child) {
			child->parent = reinterpret_cast<uintptr_t>(this);
			return child;
		}
		Node* tryHook(Node* child) {
			if (child != NULL) {
				child->parent = reinterpret_cast<uintptr_t>(this);
			}
			return child;
		}
	};
	Node* m_root;
	Allocator<Node> m_pool;
	Node* newNode(Node* parent, int key);
	void hookSubTree(Node* super, Node* root);
	void adjustAfterDelete(Node* G, int key);

public:
	RBtree(void) : m_root(NULL) {}

	bool isEmpty(void) const { return m_root == NULL; }
	bool search(int key) const;
	bool insert(int key);
	bool remove(int key);
};