package main

import "testing"

type TestCase struct {
	string1 string
	string2 string
	result bool
} 

var testCases = []TestCase{
	TestCase{"","",true},
	TestCase{"a","a",true},
	TestCase{"a","z",false},
	TestCase{"abcdef", "fedcba",true},
	TestCase{"abc","xyz",false},
	TestCase{"anagram", "naargam", true},
}

func TestValidAnagram(t *testing.T) {

	for index, testCase := range testCases {

		testResult := ValidAnagram(testCase.string1,testCase.string2)

		if testResult != testCase.result {
			t.Errorf("Test %v expected: %v but got %v", 
				index,
				testCase.result,
				ValidAnagram(testCase.string1,testCase.string2),
			)
		}
	}
}
