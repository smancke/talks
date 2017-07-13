package foo

import (
	"fmt"
	"testing"
)

func Test_Calculator(t *testing.T) {
	tests := []struct {
		op          string
		a           int
		b           int
		expected    int
		expectError bool
	}{
		{"+", 1, 1, 2, false},
		{"-", 1, 1, 0, false},
		{"*", 1, 1, 1, false},
		{"*", 2, 3, 6, false},
		{"/", 6, 2, 3, false},
		{"/", 6, 0, 0, false},
		{"foobar", 0, 0, 0, false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v %v %v == %v", test.a, test.op, test.b, test.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := Calc(test.op, test.a, test.b)
			if actual != test.expected {
				t.Logf("expected %v, but got: %v", test.expected, actual)
				t.Fail()
			}
			if test.expectError && err == nil {
				t.Logf("expected error, but got nil")
				t.Fail()
			}
			if !test.expectError && err != nil {
				t.Logf("no error expected, but got %v", err)
				t.Fail()
			}

		})
	}
}
