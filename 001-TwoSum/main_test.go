package main

import (
	"reflect"
	"testing"
)

func Test_twoSumOnePass(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		target         int
		expectedOutput []int
	}{
		{
			name:           "test happy path",
			nums:           []int{2, 7, 11, 15},
			target:         9,
			expectedOutput: []int{0, 1},
		},
		{
			name:   "test no result ",
			nums:   []int{2, 7, 11, 15},
			target: 100,
		},
	}
	for _, test := range tests {
		output := twoSumOnePass(test.nums, test.target)

		if !reflect.DeepEqual(output, test.expectedOutput) {
			t.Errorf("for %s, expected result %v, but got %v", test.name, test.expectedOutput, output)
		}
	}
}

func Test_twoSumTwoPass(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		target         int
		expectedOutput []int
	}{
		{
			name:           "test happy path",
			nums:           []int{2, 7, 11, 15},
			target:         9,
			expectedOutput: []int{0, 1},
		},
		{
			name:   "test no result ",
			nums:   []int{2, 7, 11, 15},
			target: 100,
		},
	}
	for _, test := range tests {
		output := twoSumTwoPass(test.nums, test.target)

		if !reflect.DeepEqual(output, test.expectedOutput) {
			t.Errorf("for %s, expected result %v, but got %v", test.name, test.expectedOutput, output)
		}
	}
}

func Test_twoSumBruteForce(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		target         int
		expectedOutput []int
	}{
		{
			name:           "test happy path",
			nums:           []int{2, 7, 11, 15},
			target:         9,
			expectedOutput: []int{0, 1},
		},
		{
			name:   "test no result ",
			nums:   []int{2, 7, 11, 15},
			target: 100,
		},
	}
	for _, test := range tests {
		output := twoSumBruteForce(test.nums, test.target)

		if !reflect.DeepEqual(output, test.expectedOutput) {
			t.Errorf("for %s, expected result %v, but got %v", test.name, test.expectedOutput, output)
		}
	}
}
