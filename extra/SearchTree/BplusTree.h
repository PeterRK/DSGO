#pragma once

class BplusTree {
//private:
public:
	static const int BASE_SIZE = 16;
	struct Node {
		bool	inner;
		int		cnt;
		int		data[BASE_SIZE];
		int ceil(void) const { return data[cnt - 1]; }
		int locate(int key) const;
	};

	struct Index : Node {
		static const int HALF_SZ = BASE_SIZE;
		static const int FULL_SZ = HALF_SZ * 2 - 1;
		static const int QUARTER_SZ = HALF_SZ / 2;

		int		_remain[FULL_SZ - BASE_SIZE];
		Index*	kids[FULL_SZ];

		static Index* NewNode(Allocator<Index>& pool);
		void remove(int place);
		Index* insert(int place, Index* kid, Allocator<Index>& pool);
		bool combine(Index* peer, Allocator<Index>& pool);
	};

	struct Leaf : Node {
		static const int HALF_SZ = BASE_SIZE * 2 - 1;
		static const int FULL_SZ = HALF_SZ * 2 - 1;
		static const int QUARTER_SZ = HALF_SZ / 2;

		int		_remain[FULL_SZ - BASE_SIZE];
		Leaf*	next;

		static Leaf* NewNode(Allocator<Leaf>& pool);
		void remove(int place);
		Leaf* insert(int place, int key, Allocator<Leaf>& pool);
		bool combine(Leaf* peer, Allocator<Leaf>& pool);
	};

	Index*	m_root;
	Leaf*	m_head;
	Allocator<Index>	m_ipool;
	Allocator<Leaf>		m_lpool;
	std::vector<Index*>	m_pstack;
	std::vector<int>	m_nstack;

public:
	BplusTree(void) : m_root(nullptr), m_head(nullptr) {}
	typedef void Func(int);

	void travel(Func func) const;
	bool isEmpty(void) const { return m_root == nullptr; }
	bool search(int key) const;
	bool insert(int key);
	bool remove(int key);
};

