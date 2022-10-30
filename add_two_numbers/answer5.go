package add_two_numbers

func Answer5(l1 *ListNode, l2 *ListNode) *ListNode {
	return recurs(l1, l2, 0)
}

func recurs(l1, l2 *ListNode, carry int) *ListNode {
	if l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v := v1 + v2 + carry
		carry = 0
		if v > 9 {
			carry = v / 10
			v = v % 10
		}
		return &ListNode{
			Val:  v,
			Next: recurs(l1, l2, carry),
		}
	}
	if carry > 0 {
		return &ListNode{
			Val: carry,
		}
	}
	return nil
}
