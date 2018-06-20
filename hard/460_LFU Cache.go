package main

import (
	"container/list"
	"fmt"
)

type LFUCache struct {
	kvCount  map[int]int
	freqKeys map[int]*list.List
	keyIdx   map[int]*list.Element

	cp, minFreq int
}

type node struct {
	k, v int
}

func Constructor1(capacity int) LFUCache {
	return LFUCache{
		kvCount:  make(map[int]int),
		freqKeys: make(map[int]*list.List),
		keyIdx:   make(map[int]*list.Element),
		cp:       capacity,
	}
}

func (this *LFUCache) Get(key int) int {

	if _, ok := this.keyIdx[key]; !ok {
		return -1
	}

	freq, eIdx := this.kvCount[key], this.keyIdx[key]
	if _, ok := this.freqKeys[freq+1]; !ok {
		this.freqKeys[freq+1] = list.New()
	}

	e := this.freqKeys[freq].Remove(eIdx)
	newE := this.freqKeys[freq+1].PushFront(e)
	this.kvCount[key] = freq + 1
	this.keyIdx[key] = newE

	if this.freqKeys[freq].Len() == 0 {
		this.minFreq = freq + 1
	}

	nd := e.(node)
	return nd.v
}

func (this *LFUCache) Put(key int, value int) {

	if this.cp <= 0 {
		return
	}

	if this.Get(key) != -1 {
		freq, idx := this.kvCount[key], this.keyIdx[key]
		this.freqKeys[freq].Remove(idx)
		e := this.freqKeys[freq].PushFront(node{key, value})
		this.keyIdx[key] = e
		return
	}

	if len(this.kvCount) >= this.cp {
		lastNode := this.freqKeys[this.minFreq].Back()
		nd := this.freqKeys[this.minFreq].Remove(lastNode).(node)
		delete(this.keyIdx, nd.k)
		delete(this.kvCount, nd.k)
	}

	this.minFreq = 1
	if _, ok := this.freqKeys[this.minFreq]; !ok {
		this.freqKeys[this.minFreq] = list.New()
	}

	this.kvCount[key] = this.minFreq
	newNode := this.freqKeys[this.minFreq].PushFront(node{key, value})
	this.keyIdx[key] = newNode
}

func main() {
	cache := Constructor1(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))

	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

	cache.Put(2, 1)
	fmt.Println(cache.Get(2))
}
