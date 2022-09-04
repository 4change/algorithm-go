package problem0026

func removeDuplicates(nums []int) int {
    cnt := 0

    // 过滤器思想——过滤当前元素值不等于前一元素值的元素，并放入过滤器
    for i, v := range nums {
        if (i == 0 || nums[i] != nums[i-1]) {
            nums[cnt] = v
            cnt++
        }
    }

    return cnt
}