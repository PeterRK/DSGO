#pragma once

class WeightedAVL : NonCopyable {
private:
	struct Node {
		int			key;
		unsigned 	weight : sizeof(unsigned)*8 - 3;
		int			state : 3;
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

		unsigned realWeight(void) {
			return this == nullptr ? 0 : weight;
		}
		unsigned subRank(void) {
			return left->realWeight() + 1;
		}
	};
	Node* m_root;
	Allocator<Node> m_pool;
	Node* newNode(Node* parent, int key);
	static Node* Rotate(Node* G, bool& stop);

public:
	WeightedAVL(void) : m_root(nullptr) {}

	bool isEmpty(void) const { return m_root == nullptr; }
	int search(int key) const;
	int insert(int key);
	int remove(int key);
};
