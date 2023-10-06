package main

import (
	"testing"
)

type TestCase struct {
	input []int
	expectedResult int
}

var testCases = []TestCase{
	TestCase{[]int{100,4,200,1,3,2},4},
	TestCase{[]int{0,3,7,2,5,8,4,6,0,1},9},
}

func TestLongestConsecutiveSequence(t *testing.T){
	for index, testCase := range testCases {
		testResult := LongestConsecutiveSequence(testCase.input)
		if(testResult != testCase.expectedResult) {
			t.Errorf("Test %v expected %v but got %v",
			index,
			testCase.expectedResult,
			testResult,
			)
		}
	}
}
