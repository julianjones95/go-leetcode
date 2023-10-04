package main

import (
	"testing"
	"reflect"
)

type TestCase struct {
	input []int
	kValue int
	result []int
}

var testCases = []TestCase{
	TestCase{[]int{1,1,1,2,2,3},2,[]int{1,2}},
	TestCase{[]int{1},1,[]int{1}},


}

func TestTopKFrequentElements(t *testing.T){

	for index, testCase := range testCases {
		actual := TopKFrequentElements(testCase.input,testCase.kValue)
		if !reflect.DeepEqual(actual,testCase.result) {
			t.Errorf("Test %v failed got %v expected %v",
				index,
				actual,
				testCase.result,
			)
		}
	}	
}
