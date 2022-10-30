package add_two_numbers

import (
	"testing"
)

func TestAnswer6(t *testing.T) {
	TestAddTwoNumbers(t, Answer6)
}

func BenchmarkAnswer6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkAddTwoNumbers(b, Answer6)
	}
}
