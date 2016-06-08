#include <cstdint>
#include <cstddef>
#include <new>
#include "tricks.h"
#include "RBtree.h"
////////////////////////////////////////////////////////////////////////////
bool RBtree::search(int key)
const {
	Node* root = m_root;
	while (root != nullptr) {
		if (key == root->key) {
			return true;
		}
		root = (key < root->key) ? root->left : root->right;
	}
	return false;
}
////////////////////////////////////////////////////////////////////////////
inline void RBtree::hookSubTree(Node* super, Node* root) {
	if (super == nullptr) {
		root->parent(nullptr);
		m_root = root;
	} else {
		if (root->key < super->key) {
			super->hookLeft(root);
		} else {
			super->hookRight(root);
		}
	}
}
inline RBtree::Node* RBtree::newNode(RBtree::Node* parent, int key)
{
	Node* node = m_pool.allocate();
	node->key = key;
	node->black = false;
	node->parent(parent);
	node->left = nullptr;
	node->right = nullptr;
	return node;
}
////////////////////////////////////////////////////////////////////////////
bool RBtree::insert(int key)
{
	if (m_root == nullptr) {
		m_root = newNode(nullptr, key);
		m_root->black = true;
		return true;
	}

	Node* root = m_root;
	while(true) {
		if (key < root->key) {
			if (root->left == nullptr) {
				root->left = newNode(root, key);
				break;
			}
			root = root->left;
		} else if (key > root->key) {
			if (root->right == nullptr) {
				root->right = newNode(root, key);
				break;
			}
			root = root->right;
		} else { //key == root->key
			return false;
		}
	}

	//------------红叔模式------------
	//|      bG      |      rG      |
	//|     /  \     |     /  \     |
	//|   rP    rU   |   bP    bU   |
	//|   |          |   |          |
	//|   rC         |   rC         |

	//-----------------LL形式-----------------
	//|        bG        |        bP        |
	//|       /  \       |       /  \       |
	//|     rP    bU     |     rC     rG    |
	//|    /  \          |          /  \    |
	//|  rC    x         |         x    bU  |

	//-----------------LR形式-----------------
	//|        bG        |        bC        |
	//|       /  \       |       /  \       |
	//|     rP    bU     |     rP    rG     |
	//|    / \           |    / \    / \    |
	//|      rC          |       u  v   bU  |
	//|     /  \         |                  |
	//|    u    v        |                  |

	Node* P = root;
	while (!P->black) { //违反双红禁
		Node* G = P->parent(); //必然存在，根为黑，P非根
		Node* super = G->parent();
		if (key < G->key) {
			Node* U = G->right;
			if (U != nullptr && !U->black) { //红叔模式，变色解决
				P->black = true;
				U->black = true;
				if (super != nullptr) {
					G->black = false;
					P = G->parent();
					continue; //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if (key < P->key) { //LL
					G->hookLeft(P->right, nullptr);
					P->hookRight(G);
					G->black = false;
					P->black = true;
					hookSubTree(super, P);
				} else { //LR
					Node* C = P->right;
					P->hookRight(C->left, nullptr);
					G->hookLeft(C->right, nullptr);
					C->hookLeft(P);
					C->hookRight(G);
					G->black = false;
					C->black = true;
					hookSubTree(super, C);
				}
			}
		} else {
			Node* U = G->left;
			if (U != nullptr && !U->black) { //红叔模式，变色解决
				P->black = true;
				U->black = true;
				if (super != nullptr) {
					G->black = false;
					P = G->parent();
					continue; //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if (key > P->key) { //RR
					G->hookRight(P->left, nullptr);
					P->hookLeft(G);
					G->black = false;
					P->black = true;
					hookSubTree(super, P);
				} else { //RL
					Node* C = P->left;
					P->hookLeft(C->right, nullptr);
					G->hookRight(C->left, nullptr);
					C->hookRight(P);
					C->hookLeft(G);
					G->black = false;
					C->black = true;
					hookSubTree(super, C);
				}
			}
		}
		break; //变色时才需要循环
	}
	return true;
}
////////////////////////////////////////////////////////////////////////////
bool RBtree::remove(int key)
{
	Node* target = m_root;
	while (target != nullptr && key != target->key) {
		if (key < target->key) {
			target = target->left;
		} else {
			target = target->right;
		}
	}
	if (target == nullptr) return false;

	Node* victim = nullptr;
	Node* orphan = nullptr;
	if (target->left == nullptr) {
		victim = target;
		orphan = target->right;
	}
	else if (target->right == nullptr) {
		victim = target;
		orphan = target->left;
	} else {
		victim = target->right;
		while (victim->left != nullptr) {
			victim = victim->left;
		}
		orphan = victim->right;
	}

	Node* root = victim->parent();
	if (root == nullptr) { //此时victim==target
		if (orphan != nullptr) {
			orphan->parent(nullptr);
		}
		m_root = orphan;
		if (m_root != nullptr) {
			m_root->black = true;
		}
	} else {
		if (key < root->key) {
			root->hookLeft(orphan, nullptr);
		} else {
			root->hookRight(orphan, nullptr);
		}
		if (victim->black) { //红victim随便删，黑的要考虑
			if (orphan != nullptr && !orphan->black) {
				orphan->black = true; //红子变黑顶上
			} else {
				adjustAfterDelete(root, victim->key);
			}
		}
		target->key = victim->key;
	}
	m_pool.deallocate(victim);
	return true;
}
void RBtree::adjustAfterDelete(Node* G, int key)
{
	while(true) { //剩下情况：victim黑，orphan也黑，此时victim(orphan顶替)的兄弟必然存在
		Node* super = G->parent();
		if (key < G->key) {
			Node* U = G->right; //U != nullptr
			Node* L = U->left;
			Node* R = U->right;
			if (!U->black) { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G->hookRight(L);
				U->hookLeft(G);
				U->black = true;
				G->black = false;
				hookSubTree(super, U);
				continue; //变出黑U后再行解决
			} else {
				if (L == nullptr || L->black) {
					if (R == nullptr || R->black) { //双黑，变色解决
						U->black = false;
						if (G->black && super != nullptr) {
							G = super;
							continue; //上溯
						}
						G->black = true;
					} else { //中黑外红
						G->hookRight(L, nullptr);
						U->hookLeft(G);
						U->black = G->black;
						G->black = true;
						R->black = true;
						hookSubTree(super, U);
					}
				} else { //中红
					U->hookLeft(L->right, nullptr);
					G->hookRight(L->left, nullptr);
					L->hookRight(U);
					L->hookLeft(G);
					L->black = G->black;
					G->black = true;
					hookSubTree(super, L);
				}
			}
		} else {
			Node* U = G->left; //U != nullptr
			Node* R = U->right;
			Node* L = U->left;
			if (!U->black) { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G->hookLeft(R);
				U->hookRight(G);
				U->black = true;
				G->black = false;
				hookSubTree(super, U);
				continue; //变出黑U后再行解决
			} else {
				if (R == nullptr || R->black) {
					if (L == nullptr || L->black) { //双黑，变色解决
						U->black = false;
						if (G->black && super != nullptr) {
							G = super;
							continue; //上溯
						}
						G->black = true;
					} else { //中黑外红
						G->hookLeft(R, nullptr);
						U->hookRight(G);
						U->black = G->black;
						G->black = true;
						L->black = true;
						hookSubTree(super, U);
					}
				} else { //中红
					U->hookRight(R->left, nullptr);
					G->hookLeft(R->right, nullptr);
					R->hookLeft(U);
					R->hookRight(G);
					R->black = G->black;
					G->black = true;
					hookSubTree(super, R);
				}
			}
		}
		break; //个别情况需要循环
	}
}
//----------------红叔模式----------------
//|        bG        |        bU        |
//|       /  \       |       /  \       |
//|     bO    rU     |     rG    bR     |
//|          /  \    |    /  \          |
//|        bL    bR  |  bO    bL        |

//------------------双黑------------------
//|        xG        |        bG        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bO    rU     |
//|          /  \    |          /  \    |
//|        bL    bR  |        bL    bR  |

//----------------中黑外红----------------
//|        xG        |        xU        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bR     |
//|          /  \    |    /  \          |
//|        bL    rR  |  bO    bL        |

//------------------中红------------------
//|        xG        |        xL        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bU     |
//|          /  \    |    /  \  /  \    |
//|        rL    xR  |  bO   u  v   xR  |
//|       /  \       |                  |
//|      u    v      |                  |