package main

func main(){

}

func TopKFrequentElements(nums []int, k int) []int{

	mostFrequentMap := make(map[int]int)

	var output = []int{};
	
	for i:=0; i<k; i++ {

		currentMaxValue := 0;
		currentMaxIndex := 0;
		var tempNums = []int{}

		for _, value := range nums {
			mostFrequentMap[value] = mostFrequentMap[value] + 1;
		}

		for key := range mostFrequentMap {
			if mostFrequentMap[key] > currentMaxValue {
				currentMaxValue = mostFrequentMap[key];
				currentMaxIndex = key;
			}
		}

		output = append(output, currentMaxIndex)

		for _, val := range nums {
			if val != currentMaxIndex {
				tempNums = append(tempNums, val)
			}
		}

		nums = tempNums;
	}

	return output;
}
