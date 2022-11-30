package problem0739

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i, t := range temperatures {
		// 当前元素比栈顶要大，更新结果，再去弹出栈顶元素
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			ans[topIndex] = i - topIndex
			stack = stack[:len(stack)-1]
		}

		// 当前元素比栈顶要小，直接入栈
		stack = append(stack, i)
	}

	return ans
}