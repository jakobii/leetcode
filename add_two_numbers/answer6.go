package add_two_numbers

// this c/p from leetcode to prove that my code is better.

func getNumber(l *ListNode) []int {
	n := []int{l.Val}

	for l.Next != nil {
		l = l.Next
		n = append([]int{l.Val}, n...)
	}

	return n
}

// Answer6 was c/p from leet code. its really not that good. but somehow it got a better score then
// my submission "Answer5()". so im adding it here to benchmark it.
func Answer6(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := getNumber(l1)
	n2 := getNumber(l2)

	if len(n1) != len(n2) {
		if len(n1) > len(n2) {
			n1, n2 = n2, n1
		}
		diff := len(n2) - len(n1)
		zeroes := []int{}
		for i := 0; i < diff; i++ {
			zeroes = append(zeroes, 0)
		}
		n1 = append(zeroes, n1...)
	}

	node := ListNode{}
	list := &node
	n := &node

	first := true
	carry := 0
	for i := len(n1) - 1; i >= 0; i-- {
		digit := n1[i] + n2[i] + carry
		if digit > 9 {
			digit -= 10
			carry = 1
		} else {
			carry = 0
		}
		if first {
			n.Val = digit
			first = false
		} else {
			next := ListNode{Val: digit}
			n.Next = &next
			n = &next
		}
	}
	if carry == 1 {
		next := ListNode{Val: 1}
		n.Next = &next
	}

	return list
}
