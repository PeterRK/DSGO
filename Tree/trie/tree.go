package trie

type core struct {
	keys    [8]byte
	keysCnt uint8
	level   uint8
	kidsCnt uint8
	ref     uint32
}
type node1 struct {
	core
	kids [1]*core
}
type node2 struct {
	core
	kids [2]*core
}
type node4 struct {
	core
	kids [4]*core
}
type node8 struct {
	core
	kids [8]*core
}
type node16 struct {
	core
	kids [16]*core
}
type node32 struct {
	core
	kids [32]*core
}
type node64 struct {
	core
	kids [64]*core
}
type node128 struct {
	core
	kids [128]*core
}
type node256 struct {
	core
	kids [256]*core
}
type node node256
