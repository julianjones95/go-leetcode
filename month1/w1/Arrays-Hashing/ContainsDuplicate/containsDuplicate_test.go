package main

import "testing"

type TestCase struct {
	input []int
	result bool
}

var testCases = []TestCase{
	TestCase{[]int{},false},
	TestCase{[]int{0},false},
	TestCase{[]int{1},false},
	TestCase{[]int{0,0},true},
	TestCase{[]int{1,1},true},
	TestCase{[]int{1,2,3,4,5},false},
	TestCase{[]int{1,2,3,4,5,1},true},
	TestCase{[]int{1,2,3,4,5,6},false},
	TestCase{[]int{1,2,3,4,5,6,0},false},
	TestCase{[]int{1,1,1,1,1,1,1,1,1,1,1},true},
}

func TestContainsDuplicate(t *testing.T) {

	for index, testCase := range testCases {
		resultOfTest := checkDuplicates(testCase.input)

		if(resultOfTest != testCase.result){
			t.Errorf("Expects checkDuplicates test: %v to be %v",
				index, 
				testCase.result,
			)
		} 
	}
}
