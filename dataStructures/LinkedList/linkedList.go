package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) Append(val int) *ListNode {
	if list == nil {
		return &ListNode{Val: val}
	}
	curr := list
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &ListNode{Val: val}
	return list
}

func (list *ListNode) Print() {
	if list == nil {
		return
	}
	curr := list
	for curr != nil {
		fmt.Printf("%v ", curr.Val)
		curr = curr.Next
	}
}

func (list *ListNode) Delete(val int) *ListNode {
	if list == nil {
		return nil
	}
	if list.Val == val {
		return list.Next
	}
	curr := list
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return list
}

func (list *ListNode) Prepend(val int) *ListNode {
	curr := &ListNode{Val: val}
	curr.Next = list
	return curr
}

func (list *ListNode) Search(val int) *ListNode {
	if list == nil {
		return nil
	}
	curr := list
	for curr != nil {
		if curr.Val == val {
			return curr
		}
		curr = curr.Next
	}
	return nil
}
