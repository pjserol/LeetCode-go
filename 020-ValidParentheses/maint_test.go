package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput bool
	}{
		{
			name:           "test happy path",
			input:          "()[]{}({[]})",
			expectedOutput: true,
		},
		{
			name:           "test wrong expression",
			input:          "()[]{}({[})",
			expectedOutput: false,
		},
		{
			name:           "test wrong first element expression",
			input:          "]{}({})",
			expectedOutput: false,
		},
	}
	for _, test := range tests {
		output := isValid(test.input)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected result %v, but got %v", test.name, test.expectedOutput, output)
		}
	}
}
