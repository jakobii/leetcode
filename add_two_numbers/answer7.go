package add_two_numbers

func Answer7(l1 *ListNode, l2 *ListNode) *ListNode {
	start := new(ListNode)
	cursor := start
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
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
		carry = v / 10
		cursor.Next = &ListNode{
			Val: v % 10,
		}
		cursor = cursor.Next
	}
	return start.Next
}
