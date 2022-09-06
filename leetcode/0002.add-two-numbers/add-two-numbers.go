package problem0002

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head *ListNode      // 新链表的头指针
    var tail *ListNode      // tail 指针用于遍历并链接新链表
    carry := 0              // 进位标志     

    for l1 != nil || l2 != nil {
        n1, n2 := 0, 0

        // 遍历链表
        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }

        sum := n1 + n2 + carry

        // 新链表中 当前位置的值 和 进位标志 求解
        sum, carry = sum%10, sum/10

        // 新链表中节点的链接
        if head == nil {    // 新链表头结点为空则进行初始化
            tail = &ListNode{Val: sum}
            head = tail 
        } else {        // 新链表中节点的链接(先链接新节点，再向后移动指针)
            tail.Next = &ListNode{Val: sum}
            tail = tail.Next
        }
    }

    // 新链表中最后一个进位的链接
    if carry > 0 {
        tail.Next = &ListNode{Val: carry}
    }

    return head
}