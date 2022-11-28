package problem0150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int

	for _, token := range tokens {
		val, err := strconv.Atoi(token)

		if err == nil {		// 数字则入栈
			stack = append(stack, val)
		} else {	// 运算符则弹出两个栈顶元素并计算
			y := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			x := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			z := cal(x, y, token)
			stack = append(stack, z)
		}
	}

	return stack[len(stack) - 1]
}

func cal(x, y int, opr string) int {
	if opr == "+" {
		return x + y
	} 
    if opr == "-" {
		return x - y
	} 
    if opr == "*" {
		return x * y
	} 
    if opr == "/" {
		return x / y
	} 
	
    return 0
}