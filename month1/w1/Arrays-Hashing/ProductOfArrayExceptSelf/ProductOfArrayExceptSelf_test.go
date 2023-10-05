package main

import (
	"testing"
	"reflect"
)

type TestCase struct {
	input []int
	expectedResult []int
}

var testCases = []TestCase{
	TestCase{[]int{1,2,3,4},[]int{24,12,8,6}},
	TestCase{[]int{-1,1,0,-3,3},[]int{0,0,9,0,0}},
	TestCase{[]int{0,1,0,-3,3},[]int{0,0,0,0,0}},
	TestCase{[]int{-1,1,1,-3,0},[]int{0,0,0,0,3}},
	TestCase{[]int{1,1,1,1,1,1},[]int{1,1,1,1,1,1}},
	TestCase{[]int{2},[]int{0}},
	TestCase{[]int{-1},[]int{0}},

}

func TestProductOfArrayExceptSelf(t *testing.T) {
	for index, testCase := range testCases {
		testResult := ProductOfArrayExceptSelf(testCase.input)
		if !reflect.DeepEqual(testResult,testCase.expectedResult) {
			t.Errorf("TestProductOfArrayExceptSelf test %v expects %v got %v",
				index,
				testCase.expectedResult,
				testResult,
			)
		}
	} 
}
