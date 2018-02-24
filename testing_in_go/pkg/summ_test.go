package pkg

import (
	"testing"
)

// var tests = []struct {
// 	a, b, sum int
// }{
// 	{1, 2, 3},
// 	{1, 2, 3},
// 	{1, 2, 3},
// 	{1, 2, 3},
// 	{1, 2, 3},
// }

// TEST OMIT
func TestSums(t *testing.T) {
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			actual := Sum(tc.a, tc.b)
			if actual != tc.sum {
				t.Fail()
			}
		})
	}
}

// END OMIT
