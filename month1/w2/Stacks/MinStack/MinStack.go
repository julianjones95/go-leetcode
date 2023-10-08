package main

import "fmt"

type Stack struct {
	items []int
}

func main() {

	var stack = Constructor()
	
	stack.push(1)

	fmt.Println(stack.items)
}

func (stack *Stack) push(item int) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack) pop() {
	stackLength := len(stack.items)
	stack.items = stack.items[:stackLength -1]
}

func Constructor() Stack {
	return Stack{}
}
