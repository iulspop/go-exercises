package remove_nth_from_end

type ListNode struct {
	Val int
	 Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
  if head.Next == nil { return nil }

  leader, follower := head, head
  count := 0

  for leader != nil {
    leader = leader.Next
    if count >= n + 1 { follower = follower.Next }
    count++
  }

  if count == n {
    head = follower.Next
  } else {
    follower.Next = follower.Next.Next
  }

  return head
}

/*
2
[1, 2], 2
       l
 f

- if head.Next == nil { return nil }
- init leader, follower at head
- init count = 0
- for leader != nil
  - leader = leader.next
  - if count >= nth + 1
    - follower = follower.next
  - count++
- follower.Next = follower.Next.Next
- return head


input: head, nth
output: list with nth removed

explicit req:
- nth node from the **end** of the list!!!!
- 1 to 30 nodes
- n within the size

implicit req:
- never null reference input

test cases:
[1, 2, 3, 4, 5], 2
[1, 2, 3, 5]


[1], 1
[]

[1, 2], 1
[1]

mental model:

go once and figure out the length
go twice and do deletion at the right spot

hash
  count forward as keys: address as values  
iterate to end of list and count length, then calculate which addresses to reference from hash
[1, 2, 3, 4, 5]


follower,
leader pointers

leader moves ahead every loop
followers only moves ahead after nth count

5

[1, 2, 3, 4, 5], 2
                l         
       f

0
[1], 1
   l
 f

2
[1, 2], 1
       l
 f



algo:
- if head.Next == nil { return nil }
- init leader, follower at head
- init count = 0
- for leader != nil
  - leader = leader.next
  - if count >= nth + 1
    - follower = follower.next
  - count++
- follower.Next = follower.Next.Next
- return head



HASH APPROACH:
algo:
- init map[int]*ListNode
- current = head
- count = 1
- for current != nil 
  - map[count] = current
  - current = current.next
  - count++
- map[count] = nil
- beforeindex count - nth - 1
- afterIndex count - nth + 1


*/