//go:build amd64

package hashtable

//go:noescape
func Hash32(seed uint32, str string) uint32
