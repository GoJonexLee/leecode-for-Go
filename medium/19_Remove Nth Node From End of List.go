package medium

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	
	pre, tail := head, head
	for n > 0 {
		tail = tail.
		n--
	}

	if tail == nil {
		return head.Next
	}

	for tail.Next != nil {
		pre = pre.Next
		tail = tail.Next
	}

	pre.Next = pre.Next.Next
	return head
}