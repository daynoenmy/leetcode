package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Init(array []int) *ListNode {
	var head, p *ListNode
	p = new(ListNode)
	p.Val = array[0]
	head = p
	p.Next = nil
	for i := 1; i < len(array); i++ {
		t := new(ListNode)
		t.Val = array[i]
		t.Next = nil
		head.Next = t
		head = t
	}
	return p
}
func (l *ListNode) printList() {
	for head := l; head != nil; head = head.Next {
		fmt.Printf("%d ", head.Val)
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0
	isFirst := true
	var head, p *ListNode
	for l1 != nil || l2 != nil {
		l := new(ListNode)
		l.Next = nil
		if isFirst {
			head = l
			p = head
			isFirst = false
		} else {
			p.Next = l
			p = l
		}
		if l1 == nil {
			sum := l2.Val + carry
			carry = sum / 10
			sum = sum % 10
			l.Val = sum
			l2 = l2.Next
		} else if l2 == nil {
			sum := l1.Val + carry
			carry = sum / 10
			sum = sum % 10
			l.Val = sum
			l1 = l1.Next
		} else {
			sum := l1.Val + l2.Val + carry
			carry = sum / 10
			sum = sum % 10
			l.Val = sum
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	if carry > 0 {
		l := new(ListNode)
		l.Next = nil
		l.Val = 1
		p.Next = l
	}
	return head
}
func main() {
	l1 := Init([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := Init([]int{9, 9, 9, 9})
	l3 := addTwoNumbers(l1, l2)
	l3.printList()
}
