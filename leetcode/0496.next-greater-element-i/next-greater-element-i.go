package problem0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	hashmap := map[int]int{}	// key: nums2元素值；value: 下一个更大元素
	stack := make([]int, 0)

	for i, t := range nums2 {
		// 当前元素比栈顶要大，更新结果，再去弹出栈顶元素
		for len(stack) > 0 && t > nums2[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			hashmap[nums2[topIndex]] = t
			stack = stack[:len(stack)-1]
		}

		// 当前元素比栈顶要小，直接入栈
		stack = append(stack, i)
	}

	for i := range nums1 {
		if val, ok := hashmap[nums1[i]]; ok {
			ans[i] = val
		} else {
			ans[i] = -1
		}
	}

	return ans
}