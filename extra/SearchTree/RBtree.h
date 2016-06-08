#pragma once

class RBtree : NonCopyable {
private:
	struct Node {
		int			key;
		Node*		left;
		Node*		right;
		uintptr_t	black : 1;
		uintptr_t	_parent : sizeof(uintptr_t)* 8 - 1;
		Node* parent() const {
			return (Node*)(_parent << 1);
		}
		void parent(Node* pt) {
			_parent = ((uintptr_t)pt) >> 1;
		}

		void hookLeft(Node* child, void* hint = (void*)-1) {
			if (hint != nullptr || child != nullptr) {
				child->parent(this);
			}
			this->left = child;
		}
		void hookRight(Node* child, void* hint = (void*)-1) {
			if (hint != nullptr || child != nullptr) {
				child->parent(this);
			}
			this->right = child;
		}
	};
	Node* m_root;
	Allocator<Node> m_pool;
	Node* newNode(Node* parent, int key);
	void hookSubTree(Node* super, Node* root);
	void adjustAfterDelete(Node* G, int key);

public:
	RBtree(void) : m_root(nullptr) {}

	bool isEmpty(void) const { return m_root == nullptr; }
	bool search(int key) const;
	bool insert(int key);
	bool remove(int key);
};