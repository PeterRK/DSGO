package ring

import (
	"DSGO/utils"
	"unsafe"
)

type nodeC[K comparable, V any] struct {
	node
	age uint8
	key K
	val V
}

func castC[K comparable, V any](unit *node) *nodeC[K, V] {
	return (*nodeC[K, V])(unsafe.Pointer(unit))
}

type queue struct {
	ring
	free, limit int
}

func (q *queue) reset(limit int) {
	q.init()
	if limit < 4 {
		limit = 4
	}
	q.free, q.limit = limit, limit
}

type cache[K comparable, V any] struct {
	hot, cold queue
	index     map[K]*nodeC[K, V]
	warm uint8
}

//warm一般调成0或1
func NewCacheDetail[K comparable, V any](hot, cold int, warm uint8) utils.Cache[K, V] {
	c := new(cache[K, V])
	c.init(hot, cold, warm)
	return c
}

func NewCache[K comparable, V any](capacity int) utils.Cache[K, V] {
	return NewCacheDetail[K, V](capacity/4, capacity*3/4, 1)
}

func (c *cache[K, V]) init(hot, cold int, warm uint8) {
	if warm > 3 {
		warm = 3
	}
	c.hot.reset(hot)
	c.cold.reset(cold)
	c.index = make(map[K]*nodeC[K, V])
}

func (c *cache[K, V]) Clear() {
	c.hot.reset(c.hot.limit)
	c.cold.reset(c.cold.limit)
	c.index = make(map[K]*nodeC[K, V])
}

func (c *cache[K, V]) Size() int {
	return len(c.index)
}

func (c *cache[K, V]) Capacity() int {
	return c.hot.limit + c.cold.limit
}

func (c *cache[K, V]) touch(unit *nodeC[K, V]) {
	unit.Release()
	if unit.age > c.warm { //热循环
		c.hot.pushHead(&unit.node)
		return
	}
	unit.age++
	if unit.age <= c.warm { //冷循环
		c.cold.pushHead(&unit.node)
		return
	}
	//冷热迁移
	c.hot.pushHead(&unit.node)
	if c.hot.free > 0 {
		c.hot.free--
		c.cold.free++
	} else {
		victim := c.hot.unsafePopTail()
		castC[K, V](victim).age = c.warm
		c.cold.pushHead(victim)
	}
}

func (c *cache[K, V]) Put(key K, val V) {
	if unit := c.index[key]; unit != nil {
		unit.val = val //如果已经被缓存则更新其值
		c.touch(unit)
	} else { //如果目标没有被缓存则增加此项目
		unit = new(nodeC[K, V])
		unit.age = 0
		unit.key, unit.val = key, val
		c.index[key] = unit
		c.cold.pushHead(&unit.node)
		if c.cold.free > 0 {
			c.cold.free--
		} else { //末位淘汰
			victim := c.cold.unsafePopTail()
			delete(c.index, castC[K, V](victim).key)
		}
	}
}

func (c *cache[K, V]) Get(key K) (val V, found bool) {
	if unit := c.index[key]; unit != nil {
		c.touch(unit)
		return unit.val, true
	}
	var tmp V
	return tmp, false
}

func (c *cache[K, V]) Discard(key K) {
	if unit := c.index[key]; unit != nil {
		delete(c.index, unit.key)
		unit.Release()
		if unit.age > c.warm {
			c.hot.free++
		} else {
			c.cold.free++
		}
	}
}
