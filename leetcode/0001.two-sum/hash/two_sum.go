package hash

func twoSum(nums []int, target int) []int {
	// map key: 数组值, value: 数组下标
	index := make(map[int]int, len(nums))

	for i, v := range nums {
		// map中查找到对应元素
		if j, ok := index[target-v]; ok {
			return []int{j, i}
		}

		// map中未查找到对应元素，元素值入map
		index[v] = i
	}

	return nil
}
