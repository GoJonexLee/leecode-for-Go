package hard

import (
	"container/list"
)

type LRUCache struct {
	l    *list.List
	m    map[int]*list.Element
	size int
}

type node struct {
	k, v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l:    list.New(),
		m:    make(map[int]*list.Element),
		size: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}

	this.l.MoveToFront(this.m[key])
	return this.l.Front().Value.(node).v
}

func (this *LRUCache) Put(key int, value int) {

	if e, ok := this.m[key]; ok {
		this.l.Remove(e)
		delete(this.m, key)
	}

	if this.size == len(this.m) {
		n := this.l.Back()
		this.l.Remove(n)

		nd := n.Value.(node)
		delete(this.m, nd.k)
	}

	n := node{key, value}
	this.m[key] = this.l.PushFront(n)
}
