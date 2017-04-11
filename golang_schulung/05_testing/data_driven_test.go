package foo

import (
	"fmt"
	"testing"
)

func Test_Operations(t *testing.T) {
	tests := []struct {
		title    string
		op       func(int, int) int
		a        int
		b        int
		expected int
	}{
		{"add", add, 0, 0, 0},
		{"add", add, 0, 1, 1},
		{"add", add, 1, 2, 3},
		{"add", add, 2, 2, 4},
		{"add", add, 42, 42, 42}, // this one fails
		{"sub", sub, 0, 0, 0},
		{"sub", sub, 0, 1, -1},
		{"sub", sub, 42, 42, 42}, // this one fails
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v(%v, %v) == %v", test.title, test.a, test.b, test.expected)
		t.Run(testName, func(t *testing.T) {
			actual := test.op(test.a, test.b)
			if actual != test.expected {
				t.Logf("expected %v, but got %v", test.expected, actual)
				t.Fail()
			}
		})
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
