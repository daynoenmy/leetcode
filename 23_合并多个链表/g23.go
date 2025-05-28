package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			return head.Next
		}
		if l2 == nil {
			p.Next = l1
			return head.Next
		}
		if l1.Val <= l2.Val {
			t := l1
			l1 = l1.Next
			p.Next = t
			p = p.Next
		} else {
			t := l2
			l2 = l2.Next
			p.Next = t
			p = p.Next
		}
	}
	return head.Next

}
func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}
func mergeKLists(lists []*ListNode) *ListNode {
	var lt *ListNode = nil

	for i := 0; i < len(lists); i++ {
		lt = mergeTwoLists(lt, lists[i])
	}
	return lt
}
func main() {

}
