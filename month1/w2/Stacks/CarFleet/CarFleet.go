package main

import (
	"fmt"
	"sort"
)

func main() {

}

type Stack struct{
	items []rune
}

func (stack *Stack) push(item rune) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) pop() {
	stack.items = stack.items[:len(stack.items)-1]
}

func (stack *Stack) isEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack) top() rune {
	return stack.items[len(stack.items)-1]
}

func CarFleet (target int, position []int, speed []int) int {

	var carMap = make(map[int]int)

	for i := range position {
		carMap[position[i]] = speed[i]
	}

	keys := make([]int, 0, len(carMap))


	for i := range position {
		keys = append(keys,i)
	}

	
	sort.Ints(keys)

	fmt.Println(keys)


	return 0;
}
