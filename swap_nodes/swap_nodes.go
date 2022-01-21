package swap_nodes

type ListNode struct {
		Val int
		Next *ListNode
}

func SwapNodes(head *ListNode, k int) *ListNode {
	leader := &ListNode{Next: head}
	follower := leader

	var begin *ListNode
	n := 0
	for leader != nil {
		if n == k  { begin = leader }
		if n >= k { follower = follower.Next }
		leader = leader.Next
		n++
	}

	end := follower

  begin.Val, end.Val = end.Val, begin.Val
	return head
}

/*
algo:
- return head if head.Next == nil
- leader := dummy.Next = head
- follower := dummy.Next = head
- var begin
- n := 0
- for leader != nil
  - if n == (k - 1) { begin = leader }
  - if n > (k + 1) { increment follower }
  - increment leader
  - n++
- end := follower
- begin.Next.Next, end.Next.Next = end.Next.Next, begin.Next.Next
- begin.Next, end.Next = end.Next, begin.Next
- return head

*/