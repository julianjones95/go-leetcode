package main


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

	return groupedAnagrams;
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

