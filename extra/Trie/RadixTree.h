#pragma once

class RadixTree : NonCopyable {
public:
	typedef unsigned Key;

private:
	static constexpr unsigned UINT_SZ = sizeof(Key) * 8;
	static constexpr unsigned STEP = 2; //2 or 4
	static constexpr unsigned SIZE = 1 << STEP;
	static constexpr unsigned DEPTH = UINT_SZ / STEP;
	static constexpr unsigned MASK = SIZE - 1;
	static unsigned Cut(Key key, unsigned i);

	struct Node {
		Node* kids[SIZE];
	};
	Node* NewNode(void);

	Node* m_root;
	Allocator<Node> m_pool;

public:
	RadixTree(void);

	bool Insert(Key key, void* ptr);
	const void* Search(Key key) const;
	void* Remove(Key key);
};

