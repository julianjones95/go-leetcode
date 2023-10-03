package main

func main(){

	example := []int{1,2,3,4}

	val := checkDuplicates(example)

	print(val)

}

func checkDuplicates (nums []int) bool {
	integerMap := make(map[int]int)

	for _, value := range nums {
		if integerMap[value] == 1 {
			return true;
		} else {
			integerMap[value] = 1;
		}
	}

	return false

}
