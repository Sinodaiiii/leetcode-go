package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum260614(head *ListNode) int {
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2, pre2 := slow.Next, fast.Next
	for head2 != nil {
		head2.Next, head2, pre2 = pre2, head2.Next, head2
	}
	link1, link2 := head, fast
	ans := 0
	for link2 != nil {
		ans = max(ans, link1.Val+link2.Val)
		link1, link2 = link1.Next, link2.Next
	}
	return ans
}
