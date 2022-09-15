package problem0021

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{ Val: 0 }
    protected := head

    for (list1 != nil || list2 != nil) {
        if (list2 == nil || (list1 != nil && list1.Val <= list2.Val)) {
            head.Next = list1
            list1 = list1.Next
        } else {
            head.Next = list2
            list2 = list2.Next
        }

        head = head.Next
    }

    return protected.Next
}