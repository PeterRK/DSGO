#include <cstddef>
#include <cstdint>
#include <cassert>
#include <cstring>
#include <ctime>
#include <new>
#include <algorithm>
#include "SkipList.h"
////////////////////////////////////////////////////////////////////////////
SkipList::SkipList(void)
{
	m_rand = time(nullptr);
	m_cnt = 1;
	m_ceil = FACTOR;
	m_floor = 1;
	m_level = 1;
	m_cap = 10;
	m_heads = new Node*[m_cap];
	m_knots = new Node*[m_cap];
	m_heads[0] = nullptr;
}
SkipList::~SkipList(void) {
	Node* node = m_heads[0];
	while (node != nullptr) {
		Node* tmp = node;
		node = node->next[0];
		delete tmp;
	}
	delete[] m_heads;
	delete[] m_knots;
}
inline SkipList::Node* SkipList::shadow(void) const {
	uintptr_t off = (uintptr_t)(&((Node*)0)->next);
	return (Node*)((uintptr_t)m_heads - off);
}
////////////////////////////////////////////////////////////////////////////
unsigned SkipList::size(void)
const{
	return m_cnt - 1;
}
bool SkipList::isEmpty(void)
const{
	return size() == 0;
}
bool SkipList::search(int key)
const{
	Node* knot = shadow();
	for (int i  = m_level - 1; i >= 0; i--) {
		while (knot->next[i] != nullptr && knot->next[i]->key < key) {
			knot = knot->next[i];
		}
	}
	Node* target = knot->next[0];
	return target != nullptr && target->key == key;
}
void SkipList::travel(Func func)
const{
	for (Node* pt = m_heads[0]; pt != nullptr; pt = pt->next[0]) {
		func(pt->key);
	}
}
////////////////////////////////////////////////////////////////////////////
bool SkipList::insert(int key) 
{
	if (m_level == m_cap) {
		m_cap += 5;
		Node** old = m_knots;
		m_knots = new Node*[m_cap];
		delete old;
		old = m_heads;
		m_heads = new Node*[m_cap];
		memcpy(m_heads, old, sizeof(Node*)*m_level);
		delete old;
	}

	Node* knot = shadow();
	for (int i = m_level - 1; i >= 0; i--) {
		while (knot->next[i] != nullptr && knot->next[i]->key < key) {
			knot = knot->next[i];
		}
		m_knots[i] = knot;
	}
	Node* target = knot->next[0];
	if (target != nullptr && target->key == key) return false;

	if (++m_cnt == m_ceil) {
		m_floor = m_ceil;
		m_ceil *= FACTOR;
		m_heads[m_level] = nullptr;
		m_knots[m_level] = shadow();
		m_level++;
	}

	int lv = 1;
	while (lv < m_level &&
		(m_rand = m_rand * 1103515245 + 12345) <= (((unsigned)~0) / FACTOR)
	) lv++;

	target = (Node*)operator new(sizeof(Node) + (lv-1)*sizeof(Node*));
	target->key = key;
	target->size = lv;
	for (int i  = 0; i < lv; i++) {
		target->next[i] = m_knots[i]->next[i];
		m_knots[i]->next[i] = target;
	}
	return true;
}
////////////////////////////////////////////////////////////////////////////
bool SkipList::remove(int key) 
{
	Node* knot = shadow();
	for (int i = m_level - 1; i >= 0; i--) {
		while (knot->next[i] != nullptr && knot->next[i]->key < key) {
			knot = knot->next[i];
		}
		m_knots[i] = knot;
	}
	Node* target = knot->next[0];
	if (target == nullptr || target->key != key) return false;

	int lv = std::min(target->size, m_level);
	for (int i = 0; i < lv; i++) {
		m_knots[i]->next[i] = target->next[i];
	}
	delete target;

	if (m_cnt-- == m_floor) {
		m_ceil = m_floor;
		m_floor /= FACTOR;
		m_level--;
	}
	return true;
}
////////////////////////////////////////////////////////////////////////////
