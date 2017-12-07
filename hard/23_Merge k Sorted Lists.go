package hard

import (
	"container/heap"
)

type Heap []*ListNode

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Len() int {
	return len(h)
}

func (h *Heap) Push(i interface{}) {
	*h = append(*h, i.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	hp := &Heap{}

	for _, head := range lists {
		if head != nil {
			heap.Push(hp, head)
		}
	}

	var head, pre, tmp *ListNode
	for hp.Len() != 0 {
		tmp = heap.Pop(hp).(*ListNode)
		if pre == nil {
			head = tmp
		} else {
			pre.Next = tmp
		}
		pre = tmp
		if tmp.Next != nil {
			heap.Push(hp, tmp.Next)
		}
	}

	return head
}
