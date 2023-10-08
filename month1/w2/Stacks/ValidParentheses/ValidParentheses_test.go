package main

import "testing"

type TestCase struct {
	input string
	expectedResult bool
}

var testCases = []TestCase{
	TestCase{"()", true},
	TestCase{"()[]{}", true},
	TestCase{"(]", false},
}

func TestValidParentheses(t *testing.T) {

	for index, testCase := range testCases {

		actualValue := ValidParentheses(testCase.input)

		if actualValue != testCase.expectedResult {
			t.Errorf("ValidParentheses test %v failed expected %v",
				index,
				testCase.expectedResult,
			)
		}
	}
}
