package lrucache

import (
	"unsafe"
)

type node struct {
	prev *node
	next *node
	hot  bool
	key  uint32
	val  string
}
type knot struct {
	prev  *node
	next  *node
	cnt   uint
	limit uint
}
type cache struct {
	hot  knot
	cool knot
	book map[uint32]*node
}

func knotInit(k *knot, limit uint) {
	k.cnt = 0
	k.limit = limit
	k.prev = (*node)(unsafe.Pointer(k))
	k.next = (*node)(unsafe.Pointer(k))
}

func NewCache(hot_sz, cool_sz uint) *cache {
	if cool_sz < 4 || hot_sz < cool_sz {
		return nil
	}
	var obj = new(cache)
	obj.book = make(map[uint32]*node)
	knotInit(&obj.hot, hot_sz)
	knotInit(&obj.cool, cool_sz)
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
		c.cool.cnt--
		if c.hot.cnt < c.hot.limit {
			c.hot.cnt++
		} else { //末位淘汰
			u = c.hot.prev
			u.unhook()
			delete(c.book, u.key)
		}
	}
}

func (c *cache) Insert(key uint32, val string) {
	var u, ok = c.book[key]
	if ok { //如果已经被缓存则更新其值
		u.val = val
		c.access(u)
	} else { //如果目标没有被缓存则增加此项目
		u = new(node)
		u.hot = false
		u.key, u.val = key, val
		c.book[key] = u
		u.hook((*node)(unsafe.Pointer(&c.cool)), c.cool.next)
		if c.cool.cnt < c.cool.limit {
			c.cool.cnt++
		} else { //末位淘汰
			u = c.cool.prev
			u.unhook()
			delete(c.book, u.key)
		}
	}
}

func (c *cache) Search(key uint32) (val string, ok bool) {
	u, ok := c.book[key]
	if ok {
		c.access(u)
		return u.val, true
	}
	return "", false
}

func (c *cache) Remove(key uint32) {
	var u, ok = c.book[key]
	if ok {
		delete(c.book, u.key)
		u.unhook()
		if u.hot {
			c.hot.cnt--
		} else {
			c.cool.cnt--
		}
	}
}
