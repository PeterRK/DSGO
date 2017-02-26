// +build amd64

package perfect

//go:noescape
func MurmurHash(seed uint32, str string) uint32
