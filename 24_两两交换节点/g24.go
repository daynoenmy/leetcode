package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dumHead := &ListNode{}
	dumHead.Next = head
	temp := dumHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1, node2 := temp.Next, temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dumHead.Next
}
