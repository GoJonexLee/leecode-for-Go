package head

import (
	"container/list"
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

func Constructor(capacity int) LFUCache {
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

	if this.freqKeys[this.minFreq].Len() == 0 {
		this.minFreq++
	}

	nd := e.(node)
	return nd.v
}

func (this *LFUCache) Put(key int, value int) {

	if this.cp <= 0 {
		return
	}

	if this.Get(key) != -1 {
		this.keyIdx[key].Value = node{key, value}
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
