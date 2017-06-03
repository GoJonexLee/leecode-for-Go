package medium

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	slow, fast := head, head
	for  fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil
	var tmpHead *ListNode
	for mid != nil {
		tmp := mid.Next
		mid.Next = tmpHead
		tmpHead = mid
		mid = tmp
	}

	for head != nil && tmpHead != nil {
		tmp := head.Next
		head.Next = tmpHead
		tmpHead = tmpHead.Next
		head.Next.Next = tmp
		head = tmp
	}
}
