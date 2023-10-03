package main

func main(){

	nums := []int{1,2,3,4,5,6}
	target := 3

	findTargetSumInArrayOfNums(nums, target);
}

func findTargetSumInArrayOfNums(nums []int, target int) []int{

	mapOfDifferences := make(map[int]int)

	for index, value := range nums {
		_, differenceExists := mapOfDifferences[value]
		if(differenceExists) {
			return []int{mapOfDifferences[value], index}
		}

		mapOfDifferences[target-value] = index;
	}	

	return []int{0,0}
}
