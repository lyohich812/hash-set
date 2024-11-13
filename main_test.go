package main

import (
	"fmt"
	"testing"
)

func TestFindFirstDuplicate(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect interface{}
	}{
		{
			input:  []int{2, 5, 1, 2, 3, 5, 1, 2, 4},
			expect: 2,
		},
		{
			input:  []string{"jon", "5", "doe", "jon", "3", "5"},
			expect: "jon",
		},
		{
			input:  []int{1, 2, 3, 4, 5, 5},
			expect: 5,
		},
		{
			input:  []int{1},
			expect: nil,
		},
		{
			input:  3,
			expect: nil,
		},
		{
			input:  "test",
			expect: nil,
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		result := findFirstDuplicate(test.input)

		if result != test.expect {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail`, test.input, test.expect, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.input, test.expect, result)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)

}
