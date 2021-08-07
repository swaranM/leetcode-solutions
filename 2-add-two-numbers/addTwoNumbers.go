package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummyHead := &ListNode{Val: 0}

	l3 := dummyHead

	for l1 != nil || l2 != nil {
		var l1val int
		var l2val int

		if l1 != nil {
			l1val = l1.Val
		}

		if l2 != nil {
			l2val = l2.Val
		}

		sum := l1val + l2val + carry
		carry = sum / 10
		lastDigit := sum % 10

		newNode := &ListNode{Val: lastDigit}
		l3.Next = newNode

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		l3 = l3.Next
	}
	if carry > 0 {
		newNode := &ListNode{Val: carry}
		l3.Next = newNode
		l3 = l3.Next
	}

	return dummyHead.Next
}
