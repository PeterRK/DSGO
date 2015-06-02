package hash

//对于已知输入和基础hash函数f1(x),f2(x)，构造完全hash函数F(x)=(A*f1(x)+B*f2(x))/M。
//当返回M为0时表示构造失败。
func PerfectHash(words []string, f1 func(string) uint, f2 func(string) uint) (M uint, A uint, B uint) {
	return 0, 0, 0
}
