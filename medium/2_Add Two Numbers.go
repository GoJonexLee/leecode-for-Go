package medium

type ListNode struct {
	val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    if l1 == nil {
        return l2
    }
    
    if l2 == nil {
        return l1
    }
    
    result := &ListNode{}
    ptr, carry := result, 0
    
    for l1 != nil || l2 != nil {
        if l1 != nil {
            carry += l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            carry += l2.Val
            l2 = l2.Next
        }
        ptr.Next = &ListNode{Val:carry%10, Next:nil}
        carry /= 10
        ptr = ptr.Next
    }
    
    if carry == 1 {
        ptr.Next = &ListNode{carry, nil}
    }
    
    return result.Next
}