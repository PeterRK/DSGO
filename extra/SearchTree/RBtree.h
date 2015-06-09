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

		Node* hook(Node* child) {
			child->parent(this);
			return child;
		}
		Node* tryHook(Node* child) {
			if (child != NULL) {
				child->parent(this);
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
