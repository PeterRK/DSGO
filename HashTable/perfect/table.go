package perfect

type node struct {
	code uint32
	val  []byte //nil为无效
}

type Table struct {
	hints  []uint32
	bucket []node
}

//不支持空串
func (tb *Table) Build(data [][]byte) {

}
