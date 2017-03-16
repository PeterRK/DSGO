package lrucache

import (
	"unsafe"
)

type node struct {
	prev *node
	next *node
	hot  bool
	key  int
	val  string
}
type knot struct {
	prev  *node
	next  *node
	count uint
	limit uint
}
type cache struct {
	hot  knot
	cold knot
	book map[int]*node
}

func (k *knot) initialize(limit uint) {
	k.count = 0
	k.limit = limit
	k.prev = (*node)(unsafe.Pointer(k))
	k.next = (*node)(unsafe.Pointer(k))
}

func New(hot_sz, cold_sz uint) *cache {
	if cold_sz < 4 || hot_sz < 4 {
		return nil
	}
	var obj = new(cache)
	obj.book = make(map[int]*node)
	obj.hot.initialize(hot_sz)
	obj.cold.initialize(cold_sz)
	return obj
}

func (u *node) unhook() {
	u.prev.next = u.next
	u.next.prev = u.prev
}
func (u *node) hook(prev, next *node) {
	prev.next, u.prev = u, prev
	next.prev, u.next = u, next
}

func (c *cache) access(u *node) {
	u.unhook()
	u.hook((*node)(unsafe.Pointer(&c.hot)), c.hot.next)

	if !u.hot { //冷热迁移
		u.hot = true
		c.cold.count--
		if c.hot.count < c.hot.limit {
			c.hot.count++
		} else { //末位淘汰
			u = c.hot.prev
			u.unhook()
			delete(c.book, u.key)
		}
	}
}

func (c *cache) Clear() {
	c.book = make(map[int]*node)
	c.hot.initialize(c.hot.limit)
	c.cold.initialize(c.cold.limit)
}

func (c *cache) Insert(key int, val string) {
	var u, ok = c.book[key]
	if ok { //如果已经被缓存则更新其值
		u.val = val
		c.access(u)
	} else { //如果目标没有被缓存则增加此项目
		u = new(node)
		u.hot = false
		u.key, u.val = key, val
		c.book[key] = u
		u.hook((*node)(unsafe.Pointer(&c.cold)), c.cold.next)
		if c.cold.count < c.cold.limit {
			c.cold.count++
		} else { //末位淘汰
			u = c.cold.prev
			u.unhook()
			delete(c.book, u.key)
		}
	}
}

func (c *cache) Search(key int) (val string, ok bool) {
	u, ok := c.book[key]
	if ok {
		c.access(u)
		return u.val, true
	}
	return "", false
}

func (c *cache) Remove(key int) {
	var u, ok = c.book[key]
	if ok {
		delete(c.book, u.key)
		u.unhook()
		if u.hot {
			c.hot.count--
		} else {
			c.cold.count--
		}
	}
}
