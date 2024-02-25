package maximumtwinsumofalinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
	Val int
	Next *ListNode
}
 func pairSum(head *ListNode) int {
    end := head
    mid := head
    o := 1
    for end.Next != nil {
        end = end.Next
        if o % 2 == 0 {
            mid = mid.Next
        }
        o++
    }
    if o == 2 {
        return head.Val + head.Next.Val
    }

    l := mid
    r := mid.Next
    l.Next = nil
    for r != nil {
        tmp := r.Next
        r.Next = l
        l = r
        r = tmp
    }

    mx := 0

    l = head
    for end != mid {
        if l.Val + end.Val > mx {
            mx = l.Val + end.Val
        }
        l = l.Next
        end = end.Next
    }
    return mx
}