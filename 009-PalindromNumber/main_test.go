package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedOutput bool
	}{
		{
			name:           "test happy path",
			input:          1221,
			expectedOutput: true,
		},
		{
			name:           "test happy path with odd number",
			input:          12321,
			expectedOutput: true,
		},
		{
			name:           "test happy path with zero value",
			input:          0,
			expectedOutput: true,
		},
		{
			name:           "test happy path with big number",
			input:          1234567899987654321,
			expectedOutput: true,
		},
		{
			name:           "test with number not palindrome",
			input:          123,
			expectedOutput: false,
		},
		{
			name:           "test with negative number",
			input:          -1221,
			expectedOutput: false,
		},
		{
			name:           "test with big number",
			input:          1234569999987654321,
			expectedOutput: false,
		},
	}
	for _, test := range tests {
		output := isPalindrome(test.input)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected result %v, but got %v", test.name, test.expectedOutput, output)
		}
	}
}
