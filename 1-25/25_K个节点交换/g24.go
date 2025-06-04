package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{}
	hair.Next = head
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}
