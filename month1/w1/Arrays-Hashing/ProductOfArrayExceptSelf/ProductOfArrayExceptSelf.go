package main

func main(){
}

func ProductOfArrayExceptSelf(nums []int) []int {

	if(len(nums)==1){
		return []int{0};
	}

	product := GetArrayProduct(nums)

	var outputArray []int
	var currentVal int;

	for index, value := range nums {
		if(value == 0) {
			currentVal = GetArrayProductWithoutValue(nums,index);
			outputArray = append(outputArray, currentVal)
		} else {
			currentVal = product / value
			outputArray = append(outputArray, currentVal)
		}
	}
	
	return outputArray;
}

func GetArrayProduct(nums []int) int {
	product := 1;
	for _, value := range nums {
		product = product * value;
	}
	return product
}

func GetArrayProductWithoutValue(nums []int, currentIndex int) int {
	product := 1;
	for index, value := range nums {
		if(index != currentIndex){
			product = product * value;
		}
	}
	return product
}
