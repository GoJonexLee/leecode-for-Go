package hard

import (
	"math/rand"
	"time"
)

/*
问题描述：
从单链表中随机返回一个节点的值，并且每个节点的返回概率相同；
ps:如果单链表的长度很大，并且不知道单链表的长度，能否在空间复杂度为O(1)的限制下解决该问题。

算法描述：时间复杂度O(n), 空间复杂度O(1)
1.Solution记录点链表的长度以及原始链表
2.Constructor遍历单链表，初始化Solution；
3.根据随机数，遍历单链表，返回节点值。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	count int
	next  *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	s := Solution{count: 0, next: head}
	for head != nil {
		s.count++
	}
	return s
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	n, tp := rd.Intn(this.count), this.next
	for ; n > 0; n-- {
		tp = tp.Next
	}
	return tp.Val
}
