package hashtable

type HashTable interface {
	Size() int
	IsEmpty() bool
	Insert(key string) bool
	Search(key string) bool
	Remove(key string) bool
}
