package medium

import (
	"sort"
	"bytes"
	"container/heap"
)

type Node struct {
	Num int
	Lts []int
}

type Nodes []*Node

func (n Nodes) Less(i, j int) bool {
	return n[i].Num > n[j].Num
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) Len() int {
	return len(n)
}

func (n *Nodes) Pop() interface{} {
	old := *n
	nc := len(old)
	x := old[nc-1]
	*n = old[:nc-1]
	return x
}

func (n *Nodes) Push(i interface{}) {
	*n = append(*n, i.(*Node))
}

func frequencySort(s string) string {
	ct := make(map[int]int)
	for _, tmp := range s {
		ct[int(tmp)]++
	}

	lst := make(map[int][]int)
	for word, num := range ct {
		if _, ok := lst[num]; !ok {
			lst[num] = []int{}
		}
		lst[num] = append(lst[num], word)
	}

	re := &Nodes{}
	for k, v := range lst {
		tmp := &Node{k, v}
		heap.Push(re, tmp)
	}

	r := bytes.NewBufferString("")
	for re.Len() != 0 {
		tmp := heap.Pop(re).(*Node)
		sort.Ints(tmp.Lts)
		for i := 0; i < len(tmp.Lts); i++ {
			for j := 0; j < tmp.Num; j++ {
				r.WriteString(string(tmp.Lts[i]))
			}
		}
	}
	return r.String()
}