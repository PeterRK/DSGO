#pragma once

//#define WEIGHTED_AVL

class AVLtree : NonCopyable {
private:
	struct Node {
		int			key;
#ifdef WEIGHTED_AVL
		unsigned 	weight : sizeof(unsigned)*8 - 3;
		int			state : 3;
#else
		int			state;
#endif
		Node*		parent;
		Node*		left;
		Node*		right;

		void hookLeft(Node* child, void* hint = (void*)-1) {
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
#ifdef WEIGHTED_AVL
		unsigned realWeight(void) {
			return this == nullptr ? 0 : weight;
		}
		unsigned subRank(void) {
			return left->realWeight() + 1;
		}
#else
		unsigned realWeight(void) { return 0; }
		unsigned subRank(void) { return 0; }
#endif
	};
	Node* m_root;
	Allocator<Node> m_pool;
	Node* newNode(Node* parent, int key);
	static Node* Rotate(Node* G, bool& stop);

public:
	AVLtree(void) : m_root(nullptr) {}

	bool isEmpty(void) const { return m_root == nullptr; }
	int search(int key) const;
	int insert(int key);
	int remove(int key);
};
