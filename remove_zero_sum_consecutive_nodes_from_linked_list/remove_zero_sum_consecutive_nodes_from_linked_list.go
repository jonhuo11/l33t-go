package removezerosumconsecutivenodesfromlinkedlist

import "fmt"

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
 
 func removeZeroSumSublists(head *ListNode) *ListNode {
    
    c := head

    a := []*ListNode{&ListNode{Val:0, Next: head}, head}
    sums := []int{0, head.Val}
    i := 1
    for c.Next != nil {
        c = c.Next
        i++
        sums = append(sums, sums[i-1] + c.Val)
        a = append(a, c)
    }
    a = append(a, nil)
    fmt.Println(sums)

    seen := map[int]int{} // sum, and wheere in sums it was seen
    i = 0
    for i < len(sums) {
        if index, saw := seen[sums[i]]; saw {
            a[index].Next = a[i + 1]
            for k := i - 1; sums[k] != sums[i] ; k-- {
                delete(seen, sums[k])
            }
            //el = true
        } else {
            seen[sums[i]] = i
        }
        //fmt.Println(seen)
        i++
    }
    
    if a[0].Next != nil && a[0].Next.Next == nil && a[0].Next.Val == 0 {
        return nil
    }
    return a[0].Next
}