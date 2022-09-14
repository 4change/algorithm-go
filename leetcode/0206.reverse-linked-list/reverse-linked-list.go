package problem0206

func reverseList(head *ListNode) *ListNode {
	var help *ListNode

	for head != nil {
		next := head.Next
		head.Next = help
		help = head
		head = next
	}
	
	return help
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}