package easy

/*
问题描述：
反转链表

算法描述：时间复杂度O(n), 空间复杂度O(1)
1.遍历链表；
2.头插法重新超入到新的头节点中；
3.返回新头节点的next。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	root, tp := new(ListNode), head
	for tp != nil {
		tp = head.Next
		head.Next = root.Next
		root.Next = head
		head = tp
	}
	return root.Next
}
