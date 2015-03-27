package cuckoo

type node struct {
	code [2]uint
	key  string
}
type coreTable struct {
	hash   func(str string) uint
	mask   uint
	id     uint8 //0或1
	bucket []*node
}
type HashTable struct {
	first, second coreTable
	cnt           int
}

func (table *HashTable) Size() int {
	return table.cnt
}
func (table *HashTable) IsEmpty() bool {
	return table.cnt == 0
}

func (table *HashTable) Initialize(fn1 func(str string) uint, fn2 func(str string) uint) {
	table.cnt = 0
	table.first.hash, table.second.hash = fn1, fn2
	table.first.id, table.second.id = 0, 1
	table.first.bucket, table.second.bucket = make([]*node, 32), make([]*node, 16) //保持first为second的两倍
	table.first.mask, table.second.mask = 0x1f, 0xf
}
