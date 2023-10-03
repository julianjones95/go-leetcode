package main

import (
	"testing"
	"reflect"
)

type TestCase struct {
	inputArray []int
	inputTarget int
	result []int
}

var testCases = []TestCase{
	TestCase{[]int{0,1},1,[]int{0,1}},
	TestCase{[]int{3,3},6,[]int{0,1}},
	TestCase{[]int{2,7,11,15},9,[]int{0,1}},
	TestCase{[]int{3,2,4},6,[]int{1,2}},
}



func TestTwoSum(t *testing.T) {

	for index, testCase := range testCases {
		if !reflect.DeepEqual(testCase.result,findTargetSumInArrayOfNums(testCase.inputArray,testCase.inputTarget)) {
			t.Errorf("TestTwoSum test %v expected %v got %v",
				index,
				testCase.result,
			findTargetSumInArrayOfNums(testCase.inputArray,testCase.inputTarget),
			)

		}
	}
}
