#include <cstdint>
#include <cassert>
#include <ctime>
#include <cstdio>
#include <iostream>
#include <vector>
#include <random>
#include <functional>
#include "tricks.h"
#include "RBtree.h"
#include "AVLtree.h"
#include "BplusTree.h"
#include "SkipList.h"

template<class T>
void BenchMark(std::vector<int>& list)
{
	clock_t start, end;
	const unsigned size = list.size();
	T dict;

	start = clock();
	for (unsigned i = 0; i < size; i++) {
		dict.insert(list[i]);
	}
	end = clock();
	std::cout << "[insert] " << (end - start) / (double)CLOCKS_PER_SEC << "s" << std::endl;
	start = clock();
	for (unsigned i = 0; i < size; i++) {
		dict.search(list[i]);
	}
	end = clock();
	std::cout << "[search] " << (end - start) / (double)CLOCKS_PER_SEC << "s" << std::endl;
	start = clock();
	for (unsigned i = 0; i < size; i++) {
		dict.remove(list[i]);
	}
	end = clock();
	std::cout << "[remove] " << (end - start) / (double)CLOCKS_PER_SEC << "s" << std::endl;
}

void RunAllBenchmark(void)
{
	auto rand_int = std::bind(
		std::uniform_int_distribution<int>(), std::default_random_engine(std::random_device()())
	);
	std::ios_base::sync_with_stdio(false);

	const unsigned size = 1000000;
	std::vector<int> list(size);
	for (unsigned i = 0; i < size; i++) {
		list[i] = rand_int();
	}

	std::cout << "SkipList" << std::endl;
	BenchMark<SkipList>(list);

	std::cout << "AVL Tree" << std::endl;
	BenchMark<AVLtree>(list);

	std::cout << "Red-Black Tree" << std::endl;
	BenchMark<RBtree>(list);

	std::cout << "B+ Tree" << std::endl;
	BenchMark<BplusTree>(list);
}

void DebugAVL(void)
{
	auto rand = std::bind(
		std::uniform_int_distribution<int>(), std::default_random_engine(std::random_device()())
	);

	const unsigned size = 200;
	std::vector<int> list(size);
	for (unsigned i = 0; i < size; i++) {
		list[i] = rand();
	}

	AVLtree tree;
	unsigned cnt = 0;

	for (unsigned i = 0; i < size; i++) {
		if (tree.insert(list[i]) > 0)
			cnt++;
	}

	for (unsigned i = 0; i < size; i++) {
		assert(tree.search(list[i]) != -1);
		assert(tree.insert(list[i]) < 0);
	}

	for (unsigned i = 0; i < size; i++) {
		if (tree.remove(list[i]) > 0)
			cnt--;
		assert(tree.search(list[i]) == -1);
	}

	assert(tree.isEmpty() && cnt == 0);
	assert(tree.remove(0) < 0);
}

void DebugRank(void)
{
	auto rand = std::bind(
		std::uniform_int_distribution<unsigned>(), std::default_random_engine(std::random_device()())
	);

	const unsigned size = 200;
	std::vector<int> list(size);
	for (unsigned i = 0; i < size; i++) {
		list[i] = i + 1;
	}
	for (unsigned i = 1; i < size; i++) {
		unsigned j = rand() % (i + 1);
		std::swap(list[i], list[j]);
	}
	std::vector<int> shadow(list);

	AVLtree tree;
	assert(tree.insert(shadow[0]) == 1);
	for (unsigned i = 1; i < size; i++) {
		int rank = tree.insert(shadow[i]);

		int key = shadow[i];
		unsigned a = 0, b = i;
		while (a < b) {
			unsigned m = a + (b - a) / 2;
			if (key < shadow[m]) {
				b = m;
			} else {
				a = m + 1;
			}
		}
		for (unsigned j = i; j > a; j--)
			shadow[j] = shadow[j - 1];
		shadow[a] = key;

		assert(rank == (int)(a + 1));
	}

	for (unsigned i = 0; i < size; i++) {
		assert(tree.search(i + 1) == (int)(i + 1));
	}

	for (unsigned i = 0; i < size; i++) {
		int rank = tree.remove(list[i]);

		int key = list[i];
		unsigned a = 0, b = size - i;
		while (a < b) {
			unsigned m = a + (b - a) / 2;
			if (key < shadow[m]) {
				b = m;
			} else {
				a = m + 1;
			}
		}

		for (unsigned j = a; j < size - i; j++) {
			shadow[j - 1] = shadow[j];
		}
		assert(rank == (int)a);
	}
}

int main(void)
{
//	DebugAVL();
//	DebugRank();
//	std::cout << "pass test" << std::endl;

	RunAllBenchmark();
	return 0;
}
