package main

import "leetcode/List"

type ListNode = List.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre *ListNode = head
	var cur *ListNode = nil
	var trail *ListNode = head
	for i := 0; i < n; i++ {
		trail = trail.Next
	}
	if trail == nil {
		return head.Next
	}
	for trail != nil {
		cur = pre
		trail = trail.Next
		pre = pre.Next
	}
	cur.Next = pre.Next
	return head
}
func main() {

}
