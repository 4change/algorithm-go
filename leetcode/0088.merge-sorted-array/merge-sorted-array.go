package problem0088

func merge_1(nums1 []int, m int, nums2 []int, n int)  {
    var res []int
    var i = 0
    var j = 0

    for (i < m || j < n) {
        if (j >= n || (i < m && nums1[i] <= nums2[j])) {
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

func merge_2(nums1 []int, m int, nums2 []int, n int)  {
    var i = m - 1
    var j = n - 1
    var cnt = m + n -1

    for i >= 0 || j >= 0 {
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