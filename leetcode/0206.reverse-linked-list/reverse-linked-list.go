package problem0206

func reverseList(head *ListNode) *ListNode {
	var help *ListNode			// 新链表的头结点

	for head != nil {
		next := head.Next		// 确定右边子链表的下一节点位置

		head.Next = help		// 断开

		help = head             // 确定左边子链表当前节点的位置
		head = next				// 确定右边子链表当前节点的位置
	}
	
	return help
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}