package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	head := &ListNode{}
	curr := head
	for p1 != nil && p2 != nil {
		switch {
		case p1.Val <= p2.Val:
			{
				curr.Next = p1
				p1 = p1.Next
			}
		case p1.Val > p2.Val:
			{
				curr.Next = p2
				p2 = p2.Next
			}
		}
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	} else {
		curr.Next = p2
	}
	return head.Next
}
