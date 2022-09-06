package problem0283

func moveZeroes(nums []int)  {
    cnt := 0

    // 过滤器思想——过滤大于0的元素到数组的前半部分
    for _, num := range nums {
        if (num != 0) {
            nums[cnt] = num
            cnt++
        }
    }

    // 数组的后半部分补0
    for cnt < len(nums) {
        nums[cnt] = 0
        cnt++
    }
}