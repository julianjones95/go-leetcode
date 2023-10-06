package main

import "sort"

func main(){
}

func LongestConsecutiveSequence(nums []int) int {

	maxSequence := 0;

	// Can this be done without sorting? currently runs nlog(n).
	sort.Ints(nums)

	previousValue := 0;
	currentSequence := 0;
	for _, value := range nums {
		if(previousValue == value -1){
			currentSequence = currentSequence + 1;
			previousValue = value;
		} else {
			currentSequence = 1;
			previousValue = value;
		}

		if(currentSequence > maxSequence) {
			maxSequence = currentSequence;
		}
	}

	return maxSequence;
}
