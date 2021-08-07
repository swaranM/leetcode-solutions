package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, current, next *ListNode
	current = head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	head = prev

	return head
}
