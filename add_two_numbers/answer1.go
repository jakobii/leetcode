package add_two_numbers

func Answer1(l1 *ListNode, l2 *ListNode) *ListNode {

	s1 := l1.AsSLice()
	s2 := l2.AsSLice()

	s := addSlice(s1, s2)
	return NewSinglyLinkedList(s...)

}

func addSlice(a []int, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	out := make([]int, 0, len(a)+len(b))

	carry := 0
	aCurs := 0
	bCurs := 0

	for {
		var aN, bN int
		if aCurs < len(a) {
			aN = a[aCurs]
		}
		if bCurs < len(b) {
			bN = b[bCurs]
		}

		n := aN + bN + carry

		if n > 9 {
			rem := n % 10
			carry = (n - rem) / 10
			out = append(out, rem)
		} else {
			out = append(out, n)
			carry = 0
		}

		aCurs++
		bCurs++
		if aCurs >= len(a) && bCurs >= len(b) {
			if carry > 0 {
				out = append(out, carry)
			}
			break
		}
	}
	return out
}
