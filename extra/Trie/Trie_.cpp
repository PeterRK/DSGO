#include <stdint.h>
#include <string>
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
const unsigned Trie::SIZE_LIMIT[9] = {
	1, 2, 4, 8, 16, 32, 64, 128, 256
};
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
