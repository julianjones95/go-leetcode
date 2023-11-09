package main

import "testing"

type TestCase struct {
	inputTarget int
	inputPosition []int
	inputSpeed []int
	expectedResult int
}

var testCases = []TestCase{
	TestCase{12, []int{10,8,0,5,3}, []int{2,4,1,1,3}, 3},
	TestCase{10, []int{3}, []int{3}, 1},
	TestCase{100, []int{0,2,4}, []int{4,2,1}, 1},
}


func TestCarFleet(t *testing.T){

	for index, testCase := range testCases {

		actualValue := CarFleet(
			testCase.inputTarget, 
			testCase.inputPosition,
			testCase.inputSpeed,
		)

		if actualValue != testCase.expectedResult {
			t.Errorf("CarFleet test %v failed expected %v,",
				index,
				testCase.expectedResult,
			)
		}

	}

}
