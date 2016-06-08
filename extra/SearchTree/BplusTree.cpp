#include <cstdint>
#include <cstddef>
#include <new>
#include <vector>
#include "tricks.h"
#include "BplusTree.h"
////////////////////////////////////////////////////////////////////////////
int BplusTree::Node::locate(int key)
const {
	int start = 0;
	int end = cnt - 1;
	while (start < end) {
		int mid = (start + end) / 2;
		if (key > data[mid]) {
			start = mid + 1;
		} else {
			end = mid;
		}
	} //寻找第一个大于或等于key的位置
	return start;
}
////////////////////////////////////////////////////////////////////////////
bool BplusTree::search(int key)
const {
	if (m_root == nullptr || key > m_root->ceil()) {
		return false;
	}
	Index* target = m_root;
	while (target->inner) {
		int idx = target->locate(key);
		if (key == target->data[idx]) {
			return true;
		}
		target = target->kids[idx];
	}
	return key == target->data[target->locate(key)];
}
void BplusTree::travel(Func func) 
const {
	for (auto unit = m_head; unit != nullptr; unit = unit->next) {
		for (int i  = 0; i < unit->cnt; i++) {
			func(unit->data[i]);
		}
	}
}
////////////////////////////////////////////////////////////////////////////
inline BplusTree::Index* BplusTree::Index::NewNode(Allocator<Index>& pool)
{
	Index* unit = pool.allocate();
	unit->inner = true;
	return unit;
}
inline void BplusTree::Index::remove(int place)
{
	cnt--;
	for (int i = place; i < cnt; i++) {
		data[i] = data[i + 1];
		kids[i] = kids[i + 1];
	}
}
BplusTree::Index* BplusTree::Index::insert(
	int place, Index* kid, Allocator<Index>& pool
){	//peer为分裂项，peer为nullptr时表示不分裂
	if (cnt < FULL_SZ) {
		for (int i = cnt; i > place; i--) {
			data[i] = data[i - 1];
			kids[i] = kids[i - 1];
		}
		data[place] = kid->ceil();
		kids[place] = kid;
		cnt++;
		return nullptr;
	}

	Index* peer = NewNode(pool);
	peer->cnt = cnt = HALF_SZ;
	if (place < HALF_SZ) {
		for (int i = 0; i < HALF_SZ; i++) {
			peer->data[i] = data[i + (HALF_SZ - 1)];
			peer->kids[i] = kids[i + (HALF_SZ - 1)];
		}
		for (int i = HALF_SZ - 1; i > place; i--) {
			data[i] = data[i - 1];
			kids[i] = kids[i - 1];
		}
		data[place] = kid->ceil();
		kids[place] = kid;
	}
	else {
		for (int i = FULL_SZ; i > place; i--) {
			peer->data[i - HALF_SZ] = data[i - 1];
			peer->kids[i - HALF_SZ] = kids[i - 1];
		}
		peer->data[place - HALF_SZ] = kid->ceil();
		peer->kids[place - HALF_SZ] = kid;
		for (int i = HALF_SZ; i < place; i++) {
			peer->data[i - HALF_SZ] = data[i];
			peer->kids[i - HALF_SZ] = kids[i];
		}
	}
	return peer;
}
bool BplusTree::Index::combine(
	Index* peer, Allocator<Index>& pool
){	//要求peer为unit后的节点，发生合并返回true
	int total = cnt + peer->cnt;
	if (total <= HALF_SZ + QUARTER_SZ) {
		for (int i = cnt; i < total; i++) {
			data[i] = peer->data[i - cnt];
			kids[i] = peer->kids[i - cnt];
		}
		cnt = total;
		pool.deallocate(peer);
		return true;
	}
	//分流而不合并
	int unit_sz = total / 2;
	if (cnt == unit_sz) return false;
	int peer_sz = total - unit_sz;
	if (peer->cnt > peer_sz) {
		int diff = peer->cnt - peer_sz;
		for (int i = cnt; i < unit_sz; i++) {
			data[i] = peer->data[i - cnt];
			kids[i] = peer->kids[i - cnt];
		}
		for (int i = 0; i < peer_sz; i++) {
			peer->data[i] = peer->data[i + diff];
			peer->kids[i] = peer->kids[i + diff];
		}
	} else {
		int diff = peer_sz - peer->cnt;
		for (int i = peer->cnt - 1; i >= 0; i--) {
			peer->data[i + diff] = peer->data[i];
			peer->kids[i + diff] = peer->kids[i];
		}
		for (int i = cnt - 1; i >= unit_sz; i--) {
			peer->data[i - unit_sz] = data[i];
			peer->kids[i - unit_sz] = kids[i];
		}
	}
	cnt = unit_sz;
	peer->cnt = peer_sz;
	return false;
}
////////////////////////////////////////////////////////////////////////////
inline BplusTree::Leaf* BplusTree::Leaf::NewNode(Allocator<Leaf>& pool)
{
	Leaf* unit = pool.allocate();
	unit->inner = false;
	return unit;
}
inline void BplusTree::Leaf::remove(int place)
{
	cnt--;
	for (int i = place; i < cnt; i++) {
		data[i] = data[i + 1];
	}
}
BplusTree::Leaf* BplusTree::Leaf::insert(
	int place, int key, Allocator<Leaf>& pool
){	//peer为分裂项，peer为nullptr时表示不分裂
	if (cnt < FULL_SZ) {
		for (int i = cnt; i > place; i--) {
			data[i] = data[i - 1];
		}
		data[place] = key;
		cnt++;
		return nullptr;
	}

	Leaf* peer = NewNode(pool);
	peer->next = next;
	next = peer;
	peer->cnt = cnt = HALF_SZ;
	if (place < HALF_SZ) {
		for (int i = 0; i < HALF_SZ; i++) {
			peer->data[i] = data[i + (HALF_SZ - 1)];
		}
		for (int i = HALF_SZ - 1; i > place; i--) {
			data[i] = data[i - 1];
		}
		data[place] = key;
	} else {
		for (int i = FULL_SZ; i > place; i--) {
			peer->data[i - HALF_SZ] = data[i - 1];
		}
		peer->data[place - HALF_SZ] = key;
		for (int i = HALF_SZ; i < place; i++) {
			peer->data[i - HALF_SZ] = data[i];
		}
	}
	return peer;
}
bool BplusTree::Leaf::combine(
	Leaf* peer, Allocator<Leaf>& pool
){	//要求peer为unit后的节点，发生合并返回true
	int total = cnt + peer->cnt;
	if (total <= HALF_SZ + QUARTER_SZ) {
		for (int i = cnt; i < total; i++) {
			data[i] = peer->data[i - cnt];
		}
		cnt = total;
		next = peer->next;
		pool.deallocate(peer);
		return true;
	}
	//分流而不合并
	int unit_sz = total / 2;
	if (cnt == unit_sz) return false;
	int peer_sz = total - unit_sz;
	if (peer->cnt > peer_sz) {
		int diff = peer->cnt - peer_sz;
		for (int i = cnt; i < unit_sz; i++) {
			data[i] = peer->data[i - cnt];
		}
		for (int i = 0; i < peer_sz; i++) {
			peer->data[i] = peer->data[i + diff];
		}
	} else {
		int diff = peer_sz - peer->cnt;
		for (int i = peer->cnt - 1; i >= 0; i--) {
			peer->data[i + diff] = peer->data[i];
		}
		for (int i = cnt - 1; i >= unit_sz; i--) {
			peer->data[i - unit_sz] = data[i];
		}
	}
	cnt = unit_sz;
	peer->cnt = peer_sz;
	return false;
}
////////////////////////////////////////////////////////////////////////////
bool BplusTree::insert(int key)
{
	if (m_root == nullptr) {
		m_head = Leaf::NewNode(m_lpool);
		m_head->cnt = 1;
		m_head->data[0] = key;
		m_head->next = nullptr;
		m_root = reinterpret_cast<Index*>(m_head);
		return true;
	}

	m_pstack.clear();
	m_nstack.clear();

	Leaf* leaf;
	int place;
	Index* target = m_root;
	if (key > m_root->ceil()) { //右界拓展
		while (target->inner) {
			int idx = target->cnt - 1;
			target->data[idx] = key; //之后难以修改，现在先改掉
			m_pstack.push_back(target);
			m_nstack.push_back(idx);
			target = target->kids[idx];
		}
		leaf = reinterpret_cast<Leaf*>(target);
		place = leaf->cnt;
	} else {
		while (target->inner) {
			int idx = target->locate(key);
			if (key == target->data[idx]) return false;
			m_pstack.push_back(target);
			m_nstack.push_back(idx);
			target = target->kids[idx];
		}
		leaf = reinterpret_cast<Leaf*>(target);
		place = leaf->locate(key);
		if (key == leaf->data[place]) return false;
	}

	Index* peer = reinterpret_cast<Index*>(leaf->insert(place, key, m_lpool));
	while (peer != nullptr) {
		if (m_pstack.empty()) {
			Index* unit = Index::NewNode(m_ipool);
			unit->cnt = 2;
			unit->data[0] = target->ceil();
			unit->data[1] = peer->ceil();
			unit->kids[0] = target;
			unit->kids[1] = peer;
			m_root = unit;
			break;
		} else {
			Index* parent = m_pstack.back(); m_pstack.pop_back();
			int idx = m_nstack.back(); m_nstack.pop_back();
			parent->data[idx] = target->ceil();
			target = parent;
			peer = target->insert(idx + 1, peer, m_ipool);
		}
	}
	return true;
}
////////////////////////////////////////////////////////////////////////////
bool BplusTree::remove(int key)
{
	if (m_root == nullptr || key > m_root->ceil()) return false;
	m_pstack.clear();
	m_nstack.clear();

	Index* target = m_root;
	while (target->inner) {
		int idx = target->locate(key);
		m_pstack.push_back(target);
		m_nstack.push_back(idx);
		target = target->kids[idx];
	}
	Leaf* leaf = reinterpret_cast<Leaf*>(target);
	int place = leaf->locate(key);
	if (key != leaf->data[place]) return false;

	leaf->remove(place);
	if (m_pstack.empty()) {
		if (leaf->cnt == 0) {
			m_root = nullptr;
			m_head = nullptr;
			m_lpool.deallocate(leaf);
		}
		return true;
	} //除了到根节点，unit->cnt >= 2
	bool shrink = (place == leaf->cnt);
	int new_ceil = leaf->ceil();

	Index* parent = m_pstack.back(); m_pstack.pop_back();
	place = m_nstack.back(); m_nstack.pop_back();
	if (leaf->cnt <= Leaf::QUARTER_SZ) {
		Leaf* peer = leaf;
		if (place == parent->cnt - 1) {
			leaf = reinterpret_cast<Leaf*>(parent->kids[place - 1]);
		} else {
			peer = reinterpret_cast<Leaf*>(parent->kids[++place]);
			shrink = false;
		}
		bool combined = leaf->combine(peer, m_lpool);
		parent->data[place - 1] = leaf->ceil();

		while (combined) {
			Index* index = parent; //此后代码与之前类似，但unit的类型已经不同
			index->remove(place);
			if (m_pstack.empty()) {
				if (index->cnt == 1) {
					m_root = index->kids[0];
					m_ipool.deallocate(index);
				}
				return true;
			}

			parent = m_pstack.back(); m_pstack.pop_back();
			place = m_nstack.back(); m_nstack.pop_back();
			if (index->cnt <= Index::QUARTER_SZ) {
				Index* peer = index;
				if (place == parent->cnt - 1) {
					index = parent->kids[place - 1];
				} else {
					peer = parent->kids[++place];
					shrink = false;
				}
				combined = index->combine(peer, m_ipool);
				parent->data[place - 1] = index->ceil();
				continue;
			}
			break;
		}
	}
	if (shrink) { //
		parent->data[place] = new_ceil;
		while (place == parent->cnt - 1 && !m_pstack.empty()) {
			parent = m_pstack.back(); m_pstack.pop_back();
			place = m_nstack.back(); m_nstack.pop_back();
			parent->data[place] = new_ceil;
		}
	}
	return true;
}
////////////////////////////////////////////////////////////////////////////