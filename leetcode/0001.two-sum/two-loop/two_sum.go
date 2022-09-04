package two_hash

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        num1 := nums[i]
        num2 := target - num1
        for j := i+1; j < len(nums); j++ {
            if nums[j] == num2 {
                return []int{i, j}
            }
        }
    }

    return nil
}
