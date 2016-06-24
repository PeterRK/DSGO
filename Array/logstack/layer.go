package logstack

import (
	"fmt"
	"unsafe"
)

type usec uint

const USSZ = uint(unsafe.Sizeof(usec(0)) * 8)

type layer struct {
	data []int
	mark []usec
}

func tellMark(bitvec []usec, place uint) bool {
	var idx, sft = place / USSZ, place % USSZ
	return (bitvec[idx] & (usec(1) << sft)) != 0
}

func changeMark(bitvec []usec, place uint, mk bool) {
	var idx, sft = place / USSZ, place % USSZ
	if mk {
		bitvec[idx] |= usec(1) << sft
	} else {
		bitvec[idx] &= ^(usec(1) << sft)
	}
}

func (ly *layer) reset() {
	ly.data = nil
	ly.mark = nil
}

func (ly *layer) size() uint {
	return uint(len(ly.data))
}

func (ly *layer) change(key int, mk bool) {
	var a, b = uint(0), ly.size()
	for a < b {
		var m = (a + b) / 2
		switch {
		case key > ly.data[m]:
			a = m + 1
		case key < ly.data[m]:
			b = m
		default:
			changeMark(ly.mark, m, mk)
			return
		}
	}

	if ly.size()%USSZ == 0 {
		ly.mark = append(ly.mark, 0)
	}
	ly.data = append(ly.data, 0)

	var idx = a / USSZ
	for i := (ly.size() - 1) / USSZ; i > idx; i-- {
		ly.mark[i] = (ly.mark[i] << 1) | (ly.mark[i-1] >> (USSZ - 1))
	}
	var mask = ^usec(0) << (a % USSZ)
	ly.mark[idx] = (ly.mark[idx] & ^mask) | ((ly.mark[idx] & mask) << 1)

	changeMark(ly.mark, a, mk)

	for b = ly.size() - 1; b > a; b-- {
		ly.data[b] = ly.data[b-1]
	}
	ly.data[a] = key
}

//返回0表示没有，返回1表示有，返回-1表示标记删除
func (ly *layer) search(key int) int {
	var a, b = uint(0), ly.size()
	for a < b {
		var m = (a + b) / 2
		switch {
		case key > ly.data[m]:
			a = m + 1
		case key < ly.data[m]:
			b = m
		default:
			if tellMark(ly.mark, m) {
				return 1
			} else {
				return -1
			}
		}
	}
	return 0
}

func (ly *layer) extend(key int, mk bool) {
	if ly.size()%USSZ == 0 {
		ly.mark = append(ly.mark, 0)
	}
	changeMark(ly.mark, ly.size(), mk)
	ly.data = append(ly.data, key)
}
func (ly *layer) merge(vic *layer) {
	var nly layer
	var a, b = uint(0), uint(0)
	for a < ly.size() && b < vic.size() {
		if ly.data[a] < vic.data[b] {
			nly.extend(ly.data[a], tellMark(ly.mark, a))
			a++
		} else {
			if ly.data[a] == vic.data[b] {
				a++
			}
			nly.extend(vic.data[b], tellMark(vic.mark, b))
			b++
		}
	}
	for a < ly.size() {
		nly.extend(ly.data[a], tellMark(ly.mark, a))
		a++
	}
	for b < vic.size() {
		nly.extend(vic.data[b], tellMark(vic.mark, b))
		b++
	}
	*ly = nly
}

func (ly *layer) compact() {
	var w = uint(0)
	for r := uint(0); r < ly.size(); r++ {
		if tellMark(ly.mark, r) {
			ly.data[w] = ly.data[r]
			w++
		}
	}
	ly.data = ly.data[:w]
	ly.mark = ly.mark[0 : (w+(USSZ-1))/USSZ]
	for i := 0; i < len(ly.mark); i++ {
		ly.mark[i] = ^usec(0)
	}
}

func (ly *layer) print() {
	for i := uint(0); i < ly.size(); i++ {
		if tellMark(ly.mark, i) {
			fmt.Printf("X%d ", ly.data[i])
		} else {
			fmt.Printf("_%d ", ly.data[i])
		}
	}
	fmt.Println()
}

func (ly *layer) debug() {
	for i := uint(0); i < ly.size(); i++ {
		if tellMark(ly.mark, i) {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
}
