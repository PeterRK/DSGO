#pragma once

class NonCopyable {
protected:
	NonCopyable(void) {}
	~NonCopyable(void) {}
private:
	NonCopyable(const NonCopyable&) = delete;
	const NonCopyable& operator=(const NonCopyable&) = delete;
};

template<typename T, unsigned N=1024>
class Allocator : NonCopyable {
private:
	static const unsigned SIZE = N-1;
	union Obj {
		char	data[sizeof(T)];
		Obj*	next;
	};
	struct Block {
		Obj		data[SIZE];
		Block*	next;
	};

	unsigned	m_spot;
	Block*		m_crrent;
	Block*		m_head;
	Obj*		m_free;
	unsigned	m_balance;

public:
	Allocator(void)
		: m_spot(0), m_crrent(new Block), m_head(m_crrent),
		m_free(nullptr), m_balance(0)
	{ m_crrent->next = nullptr; }
	~Allocator(void) {
		while (m_head != nullptr) {
			m_crrent = m_head;
			m_head = m_head->next;
			operator delete(m_crrent);
		}
	}
	unsigned balance(void) const { return m_balance; }

	T* allocate(void) {
		Obj* obj = m_free;
		if (obj != nullptr) {
			m_free = m_free->next;
		} else {
			if (m_spot == SIZE) {
				m_spot = 0;
				m_crrent = m_crrent->next =
					reinterpret_cast<Block*>(operator new(sizeof(Block)));
				m_crrent->next = nullptr;
			}
			obj = &m_crrent->data[m_spot++];
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
