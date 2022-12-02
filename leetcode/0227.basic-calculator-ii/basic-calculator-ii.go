package problem0227

import (
	"strconv"
)

func calculate(s string) int {
	s += " "
	tokens := []string{}	// 存放后缀表达式的栈
	ops := []rune{}			// 存放操作符的栈
	number := ""			// 存放字符串中的数字
	for _, char := range s {
		if (char >= '0' && char <= '9') {	// 字符串中数字的拼接
			number += string(char)
			continue
		} else {
			if number != "" {	// 字符串中数字的处理——数字放入后缀表达式栈
				tokens = append(tokens, number)
				number = ""
			}

			if char == ' ' {	// 字符串中空格的处理
				continue
			}

			curRank := getRank(char)	// 字符串中操作符的处理
			for (len(ops) > 0 && getRank(ops[len(ops) - 1]) >= curRank) {
				// 操作符栈中栈顶元素优先级比当前元素优先级高——弹出操作符栈的栈顶元素，并将其入栈到后缀表达式栈
				tokens = append(tokens, string(ops[len(ops) - 1]))
				ops = ops[:len(ops) - 1]
			}
			// 操作符栈中栈顶元素优先级比当前元素优先级低——当前元素入栈操作符栈
			ops = append(ops, char)
		}
	}

	// 操作符栈中所有元素弹出，并入栈到后缀表达式栈
	for (len(ops) > 0) {
		tokens = append(tokens, string(ops[len(ops) - 1]))
		ops = ops[:len(ops) - 1]
	}

	return evalRPN(tokens)
}

func getRank(char rune) int {
	if (char == '+' || char == '-') {
		return 1
	}
	if (char == '*' || char == '/') {
		return 2
	}
	return 0
}

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