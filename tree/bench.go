package tree

import (
	"DSGO/linkedlist/skiplist"
	"DSGO/tree/avl"
	"DSGO/tree/avl/weak"
	"DSGO/tree/bplus"
	"DSGO/tree/redblack"
	"DSGO/tree/redblack/compact"
	"fmt"
	"math/rand"
	"time"
)

type U32Set interface {
	Size() int
	IsEmpty() bool
	Clear()
	Search(uint32) bool
	Insert(uint32) bool
	Remove(uint32) bool
}

type result struct {
	insert time.Duration
	search time.Duration
	remove time.Duration
}

func (r *result) add(d result) {
	r.insert += d.insert
	r.search += d.search
	r.remove += d.remove
}

func benchmark(list []uint32, create func() U32Set) result {
	len1, len2, len3 := len(list)/3, len(list)*2/3, len(list)
	set := create()
	var res result

	start := time.Now()
	for i := 0; i < len2; i++ {
		set.Insert(list[i])
	}
	finish := time.Now()
	res.insert += finish.Sub(start)

	start = finish
	for i := len1; i < len2; i++ {
		set.Remove(list[i])
	}
	finish = time.Now()
	res.remove += finish.Sub(start)

	start = finish
	for i := 0; i < len2; i++ {
		set.Search(list[i])
	}
	finish = time.Now()
	res.search += finish.Sub(start)

	start = finish
	for i := len2; i < len3; i++ {
		set.Insert(list[i])
	}
	finish = time.Now()
	res.insert += finish.Sub(start)

	start = finish
	for i := 0; i < len1; i++ {
		set.Remove(list[i])
	}
	finish = time.Now()
	res.remove += finish.Sub(start)

	start = finish
	for i := len2; i < len3; i++ {
		set.Search(list[i])
	}
	finish = time.Now()
	res.search += finish.Sub(start)

	return res
}

func Benchmark(size, round int, extend bool) {
	if round < 1 {
		return
	}
	if size < 10000 {
		size = 10000
	}
	list := make([]uint32, size)
	rand.Seed(time.Now().UnixNano())

	var sklTime result
	runSkl := func() {
		sklTime.add(benchmark(list, func() U32Set {
			skl := skiplist.New[uint32]()
			return skl
		}))
	}

	var bpTime result
	runBp := func() {
		bpTime.add(benchmark(list, func() U32Set {
			return new(bplus.Tree[uint32])
		}))
	}

	var rbTime result
	runRb := func() {
		rbTime.add(benchmark(list, func() U32Set {
			if extend {
				return new(compact.Tree[uint32])
			} else {
				return new(redblack.Tree[uint32])
			}
		}))
	}

	var avlTime result
	runAvl := func() {
		avlTime.add(benchmark(list, func() U32Set {
			return new(avl.Tree[uint32])
		}))
	}

	var wavlTime result
	runWavl := func() {
		wavlTime.add(benchmark(list, func() U32Set {
			return new(weak.Tree[uint32])
		}))
	}

	for i := 0; i < round; i++ {
		for j := 0; j < size; j++ {
			list[j] = rand.Uint32()
		}
		if extend {
			runSkl()
		}

		runBp()
		runRb()
		runAvl()
		runWavl()

		runBp()
		runAvl()
		runWavl()
		runRb()

		runBp()
		runWavl()
		runRb()
		runAvl()
	}

	div := time.Duration(round)

	if extend {
		fmt.Println("SKL")
		fmt.Println("Insert", sklTime.insert/div)
		fmt.Println("Search", sklTime.search/div)
		fmt.Println("Remove", sklTime.remove/div)
	}

	div *= 3

	fmt.Println("\nB+")
	fmt.Println("Insert", bpTime.insert/div)
	fmt.Println("Search", bpTime.search/div)
	fmt.Println("Remove", bpTime.remove/div)

	fmt.Println("\nRB")
	fmt.Println("Insert", rbTime.insert/div)
	fmt.Println("Search", rbTime.search/div)
	fmt.Println("Remove", rbTime.remove/div)

	fmt.Println("\nAVL")
	fmt.Println("Insert", avlTime.insert/div)
	fmt.Println("Search", avlTime.search/div)
	fmt.Println("Remove", avlTime.remove/div)

	fmt.Println("\nWAVL")
	fmt.Println("Insert", wavlTime.insert/div)
	fmt.Println("Search", wavlTime.search/div)
	fmt.Println("Remove", wavlTime.remove/div)
}
