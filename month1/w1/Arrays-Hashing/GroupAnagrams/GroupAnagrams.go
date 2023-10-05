package main

import "sort"

func main(){
	print(checkValidAnagram("tan","nat"))
}

func GroupAnagrams(strs []string) [][]string {
	
	var groupedAnagrams = [][]string{{strs[0]}}

	for i:=1; i<len(strs); i++ {

		var endOfSet = len(groupedAnagrams)-1;

		for j:=0; j<len(groupedAnagrams); j++ {

			if(checkValidAnagram(strs[i], groupedAnagrams[j][0])){
				groupedAnagrams[j] = append(groupedAnagrams[j],strs[i])
				break;
			}

			if(j==endOfSet){
				groupedAnagrams = append(groupedAnagrams, []string{strs[i]})
				break;
			}
		}
	}

	sortSliceByNumElements(groupedAnagrams)

	sortSliceAlphabeticallyByFirstLetter(groupedAnagrams)

	return groupedAnagrams;
}

func sortSliceByNumElements(nums [][]string) [][]string{
	sort.SliceStable(nums, func(i, j int) bool {
		return len(nums[i]) < len(nums[j]) 	
	})

	return nums
}

func sortSliceByFirstLetter(nums [][]string) [][]string {
	for i:=0; i<len(nums); i++ {
		sort.Slice(nums[i], func(k, l int) bool { 
			return nums[i][k] < nums[i][l]
		})
	}

	return nums
}

func checkValidAnagram (str1 string, str2 string) bool {

	anagramMap := make(map[rune]int)

	for _, char := range str1 {
		anagramMap[char] = anagramMap[char] + 1;
	}
	
	for _, char := range str2 {

		_, keyExists := anagramMap[char]

		if keyExists {
			anagramMap[char] = anagramMap[char] - 1;
		}
	}

	anagramSum := 0;

	for _, val := range anagramMap {
		anagramSum = anagramSum + val;
	}

	return anagramSum == 0;
}

