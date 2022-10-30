package add_two_numbers

import (
	"testing"
)

func TestAnswer5(t *testing.T) {
	TestAddTwoNumbers(t, Answer5)
}

func BenchmarkAnswer5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkAddTwoNumbers(b, Answer5)
	}
}
