package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"([](){([])})", 0},
		{"()[]}", 5},
		{"{{[()]]", 7},
		{"{{{[][][]", 3},
		{"{*{{}", 3},
		{"[[*", 2},
		{"{*}", 0},
		{"{{", 2},
		{"{}", 0},
		{"", 0},
		{"}", 1},
		{"*{}", 0},
		{"{{{**[][][]", 3},
		{"foo(bar[i)", 10},
	}

	for _, test := range tests {
		result := check(test.input)
		if result != test.expected {
			t.Errorf("check(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
