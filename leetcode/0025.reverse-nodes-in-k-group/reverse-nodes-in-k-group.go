package main

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    protect := &ListNode{ Val: 0, Next: head }
    pre := protect

    for (head != nil) {
// 1. 分组，往后遍历 k-1 步为一组，head & end
        end := getEnd(head, k)
        if (end == nil) {
            break
        }
        nextGrpHead := end.Next

// 2. 组内做反转（206 题解）
        reverseList(head, nextGrpHead)

// 3. 维护组间的边， 前一个组、后一个组的边
        pre.Next = end
        head.Next = nextGrpHead
        pre = head
        head = nextGrpHead
    }

    return protect.Next
}

func reverseList(head *ListNode, nextGrpHead *ListNode) {
	help := head
    head = head.Next

	for head != nextGrpHead {
		next := head.Next		// 确定右边子链表的下一节点位置

		head.Next = help		// 断开

		help = head             // 确定左边子链表当前节点的位置
		head = next				// 确定右边子链表当前节点的位置
	}
}

// return end, k - 1 步
// return nil, 分组不够 K 个
func getEnd(head *ListNode, k int) *ListNode {
    for head != nil {
        k--
        if k == 0 {
            return head
        }
        head = head.Next
    }
    return nil
}