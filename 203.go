/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var i *ListNode
    j := head
    for (j!=nil) {
        if j.Val == val{
            if j==head {
                head = j.Next
            }else{
                for(j!=nil && j.Val == val){
                    j=j.Next
                }
                i.Next = j
            }
        }
        i=j
        if j!=nil {
            j=j.Next
        }
    }
    return head    
}