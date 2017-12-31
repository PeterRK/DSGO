package perfect

import (
	"errors"
)

type node struct {
	filled bool
	code   uint32
	val    string
}

type Table struct {
	hint   []uint32
	bucket []node
}

const VALUES_PER_HINT = 4
const TRYS_LIMIT = 1000
const DEFAULT_SEED = uint32(0)
const SEED_STEP = 0x9e3779b9 //非零魔数

type memo struct {
	idx uint32
	lst []string
}

func (tb *Table) Build(data []string) error {
	return tb.BuildWithSeed(data, 0)
}

//不支持空串，不支持空集
func (tb *Table) BuildWithSeed(data []string, seed uint32) error {
	m := len(data)
	if m == 0 {
		return errors.New("cannot build empty table")
	}
	n := uint32(m + m/10) //十分之一的冗余，目标容积率0.9
	if i := binarySearch(primes, n); i < len(primes) {
		n = primes[i] //
	}

	hn := uint32((m + (VALUES_PER_HINT - 1)) / VALUES_PER_HINT)
	book := make([]memo, hn)
	for i := uint32(0); i < hn; i++ {
		book[i].idx = i
	}

	//初级Hash
	for _, val := range data {
		code := MurmurHash(DEFAULT_SEED, val)
		cell := &book[code%hn]
		cell.lst = append(cell.lst, val)
	}

	//剔除空项目
	for i := 0; i < len(book); {
		if len(book[i].lst) == 0 {
			book[i] = book[len(book)-1]
			book = book[:len(book)-1] //抽尾补中
		} else {
			i++
		}
	}
	sort(book) //根据lst长度排序

	hint := make([]uint32, hn)
	bucket := make([]node, n)

	dirty := make([]uint32, 0, len(book[len(book)-1].lst))

	//先大后下小
	for i := len(book) - 1; i >= 0; i-- {
		trys := TRYS_LIMIT
		for ; trys > 0; trys-- {
			dirty = dirty[:0]
			seed += SEED_STEP

			j := 0
			for lst := book[i].lst; j < len(lst); j++ {
				code := MurmurHash(seed, lst[j])
				index := code % n
				if bucket[index].filled {
					break //冲突了
				}
				bucket[index] = node{true, code, lst[j]}
				dirty = append(dirty, index)
			}
			if j == len(book[i].lst) {
				hint[book[i].idx] = seed
				break //成功
			}
			//失败回滚
			for j = 0; j < len(dirty); j++ {
				bucket[dirty[j]] = node{false, 0, ""}
			}
		}
		if trys <= 0 {
			return errors.New("cannot resolve conflicts")
		}
	}

	tb.hint = hint
	tb.bucket = bucket
	return nil
}

func (tb *Table) Serach(val string) bool {
	hn, n := uint32(len(tb.hint)), uint32(len(tb.bucket))
	if hn == 0 || n == 0 {
		return false
	}

	index := MurmurHash(DEFAULT_SEED, val) % hn
	code := MurmurHash(tb.hint[index], val)

	cell := tb.bucket[code%n]
	return cell.code == code && cell.val == val
}
