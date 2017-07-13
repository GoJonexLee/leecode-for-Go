package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	la1, la2 := listToArray(l1), listToArray(l2)

	var re *ListNode
	var lBegin, rBegin int
	var pre, sum int
	for lBegin, rBegin = len(la1) - 1, len(la2) -1; lBegin >= 0 && rBegin >= 0; {
		sum = la1[lBegin] + la2[rBegin] + pre
		pre, sum = sum / 10, sum % 10
		tmpNode := &ListNode{Val: sum, Next: re}
		re = tmpNode
		lBegin, rBegin = lBegin-1, rBegin-1
	}

	if lBegin >= 0 {
		re = addList(la1, re, &pre, lBegin)
	} 

	if rBegin >= 0 {
		re = addList(la2, re, &pre, rBegin)
	}

	if pre > 0 {
		tmpNode := &ListNode{Val: pre, Next: re}
		re = tmpNode
	}
	return re
}

func addList(list []int, slist *ListNode, pre *int, pos int) *ListNode {
	for pos >= 0 {
		sum := list[pos] + *pre
		*pre, sum = sum / 10, sum % 10
		tmpNode := &ListNode{Val: sum, Next: slist}
		slist = tmpNode
		pos--
	}
	return slist
}

func listToArray(lst *ListNode) []int {
	
	var re []int
	for lst != nil {
		re = append(re, lst.Val)
		lst = lst.Next
	}

	return re
}

func main() {
	value1 := &ListNode{1, nil}
	value2 := &ListNode{9, &ListNode{9, nil}}
	
	addTwoNumbers(value1, value2)
}