package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	cases := []struct {
		name     string
		in       int
		expected int
	}{
		{
			name:     "factorial of 5 is 120",
			in:       5,
			expected: 120,
		},
		{
			name:     "factorial of 0 is 1",
			in:       0,
			expected: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := factorial(c.in)
			if got != c.expected {
				t.Errorf("want: %v, got: %v", c.expected, got)
			}
		})
	}
}
