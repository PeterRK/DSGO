#include <cstdint>
#include <cstddef>
#include <new>
#include "tricks.h"
#include "AVLtree.h"
////////////////////////////////////////////////////////////////////////////
bool AVLtree::search(int key)
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
inline AVLtree::Node* AVLtree::newNode(AVLtree::Node* parent, int key)
{
	Node* node = m_pool.allocate();
	node->key = key;
	node->state = 0;
	node->parent = parent;
	node->left = NULL;
	node->right = NULL;
	return node;
}
////////////////////////////////////////////////////////////////////////////
//--------------LR形式--------------
//|       G       |       C       |
//|      / \      |      / \      |
//|     P         |     P   G     |
//|    / \        |    / \ / \    |
//|       C       |      u v      |
//|      / \      |               |
//|     u   v     |               |

//--------------LL形式--------------
//|       G       |       P       |
//|      / \      |      / \      |
//|     P         |     C   G     |
//|    / \        |    / \ / \    |
//|   C   x       |        x      |
//|  / \          |               |
//|               |               |
AVLtree::Node* AVLtree::Rotate(Node* G, bool& stop)
{
	//assert(G->state == 2 || G->state == -2);
	stop = false;
	Node* root = NULL;
	if (G->state == 2) { //左倾右旋
		Node* P = G->left;
		if (P->state == -1) { //LR
			Node* C = P->right; //一定非NULL
			P->right = P->tryHook(C->left);
			G->left = G->tryHook(C->right);
			C->left = C->hook(P);
			C->right = C->hook(G);
			switch (C->state) {
			case 1:
				G->state = -1;
				P->state = 0;
				break;
			case -1:
				G->state = 0;
				P->state = 1;
				break;
			default:
				G->state = 0;
				P->state = 0;
				break;
			}
			C->state = 0;
			root = C;
		} else { //LL
			G->left = G->tryHook(P->right);
			P->right = P->hook(G);
			if (P->state == 0) { //不降高旋转
				G->state = 1;
				P->state = -1;
				stop = true;
			} else { //P->state == 1
				G->state = 0;
				P->state = 0;
			}
			root = P;
		}
	}
	else { //右倾左旋(P->state==-2)
		Node* P = G->right;
		if (P->state == 1) { //RL
			Node* C = P->left; //一定非NULL
			P->left = P->tryHook(C->right);
			G->right = G->tryHook(C->left);
			C->right = C->hook(P);
			C->left = C->hook(G);
			switch (C->state) {
				case -1:
					G->state = 1;
					P->state = 0;
					break;
				case 1:
					G->state = 0;
					P->state = -1;
					break;
				default:
					G->state = 0;
					P->state = 0;
					break;
			}
			C->state = 0;
			root = C;
		} else { //RR
			G->right = G->tryHook(P->left);
			P->left = P->hook(G);
			if (P->state == 0) { //不降高旋转
				G->state = -1;
				P->state = 1;
				stop = true;
			} else { //P->state == -1
				G->state = 0;
				P->state = 0;
			}
			root = P;
		}
	}
	return root;
}
////////////////////////////////////////////////////////////////////////////
bool AVLtree::insert(int key)
{
	if (m_root == NULL) {
		m_root = newNode(NULL, key);
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

	while(true) {
		int8_t state = root->state;
		root->state += (key < root->key)? 1 : -1;
		if (state == 0 && root->parent != NULL) {
			root = root->parent;
			continue;
		}
		if (state != 0 && root->state != 0) { //2 || -2
			Node* super = root->parent;
			bool stop;
			root = Rotate(root, stop);
			if (super == NULL) {
				m_root = super->hook(root);
			} else {
				if (key < super->key) {
					super->left = super->hook(root);
				} else {
					super->right = super->hook(root);
				}
			}
		}
		break;
	}
	return true;
}
////////////////////////////////////////////////////////////////////////////
bool AVLtree::remove(int key)
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
	} else if (target->right == NULL) {
		victim = target;
		orphan = target->left;
	} else {
		if (target->state == 1) {
			victim = target->left;
			while (victim->right != NULL) {
				victim = victim->right;
			}
			orphan = victim->left;
		} else {
			victim = target->right;
			while (victim->left != NULL) {
				victim = victim->left;
			}
			orphan = victim->right;
		}
	}

	Node* root = victim->parent;
	if (root == NULL) { //此时victim==target
		m_root = root->tryHook(orphan);
	} else {
		key = victim->key;
		int8_t state = root->state;
		if (key < root->key) {
			root->left = root->tryHook(orphan);
			root->state--;
		} else {
			root->right = root->tryHook(orphan);
			root->state++;
		}

		while (state != 0) { //如果原平衡因子为0则子树高度不变
			bool stop;
			Node* super = root->parent;
			if (super == NULL) {
				if (root->state != 0) { //2 || -2
					root = Rotate(root, stop);
					m_root = super->hook(root);
				}
				break;
			} else {
				if (root->state != 0) { //2 || -2
					root = Rotate(root, stop);
					if (key < super->key) {
						super->left = super->hook(root);
					} else {
						super->right = super->hook(root);
					}
					if (stop) break;
				}
				root = super;
				state = root->state;
				root->state -= (key < root->key) ? 1 : -1;
			}
		}
		target->key = key;
	}
	m_pool.deallocate(victim);
	return true;
}
////////////////////////////////////////////////////////////////////////////


