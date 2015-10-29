#pragma once

class Trie : NonCopyable {
private:
	struct Node {
		static constexpr uint8_t CAP = sizeof(void*);
		typedef Node* Pointer;
		char key[CAP];
		uint8_t cnt;
		uint8_t lv;
		uint16_t brs;
		uint32_t ref;
		union {
			Pointer next;
			Pointer* kids;
		};
	};
	Allocator<Node> m_pool;

	typedef Node::Pointer* (Trie::*AllocatePtrArray)(void);
	typedef void (Trie::*DeallocatePtrArray)(Node::Pointer* ptr);

#define ALLOCATEOR(lv, sz) \
	Allocator<Node::Pointer[1<<lv], sz> m_ptrPool##lv;	\
	Node::Pointer* AllocatePtrArray##lv(void);			\
	void DeallocatePtrArray##lv(Node::Pointer* ptr);
	ALLOCATEOR(1, 255) ALLOCATEOR(2, 255) ALLOCATEOR(3, 127) ALLOCATEOR(4, 127)
	ALLOCATEOR(5, 127) ALLOCATEOR(6, 127) ALLOCATEOR(7,  63) ALLOCATEOR(8,  31)
#undef ALLOCATEOR

	static const AllocatePtrArray ALLOCATE[9];
	static const DeallocatePtrArray DEALLOCATE[9];

	Node* m_root;

	Node* NewNode(void);
	static uint16_t SearchKid(Node* node, char ch);
	Node* CreateTail(const std::string& data, unsigned begin);
	void Split(Node* node, uint8_t mk);
	void AddKid(Node* node, Node* peer);

public:
	Trie(void);

	uint16_t Search(const std::string& data);
	void Insert(const std::string& data);
	void Remove(const std::string& data, bool all=false);
};
