package easy

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head, nhead *ListNode
	if l1.Val < l2.Val {
		head, nhead = l1, l2
	} else {
		head, nhead = l2, l1
	}

	head.Next = mergeTwoLists(head.Next, nhead)
	return head
}