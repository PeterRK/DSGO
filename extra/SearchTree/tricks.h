#pragma once

class NonCopyable {
protected:
	NonCopyable(void) {}
	~NonCopyable(void) {}
private:
	NonCopyable(const NonCopyable&) = delete;
	const NonCopyable& operator=(const NonCopyable&) = delete;
};

template<typename T, unsigned N = 255>
class Allocator final : NonCopyable {
	static_assert(N >= 7, "N is not big enough!");
	static_assert(N <= 4095, "N is too big!");
private:
	union Obj {
		char space[sizeof(T)];
		Obj* next;
	};
	struct Block {
		Block* next;
	};
	struct DataBlock : public Block {
		Obj data[N];
	};
	struct HookBlock : public Block {
		Obj* hook[N];
	};
	static Block* NewBlock(void) {
		if (isHookMode) {
			return reinterpret_cast<HookBlock*>(operator new(sizeof(HookBlock)));
		} else {
			return reinterpret_cast<DataBlock*>(operator new(sizeof(DataBlock)));
		}
	}
	static void DeleteBlock(Block* block, unsigned cnt) {
		if (isHookMode) {
			Obj** hook = reinterpret_cast<HookBlock*>(block)->hook;
			for (unsigned i = 0; i < cnt; i++) {
				operator delete(hook[i]);
			}
		}
		operator delete(block);
	}

	unsigned	m_spot;
	Block*		m_crrent;
	Block*		m_head;
	Obj*		m_free;
	unsigned	m_balance;

public:
	static constexpr bool isHookMode = sizeof(T) > 128;
	unsigned balance(void) const { return m_balance; }

	Allocator(void)
		: m_spot(0), m_free(nullptr), m_balance(0)
	{
		m_head = m_crrent = NewBlock();
		m_head->next = nullptr;
	}
	~Allocator(void) {
		DeleteBlock(m_crrent, m_spot);
		while (m_head != m_crrent) {
			Block* tmp = m_head;
			m_head = m_head->next;
			DeleteBlock(tmp, N);
		}
	}

	T* allocate(void) {
		Obj* obj = m_free;
		if (obj != nullptr) {
			m_free = m_free->next;
		} else {
			if (m_spot == N) {
				m_spot = 0;
				m_crrent = m_crrent->next = NewBlock();
				m_crrent->next = nullptr;
			}
			if (isHookMode) {
				HookBlock* block = reinterpret_cast<HookBlock*>(m_crrent);
				obj = block->hook[m_spot++] = 
					reinterpret_cast<Obj*>(operator new(sizeof(Obj)));
			} else {
				DataBlock* block = reinterpret_cast<DataBlock*>(m_crrent);
				obj = &block->data[m_spot++];
			}
		}
		m_balance++;
		return reinterpret_cast<T*>(obj);
	}
	void deallocate(T* p) {
		Obj* obj = reinterpret_cast<Obj*>(p);
		obj->next = m_free;
		m_free = obj;
		m_balance--;
	}
};