package main

import (
	"testing"
	"reflect"
)

type TestCase struct {
	input []string
	expected [][]string
}

var testCases = []TestCase{
	TestCase{
		[]string{"eat","tea","tan","ate","nat","bat"},
		[][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}},
	},
	TestCase{
		[]string{""},
		[][]string{{""}},
	},
	TestCase{
		[]string{"a"},
		[][]string{{"a"}},
	},
}

func TestGroupAnagrams(t *testing.T) {

	for index, testCase := range testCases {
		funcResult := GroupAnagrams(testCase.input)
		if !reflect.DeepEqual(funcResult, testCase.expected) {
			t.Errorf("GroupAnagram test %v expected %v but got %v",
				index,
				testCase.expected,
				funcResult,
			)
		}
	}
}
