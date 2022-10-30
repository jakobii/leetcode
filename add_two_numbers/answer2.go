package add_two_numbers

func Answer2(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersRecurse(l1, l2, 0)
}

func addTwoNumbersRecurse(l1, l2 *ListNode, carry int) *ListNode {
	var val, val1, val2, carry2 int
	var next1, next2 *ListNode
	if l1 != nil {
		next1 = l1.Next
		val1 = l1.Val
	}
	if l2 != nil {
		next2 = l2.Next
		val2 = l2.Val
	}
	val = val1 + val2 + carry
	if val > 9 {
		rem := val % 10
		carry2 = val / 10
		val = rem
	}

	next := &ListNode{
		Val: val,
	}

	if next1 == nil && next2 == nil {
		if carry2 > 0 {
			next.Next = &ListNode{
				Val: carry2,
			}
		}
	} else {
		next.Next = addTwoNumbersRecurse(next1, next2, carry2)
	}

	return next
}
