package add_two_numbers

import (
	"testing"
)

func TestAnswer7(t *testing.T) {
	TestAddTwoNumbers(t, Answer7)
}

func BenchmarkAnswer7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkAddTwoNumbers(b, Answer7)
	}
}
