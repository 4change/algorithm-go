package problem0088

// 解法1：从前往后遍历，使用到临时数组 res 存放中间结果
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func merge_1(nums1 []int, m int, nums2 []int, n int)  {
    var res []int       // 新数组用于存储比较结果
    var i = 0
    var j = 0

    for (i < m || j < n) {
        if (j >= n || (i < m && nums1[i] <= nums2[j])) {    // 注意边界条件的判定
            res = append(res, nums1[i])
            i++
        } else {
            res = append(res, nums2[j])
            j++
        }
    }

    for i, num := range res {
        nums1[i] = num
    }
}

// 解法二：从后往前遍历，只使用 nums1 作为存放最终结果的数组
// 时间复杂度: O(m+n)
// 空间复杂度：O(1)
func merge_2(nums1 []int, m int, nums2 []int, n int)  {
    var i = m - 1
    var j = n - 1           // nums2 数组的最后一个元素的下标
    var cnt = m + n -1      // nums1 数组的最后一个元素的下标

    for i >= 0 || j >= 0 {
        // 注意边界条件的判定
        // j < 0 || (i >= 0 && nums1[i] > nums2[j])    表示 "num2 已遍历完" or "num1 和 num2 未遍历完，并且 num1[i] > num2[j]"
        if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {   
            nums1[cnt] = nums1[i]
            i--
        } else {
            nums1[cnt] = nums2[j]
            j--
        }
        cnt--
    }
}