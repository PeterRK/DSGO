package hash

//对于已知输入和基础hash函数f1(x),f2(x)，构造完全hash函数F(x)=(A*f1(x)+B*f2(x))/M。
//当返回M为0时表示构造失败。
func PerfectHash(words [][]byte, f1, f2 func([]byte) uint) (M uint, A, B uint) {
	var size = len(words)
	var tb1, tb2 = make([]uint, size), make([]uint, size)
	for i, word := range words {
		tb1[i], tb2[i] = f1(word), f2(word)
	}

	return 0, 0, 0
}
