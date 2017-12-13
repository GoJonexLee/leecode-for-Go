package medium

import (
	"container/heap"
	"sort"
)

type Node struct {
	Words []string
	Cont  int
}

type Nodes []*Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) Less(i, j int) bool {
	return n[i].Cont > n[j].Cont
}

func (n *Nodes) Push(i interface{}) {
	*n = append(*n, i.(*Node))
}

func (n *Nodes) Pop() interface{} {
	old := *n
	nc := len(old)
	x := old[nc-1]
	*n = old[0 : nc-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	tmpMap := make(map[string]int)

	for _, word := range words {
		tmpMap[word]++
	}

	ctMap := make(map[int][]string)
	for wd, ct := range tmpMap {
		if _, ok := ctMap[ct]; !ok {
			ctMap[ct] = []string{}
		}
		ctMap[ct] = append(ctMap[ct], wd)
	}

	for _, v := range ctMap {
		sort.Strings(v)
	}

	hp := &Nodes{}
	for k, v := range ctMap {
		tmp := &Node{v, k}
		heap.Push(hp, tmp)
	}

	re := []string{}
LEE:
	tmpNode := heap.Pop(hp).(*Node)
	for i := 0; i < len(tmpNode.Words); i++ {
		re = append(re, tmpNode.Words[i])
		if len(re) == k {
			break
		}
	}
	if len(re) < k {
		goto LEE
	}

	return re
}
