#include <cstdint>
//#include <cstdlib>
#include <ctime>
#include <cstdio>
#include <vector>
#include <random>
#include <functional>
#include "tricks.h"
#include "AVLtree.h"
#include "BplusTree.h"
#include "RBtree.h"
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
	printf("[insert] %.3fs\n", (end - start) / (double)CLOCKS_PER_SEC);
	start = clock();
	for (unsigned i = 0; i < size; i++) {
		dict.search(list[i]);
	}
	end = clock();
	printf("[search] %.3fs\n", (end - start) / (double)CLOCKS_PER_SEC);
	start = clock();
	for (unsigned i = 0; i < size; i++) {
		dict.remove(list[i]);
	}
	end = clock();
	printf("[remove] %.3fs\n", (end - start) / (double)CLOCKS_PER_SEC);
}

int main(int argc, char* argv[])
{
//	srand((unsigned)time(NULL));
	auto rand = std::bind(
		std::uniform_int_distribution<int>(), std::default_random_engine(std::random_device()())
	);

	const unsigned size = 1000000;
	std::vector<int> list(size);
	for (unsigned i = 0; i < size; i++) {
		list[i] = rand();
	}

	puts("AVL Tree");
	BenchMark<AVLtree>(list);

	puts("Red-Black Tree");
	BenchMark<RBtree>(list);

	puts("B+ Tree");
	BenchMark<BplusTree>(list);

	puts("SkipList");
	BenchMark<SkipList>(list);

	return 0;
}
