package problem0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	hashmap := map[int]int{}	// key: nums2元素值；value: nums2中下一个更大元素
	stack := make([]int, 0)		// 单调栈（单调递减）

	for i, t := range nums2 {
		// 当前元素值大于栈顶元素值，更新哈希表，再去弹出栈顶元素
		for len(stack) > 0 && t > nums2[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			hashmap[nums2[topIndex]] = t
			stack = stack[:len(stack)-1]
		}

		// 当前元素值小于等于栈顶元素值，直接入栈
		stack = append(stack, i)
	}

	// 根据哈希表确定结果数组
	for i := range nums1 {
		if val, ok := hashmap[nums1[i]]; ok {
			res[i] = val
		} else {
			res[i] = -1
		}
	}

	return res
}