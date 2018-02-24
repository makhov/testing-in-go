package pkg

import (
	"testing"
)

var tests = []struct {
	a, b, sum int
}{
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},
}

// TEST OMIT
func TestSum(t *testing.T) {
	for _, tc := range tests {
		actual := Sum(tc.a, tc.b)
		if actual != tc.sum {
			t.Fail()
		}
	}
}

func TestSum2(t *testing.T) {
	tests := append(tests, struct{ a, b, sum int }{1, 2, 4})
	for _, tc := range tests {
		actual := Sum(tc.a, tc.b)
		if actual != tc.sum {
			t.FailNow()
		}
	}
}

func TestSum3(t *testing.T) {
	for _, tc := range tests {
		actual := Sum(tc.a, tc.b)
		if actual != tc.sum {
			t.Fail()
		}
	}
}

func TestSum4(t *testing.T) {
	for _, tc := range tests {
		actual := Sum(tc.a, tc.b)
		if actual != tc.sum {
			t.Fail()
		}
	}
}
func TestSum5(t *testing.T) {
	for _, tc := range tests {
		// ¯\_(ツ)_/¯
		actual := Sum(tc.a, tc.b)
		if actual != tc.sum {
			t.Fail()
		}
	}
}
func TestSum6(t *testing.T) {
	for _, tc := range tests {
		actual := Sum(tc.a, tc.b)
		if actual != tc.sum {
			t.Fail()
		}
	}
}
func TestSum7(t *testing.T) {
	for _, tc := range tests {
		actual := Sum(tc.a, tc.b)
		if actual != tc.sum {
			t.Fail()
		}
	}
}

// END OMIT
