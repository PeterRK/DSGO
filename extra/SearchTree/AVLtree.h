#pragma once

class AVLtree : NonCopyable {
private:
	struct Node {
		int		key;
		int8_t	state;
		Node*	parent;
		Node*	left;
		Node*	right;

		void hookLeft(Node* child, void* hint=(void*)-1) {
			if (hint != nullptr || child != nullptr) {
				child->parent = this;
			}
			this->left = child;
		}
		void hookRight(Node* child, void* hint = (void*)-1) {
			if (hint != nullptr || child != nullptr) {
				child->parent = this;
			}
			this->right = child;
		}
	};
	Node* m_root;
	Allocator<Node> m_pool;
	Node* newNode(Node* parent, int key);
	static Node* Rotate(Node* G, bool& stop);

public:
	AVLtree(void) : m_root(nullptr) {}

	bool isEmpty(void) const { return m_root == nullptr; }
	bool search(int key) const;
	bool insert(int key);
	bool remove(int key);
};