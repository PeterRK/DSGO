#pragma once

class NonCopyable {
protected:
	NonCopyable(void) {}
	~NonCopyable(void) {}
private:
	NonCopyable(const NonCopyable&);
	const NonCopyable& operator=(const NonCopyable&);
};

template<typename T>
class Allocator : NonCopyable {
private:
	static const int SIZE = 1023;
	union Obj {
		char	data[sizeof(T)];
		Obj*	next;
	};
	struct Block {
		Obj		data[SIZE];
		Block*	next;
	};

	int		m_spot;
	Block*	m_crrent;
	Block*	m_head;
	Obj*	m_free;
	int		m_balance;

public:
	Allocator(void)
		: m_spot(0), m_crrent(new Block), m_head(m_crrent),
		m_free(NULL), m_balance(0)
	{ m_crrent->next = NULL; }
	~Allocator(void) {
		while (m_head != NULL) {
			m_crrent = m_head;
			m_head = m_head->next;
			operator delete(m_crrent);
		}
	}
	int balance(void) const { return m_balance; }

	T* allocate(void) {
		Obj* obj = m_free;
		if (obj != NULL) {
			m_free = m_free->next;
		} else {
			if (m_spot == SIZE) {
				m_spot = 0;
				m_crrent = m_crrent->next =
					reinterpret_cast<Block*>(operator new(sizeof(Block)));
				m_crrent->next = NULL;
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