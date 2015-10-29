#include <stdint.h>
#include <cassert>
#include <string>
#include <algorithm>
#include "tricks.h"
#include "Trie.h"
////////////////////////////////////////////////////////////////////////////
#define ALLOC_IMP(lv) \
	Trie::Node::Pointer* Trie::AllocatePtrArray##lv(void) {		\
		return (Node::Pointer*)m_ptrPool##lv.allocate();		\
	}															\
	void Trie::DeallocatePtrArray##lv(Node::Pointer* ptr) {		\
		m_ptrPool##lv.deallocate((Node::Pointer(*)[1<<lv])ptr);	\
	}
ALLOC_IMP(1) ALLOC_IMP(2) ALLOC_IMP(3) ALLOC_IMP(4)
ALLOC_IMP(5) ALLOC_IMP(6) ALLOC_IMP(7) ALLOC_IMP(8)
#define FUNCX(lv) &Trie::AllocatePtrArray##lv
const Trie::AllocatePtrArray Trie::ALLOCATE[9] = {
	nullptr, FUNCX(1), FUNCX(2), FUNCX(3), FUNCX(4), FUNCX(5), FUNCX(6), FUNCX(7), FUNCX(8)
};
#undef FUNCX
#define FUNCX(lv) &Trie::DeallocatePtrArray##lv
const Trie::DeallocatePtrArray Trie::DEALLOCATE[9] = {
	nullptr, FUNCX(1), FUNCX(2), FUNCX(3), FUNCX(4), FUNCX(5), FUNCX(6), FUNCX(7), FUNCX(8)
};
#undef FUNCX
////////////////////////////////////////////////////////////////////////////
inline static uint16_t LEVEL_CAPACITY(uint8_t lv) {
	return 1U << lv;
}
////////////////////////////////////////////////////////////////////////////
inline Trie::Node* Trie::NewNode(void)
{
	Node* node = m_pool.allocate();
	node->cnt = 0;
	node->lv = 0;
	node->brs = 0;
	node->ref = 0;
	return node;
}
Trie::Trie(void)
{
	m_root = NewNode();
	m_root->ref = 1;		//空串总是有效
}
////////////////////////////////////////////////////////////////////////////
inline uint16_t Trie::SearchKid(Node* node, char ch)
{	//lv != 0
	assert(node->lv != 0);
	uint16_t start = 0, end = node->brs;
	while (start < end) {
		uint16_t mid = (start + end) / 2;
		if (ch > node->kids[mid]->key[0]) {
			start = mid + 1;
		} else {
			end = mid;
		}
	}
	return start;
}
uint16_t Trie::Search(const std::string& data)
{	//除了查找，还要尝试节缩
	Node* root = m_root;
	uint8_t mk = 0;
	for (unsigned idx = 0; idx < data.size(); idx++) {
		if (mk == root->cnt) { //下探
			if (root->brs == 0) return 0;
			if (root->brs == 1 && root->ref == 0 && root->cnt < Node::CAP) {
				Node* kid = (root->lv == 0) ? root->next : root->kids[0];
				if (root->cnt + kid->cnt > Node::CAP) { //半缩
					uint8_t j = 0;
					for (uint8_t i = root->cnt; i < Node::CAP; i++) {
						root->key[i] = kid->key[j++];
					}
					root->cnt = Node::CAP;
					uint8_t i = 0;
					while (j < kid->cnt) {
						kid->key[i++] = kid->key[j++];
					}
					kid->cnt = i;
				} else { //全缩
					for (uint8_t i = 0; i < kid->cnt; i++) {
						root->key[root->cnt++] = kid->key[i];
					}
					root->ref = kid->ref;
					root->brs = kid->brs;
					if (root->lv != 0) {
						(this->*DEALLOCATE[root->lv])(root->kids);
					}
					root->lv = kid->lv;
					root->kids = kid->kids;
					m_pool.deallocate(kid);
				}
			} else { //单纯下探
				if (root->lv == 0) {
					root = root->next;
					if (root->key[0] != data[idx]) return 0;
				} else {
					uint16_t spot = SearchKid(root, data[idx]);
					if (spot == root->brs
						|| root->kids[spot]->key[0] != data[idx]
					) return 0;
					root = root->kids[spot];
				}
				mk = 1;
				continue;
			}
		}
		if (data[idx] != root->key[mk++]) return 0;
	}
	if (mk != root->cnt) return 0;
	return root->ref;
}
////////////////////////////////////////////////////////////////////////////
Trie::Node* Trie::CreateTail(const std::string& data, unsigned begin)
{
	Node* head = NewNode();
	Node* last = head;
	for (; begin + Node::CAP < data.size(); begin += Node::CAP) {
		for (uint8_t i = 0; i < Node::CAP; i++) {
			last->key[i] = data[begin + i];
		}
		last->cnt = Node::CAP;
		last->brs = 1;
		last = last->next = NewNode();
	}
	while (begin < data.size()) {
		last->key[last->cnt++] = data[begin++];
	}
	last->ref = 1;
	return head;
}
void Trie::Split(Node* node, uint8_t mk)
{
	Node* kid = NewNode();
	for (uint8_t i = mk; i < node->cnt; i++) {
		kid->key[kid->cnt++] = node->key[i];
	}
	node->cnt = mk;
	kid->ref = node->ref;
	kid->brs = node->brs;
	kid->lv = node->lv;
	kid->kids = node->kids;
	node->ref = 0;
	node->brs = 1;
	node->lv = 0;
	node->next = kid;
}
inline void Trie::AddKid(Node* node, Node* peer)
{	//从1到2
	assert(node->lv == 0 && node->brs == 1);
	Node* kid = node->next;
	if (kid->key[0] > peer->key[0]) {
		std::swap(kid, peer);
	}
	node->brs = 2;
	node->lv = 1;
	node->kids = (this->*ALLOCATE[1])();
	node->kids[0] = kid;
	node->kids[1] = peer;
}
void Trie::Insert(const std::string& data)
{
	if (data.empty()) return;
	Node* root = m_root;
	uint8_t mk = 0;
	for (unsigned idx = 0; idx < data.size(); idx++) {
		if (mk == root->cnt) {	//下探
			if (root->lv == 0) {
				if (root->brs == 0) {
					root->brs = 1;
					root->next = CreateTail(data, idx);
					return;
				}
				if (root->next->key[0] != data[idx]) {
					AddKid(root, CreateTail(data, idx));
					return;
				}
				root = root->next;
			} else {
				uint16_t spot = SearchKid(root, data[idx]);
				if (spot == root->brs
					|| root->kids[spot]->key[0] != data[idx]
				){	//扩展
					if (root->brs == LEVEL_CAPACITY(root->lv)) {
						Node::Pointer* old = root->kids;
						auto deallocate = DEALLOCATE[root->lv];
						root->kids = (this->*ALLOCATE[++root->lv])();
						for (uint16_t i = 0; i < spot; i++) {
							root->kids[i] = old[i];
						}
						for (uint16_t i = root->brs; i > spot; i--) {
							root->kids[i] = old[i - 1];
						}
						(this->*deallocate)(old);
					} else {
						for (uint16_t i = root->brs; i > spot; i--) {
							root->kids[i] = root->kids[i - 1];
						}
					}
					root->kids[spot] = CreateTail(data, idx);
					root->brs++;
					return;
				}
				root = root->kids[spot];
			}
			mk = 1;
		} else {
			if (data[idx] != root->key[mk]) {
				Split(root, mk);
				AddKid(root, CreateTail(data, idx));
				return;
			}
			mk++;
		}
	}
	if (mk != root->cnt) {
		Split(root, mk);
	}
	root->ref++;
}	
////////////////////////////////////////////////////////////////////////////
void Trie::Remove(const std::string& data, bool all)
{
	if (data.empty()) return;
	Node* knot = nullptr;
	uint16_t branch = 0;	//记录独苗分支节点
	Node* root = m_root;
	uint8_t mk = 0;
	for (unsigned idx = 0; idx < data.size(); idx++) {
		if (mk == root->cnt) {	//下探
			if (root->brs == 0) return;
			if (root->lv == 0) {
				if (root->ref != 0) knot = root;
				root = root->next;
				if (root->key[0] != data[idx]) return;
			} else {
				uint16_t spot = SearchKid(root, data[idx]);
				if (spot == root->brs
					|| root->kids[spot]->key[0] != data[idx]
				) return;
				if (root->ref != 0 || root->brs > 1) {
					knot = root;
					branch = spot;
				}
				root = root->kids[spot];
			}
			mk = 1;
		} else {
			if (data[idx] != root->key[mk]) return;
			mk++;
		}
	}
	if (mk != root->cnt || root->ref == 0) return;
	if (all) {
		root->ref = 0;
	} else {
		root->ref--;
	}
	if (root->ref == 0 && root->brs == 0) {	//删除独苗
		knot->brs--;
		Node* tail = nullptr;
		if (knot->lv == 0) {
			tail = knot->next;
		} else {
			tail = knot->kids[branch];
			if (knot->brs < LEVEL_CAPACITY(knot->lv) / 4) {	//错级收缩，抗抖动
				Node::Pointer* old = knot->kids;
				auto deallocate = DEALLOCATE[knot->lv];
				knot->kids = (this->*ALLOCATE[--knot->lv])();
				for (uint16_t i = 0; i < branch; i++) {
					knot->kids[i] = old[i];
				}
				for (uint16_t i = branch; i < knot->brs; i++) {
					knot->kids[i] = old[i + 1];
				}
				(this->*deallocate)(old);
			} else {
				for (uint16_t i = branch; i < knot->brs; i++) {
					knot->kids[i] = knot->kids[i + 1];
				}
			}
		}
		for (bool lock = true; lock; ) {
			if (tail == root) lock = false;
			Node* tmp = tail;
			if (tail->lv == 0) {
				tail = tail->next;
			} else {
				tail = tail->kids[0];
				(this->*DEALLOCATE[tmp->lv])(tmp->kids);
			}
			m_pool.deallocate(tmp);
		}
	}
}
////////////////////////////////////////////////////////////////////////////