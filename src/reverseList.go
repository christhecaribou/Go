/*
CHRIS FELLI, 2019
Reverse a linked list
*/

//Package gocode ...
package gocode

//ListNode : A linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// If head is nil
	if nil == head || nil == head.Next {
		return head
	}

	var pre *ListNode
	cur := head
	for nil != cur {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	head = pre

	return head
}
