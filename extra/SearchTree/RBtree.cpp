#include <cstdint>
#include <cstddef>
#include <new>
#include "tricks.h"
#include "RBtree.h"
////////////////////////////////////////////////////////////////////////////
bool RBtree::search(int key)
const {
	Node* root = m_root;
	while (root != NULL) {
		if (key == root->key) {
			return true;
		}
		root = (key < root->key) ? root->left : root->right;
	}
	return false;
}
////////////////////////////////////////////////////////////////////////////
inline void RBtree::hookSubTree(Node* super, Node* root) {
	if (super == NULL) {
		m_root = super->hook(root);
	} else {
		if (root->key < super->key) {
			super->left = super->hook(root);
		} else {
			super->right = super->hook(root);
		}
	}
}
inline RBtree::Node* RBtree::newNode(RBtree::Node* parent, int key)
{
	Node* node = m_pool.allocate();
	node->key = key;
	node->black = false;
	node->parent = reinterpret_cast<uintptr_t>(parent);
	node->left = NULL;
	node->right = NULL;
	return node;
}
////////////////////////////////////////////////////////////////////////////
bool RBtree::insert(int key)
{
	if (m_root == NULL) {
		m_root = newNode(NULL, key);
		m_root->black = true;
		return true;
	}

	Node* root = m_root;
	while(true) {
		if (key < root->key) {
			if (root->left == NULL) {
				root->left = newNode(root, key);
				break;
			}
			root = root->left;
		} else if (key > root->key) {
			if (root->right == NULL) {
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
	while (!P->black) { //违法双红禁
		Node* G = reinterpret_cast<Node*>(P->parent); //必然存在，根为黑，P非根
		Node* super = reinterpret_cast<Node*>(G->parent);
		if (key < G->key) {
			Node* U = G->right;
			if (U != NULL && !U->black) { //红叔模式，变色解决
				P->black = true;
				U->black = true;
				if (super != NULL) {
					G->black = false;
					P = reinterpret_cast<Node*>(G->parent);
					continue; //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if (key < P->key) { //LL
					G->left = G->tryHook(P->right);
					P->right = P->hook(G);
					G->black = false;
					P->black = true;
					hookSubTree(super, P);
				} else { //LR
					Node* C = P->right;
					P->right = P->tryHook(C->left);
					G->left = G->tryHook(C->right);
					C->left = C->hook(P);
					C->right = C->hook(G);
					G->black = false;
					C->black = true;
					hookSubTree(super, C);
				}
			}
		} else {
			Node* U = G->left;
			if (U != NULL && !U->black) { //红叔模式，变色解决
				P->black = true;
				U->black = true;
				if (super != NULL) {
					G->black = false;
					P = reinterpret_cast<Node*>(G->parent);
					continue; //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if (key > P->key) { //RR
					G->right = G->tryHook(P->left);
					P->left = P->hook(G);
					G->black = false;
					P->black = true;
					hookSubTree(super, P);
				} else { //RL
					Node* C = P->left;
					P->left = P->tryHook(C->right);
					G->right = G->tryHook(C->left);
					C->right = C->hook(P);
					C->left = C->hook(G);
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
	while (target != NULL && key != target->key) {
		if (key < target->key) {
			target = target->left;
		} else {
			target = target->right;
		}
	}
	if (target == NULL) return false;

	Node* victim = NULL;
	Node* orphan = NULL;
	if (target->left == NULL) {
		victim = target;
		orphan = target->right;
	}
	else if (target->right == NULL) {
		victim = target;
		orphan = target->left;
	} else {
		victim = target->right;
		while (victim->left != NULL) {
			victim = victim->left;
		}
		orphan = victim->right;
	}

	Node* root = reinterpret_cast<Node*>(victim->parent);
	if (root == NULL) { //此时victim==target
		m_root = root->tryHook(orphan);
		if (m_root != NULL) {
			m_root->black = true;
		}
	} else {
		if (key < root->key) {
			root->left = root->tryHook(orphan);
		} else {
			root->right = root->tryHook(orphan);
		}
		if (victim->black) { //红victim随便删，黑的要考虑
			if (orphan != NULL && !orphan->black) {
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
		Node* super = reinterpret_cast<Node*>(G->parent);
		if (key < G->key) {
			Node* U = G->right; //U != NULL
			Node* L = U->left;
			Node* R = U->right;
			if (!U->black) { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G->right = G->hook(L);
				U->left = U->hook(G);
				U->black = true;
				G->black = false;
				hookSubTree(super, U);
				continue; //变出黑U后再行解决
			} else {
				if (L == NULL || L->black) {
					if (R == NULL || R->black) { //双黑，变色解决
						U->black = false;
						if (G->black && super != NULL) {
							G = super;
							continue; //上溯
						}
						G->black = true;
					} else { //中黑外红
						G->right = G->tryHook(L);
						U->left = U->hook(G);
						U->black = G->black;
						G->black = true;
						R->black = true;
						hookSubTree(super, U);
					}
				} else { //中红
					U->left = U->tryHook(L->right);
					G->right = G->tryHook(L->left);
					L->right = L->hook(U);
					L->left = L->hook(G);
					L->black = G->black;
					G->black = true;
					hookSubTree(super, L);
				}
			}
		} else {
			Node* U = G->left; //U != NULL
			Node* R = U->right;
			Node* L = U->left;
			if (!U->black) { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G->left = G->hook(R);
				U->right = U->hook(G);
				U->black = true;
				G->black = false;
				hookSubTree(super, U);
				continue; //变出黑U后再行解决
			} else {
				if (R == NULL || R->black) {
					if (L == NULL || L->black) { //双黑，变色解决
						U->black = false;
						if (G->black && super != NULL) {
							G = super;
							continue; //上溯
						}
						G->black = true;
					} else { //中黑外红
						G->left = G->tryHook(R);
						U->right = U->hook(G);
						U->black = G->black;
						G->black = true;
						L->black = true;
						hookSubTree(super, U);
					}
				} else { //中红
					U->right = U->tryHook(R->left);
					G->left = G->tryHook(R->right);
					R->left = R->hook(U);
					R->right = R->hook(G);
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
