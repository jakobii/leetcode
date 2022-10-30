package add_two_numbers

import (
	"fmt"
	"testing"
)

type AddTwoNumbersFunc = func(l1 *ListNode, l2 *ListNode) *ListNode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) AsSLice() []int {
	out := make([]int, 0)
	c := l
	for {
		out = append(out, c.Val)
		if c.Next == nil {
			break
		}
		c = c.Next
	}
	return out
}

func NewSinglyLinkedList(numbers ...int) *ListNode {
	if len(numbers) < 1 {
		return nil
	}
	var prev *ListNode
	for i := len(numbers) - 1; i > -1; i-- {
		ln := &ListNode{
			Val:  numbers[i],
			Next: prev,
		}
		prev = ln
	}
	return prev
}

func NewSinglyLinkedListReversed(numbers ...int) *ListNode {
	if len(numbers) < 1 {
		return nil
	}
	var prev *ListNode
	for i := 0; i < len(numbers); i++ {
		ln := &ListNode{
			Val:  numbers[i],
			Next: prev,
		}
		prev = ln
	}
	return prev
}

func EqSinglyLinkedList(want *ListNode, got *ListNode) bool {
	s1 := want.AsSLice()
	s2 := got.AsSLice()
	return equal(s1, s2)

}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestAddTwoNumbers
func TestAddTwoNumbers(t *testing.T, addTwoNumbers AddTwoNumbersFunc) {
	for i, x := range examples {
		t.Run(fmt.Sprintf("Example %v", i), func(t *testing.T) {
			got := addTwoNumbers(x.l1, x.l2)
			if !equal(x.want, got.AsSLice()) {
				t.Fatalf(
					"\nl1: %+v\nl2: %+v\nwant: %+v\ngot: %+v",
					x.l1.AsSLice(),
					x.l2.AsSLice(),
					x.want,
					got.AsSLice(),
				)
			}
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B, addTwoNumbers AddTwoNumbersFunc) {
	for _, x := range examples {
		got := addTwoNumbers(x.l1, x.l2)
		if !equal(x.want, got.AsSLice()) {
			b.Fatalf(
				"\nl1: %+v\nl2: %+v\nwant: %+v\ngot: %+v",
				x.l1.AsSLice(),
				x.l2.AsSLice(),
				x.want,
				got.AsSLice(),
			)
		}
	}
}

var examples = []struct {
	l1   *ListNode
	l2   *ListNode
	want []int
}{
	{
		l1:   NewSinglyLinkedListReversed(2, 4, 3),
		l2:   NewSinglyLinkedListReversed(5, 6, 4),
		want: []int{7, 0, 8},
	},
	{
		l1:   NewSinglyLinkedListReversed(0),
		l2:   NewSinglyLinkedListReversed(0),
		want: []int{0},
	},
	{
		l1:   NewSinglyLinkedListReversed(9, 9, 9, 9, 9, 9, 9),
		l2:   NewSinglyLinkedListReversed(9, 9, 9, 9),
		want: []int{8, 9, 9, 9, 0, 0, 0, 1},
	},
	{
		l1:   NewSinglyLinkedListReversed(1, 1, 1),
		l2:   NewSinglyLinkedListReversed(1, 1, 1),
		want: []int{2, 2, 2},
	},
}
