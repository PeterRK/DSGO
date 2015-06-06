#pragma once

class SkipList {
private:
	static const unsigned FACTOR = 3;
	struct Node {
		int			key;
		int			size;
		Node*		next[1];
	};

	uint32_t	m_rand;
	int			m_cnt, m_ceil, m_floor, m_level;
	int			m_cap;
	Node**		m_heads;
	Node**		m_knots;

	Node* shadow(void) const;

public:
	SkipList(void);
	~SkipList(void);

	typedef void Func(int);
	void travel(Func func) const;
	unsigned size(void) const;
	bool isEmpty(void) const;
	bool search(int key) const;
	bool insert(int key);
	bool remove(int key);
};
