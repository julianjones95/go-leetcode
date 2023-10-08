package main

import "fmt"

func main(){

}

type Stack struct {
	items []rune
}

func (stack *Stack) push(item rune) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) pop() {
	stack.items = stack.items[:len(stack.items)-1]
}

func (stack *Stack) isEmpty() bool {
	length := len(stack.items);
	return length == 0
}

func (stack *Stack) top() rune {
	return stack.items[len(stack.items)-1]
}

func ValidParentheses(s string) bool {

	var currentParentheses Stack = Stack{};

	for _, value := range s {
		if value == '{' || value == '[' || value == '(' {
			currentParentheses.push(value)
		} else if value == '}' && currentParentheses.top() == '{' {
			currentParentheses.pop()
		} else if value == ')' && currentParentheses.top() == '(' {
			currentParentheses.pop()
		} else if value == ']' && currentParentheses.top() == '[' {
			currentParentheses.pop()
		}
	}

	fmt.Printf("\n parens after %c", currentParentheses.items)

	return currentParentheses.isEmpty();
}
