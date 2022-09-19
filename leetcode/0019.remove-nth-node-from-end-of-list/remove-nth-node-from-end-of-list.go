package main

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    protected := &ListNode{ Val:0, Next: head}
    fast, slow := head, protected

    for i := 0; i < n; i++ {
        fast = fast.Next
    }

    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next

    return protected.Next
}