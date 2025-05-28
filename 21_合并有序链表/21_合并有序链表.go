package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l := new(ListNode)
	head := l
	l.Next = nil
	for list1 != nil || list2 != nil {
		if list1 == nil {
			for list2 != nil {
				node := new(ListNode)
				node.Val = list1.Val
				l.Next = node
				node.Next = nil
				l = l.Next
				list2 = list2.Next
			}
			return head.Next
		}
		if list2 == nil {
			for list1 != nil {
				node := new(ListNode)
				node.Val = list1.Val
				l.Next = node
				node.Next = nil
				l = l.Next
				list1 = list1.Next
			}
			return head.Next
		}
		if list1.Val <= list2.Val {
			node := new(ListNode)
			node.Val = list1.Val
			l.Next = node
			node.Next = nil
			l = l.Next
			list1 = list1.Next
		} else {
			node := new(ListNode)
			node.Val = list1.Val
			l.Next = node
			node.Next = nil
			l = l.Next
			list2 = list2.Next
		}
	}
	return head.Next
}
