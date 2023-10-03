package main

func main(){
}

func ValidAnagram(s string, t string) bool {

	AnagramMap := make(map[rune]int)

	for _, char := range s {
		AnagramMap[char] = AnagramMap[char] + 1;
	}

	for _, char := range t {

		_, KeyExists := AnagramMap[char]
		if KeyExists {
			AnagramMap[char] = AnagramMap[char] - 1;
		}
	}

	mapSum := 0;

	for _, value := range AnagramMap {
		mapSum = mapSum + value;
	}

	var isThisAnAnagram bool = mapSum == 0


	return isThisAnAnagram;
}
