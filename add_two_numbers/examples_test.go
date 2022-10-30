package add_two_numbers

import (
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList(1, 2, 3)

	if list.Val != 1 {
		t.Fatalf("want: 1, got: %v", list.Val)
	}
	if list.Next.Val != 2 {
		t.Fatalf("want: 2, got: %v", list.Next.Val)
	}
	if list.Next.Next.Val != 3 {
		t.Fatalf("want: 3, got: %v", list.Next.Next.Val)
	}
	if list.Next.Next.Next != nil {
		t.Fatalf("want: nil, got: %+v", list.Next.Next.Next)
	}
}

func TestListNode_AsSLice_small(t *testing.T) {
	got := NewSinglyLinkedList(1, 2, 3).AsSLice()
	want := []int{1, 2, 3}
	if !equal(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestListNode_AsSLice_large(t *testing.T) {
	got := NewSinglyLinkedList(9, 9, 9, 9, 9, 9, 9).AsSLice()
	want := []int{9, 9, 9, 9, 9, 9, 9}
	if !equal(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestNewSinglyLinkedListReversed(t *testing.T) {
	got := NewSinglyLinkedListReversed(1, 2, 3).AsSLice()
	want := []int{3, 2, 1}
	if !equal(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
