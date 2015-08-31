#include "tricks.h"
#include "RadixTree.h"

inline unsigned RadixTree::Cut(Key key, unsigned i)
{
	return (key >> ((UINT_SZ - STEP) - i*STEP)) & MASK;
}
inline RadixTree::Node* RadixTree::NewNode(void)
{
	Node* node = m_pool.allocate();
	for (unsigned i = 0; i < SIZE; i++) {
		node->kids[i] = nullptr;
	}
	return node;
}
RadixTree::RadixTree(void)
{
	m_root = NewNode();
}

bool RadixTree::Insert(Key key, void* ptr)
{
	Node* root = m_root;
	for (unsigned i = 0; i < DEPTH - 1; i++) {
		unsigned idx = Cut(key, i);
		if (root->kids[idx] == nullptr) {
			root->kids[idx] = NewNode();
		}
		root = root->kids[idx];
	}
	unsigned idx = key & MASK;
	if (root->kids[idx] != nullptr) return false;
	root->kids[idx] = static_cast<Node*>(ptr);
	return true;
}

const void* RadixTree::Search(Key key)
const {
	const Node* root = m_root;
	for (unsigned i = 0; i < DEPTH && root != nullptr; i++) {
		root = root->kids[Cut(key, i)];
	}
	return root;
}

void* RadixTree::Remove(Key key)
{
	Node* path[DEPTH];
	path[0] = m_root;
	for (unsigned i = 0; i < DEPTH - 1; i++) {
		path[i + 1] = path[i]->kids[Cut(key, i)];
		if (path[i + 1] == nullptr) return nullptr;
	}
	unsigned idx = key & MASK;
	void* ptr = path[DEPTH - 1]->kids[idx];
	if (ptr != nullptr) {
		path[DEPTH - 1]->kids[idx] = nullptr;
		for (unsigned i = DEPTH - 1; i != 0; i--) {
			unsigned j = 0;
			while (j < SIZE && path[i]->kids[j] == nullptr) j++;
			if (j == SIZE) { //ШЋПе
				unsigned idx = Cut(key, i - 1);
				m_pool.deallocate(path[i - 1]->kids[idx]);
				path[i - 1]->kids[idx] = nullptr;
			}
		}
	}
	return ptr;
}
