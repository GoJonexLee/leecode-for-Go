package easy

type ListNost struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	
	newhead := &ListNode{0, head}	// 需要重新起一个头结点的前结点
	pre := newhead

	for pre.Next != nil {
		if pre.Next.Val == val {
			tmp := pre.Next
			pre.Next = tmp.Next
			tmp.Next = nil
		} else {
			pre = pre.Next
		}
	}

	return newhead.Next
}

func removeElement(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElement(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
}