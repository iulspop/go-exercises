package stack

type ListNode struct {
	Val int
  Next *ListNode
}

type linkedListStack struct {
	head *ListNode
	length int
}

func (s *linkedListStack) Push(val int)  {
	s.head = &ListNode{ val, s.head }
	s.length++
}

func (s *linkedListStack) Pop() int {
	val := s.head.Val
	s.head = s.head.Next
	s.length--
	return val
}

func (s *linkedListStack) Peak() int {
	return s.head.Val
}

func (s *linkedListStack) Size() int {
	return s.length
}

func NewLinkedListStack() Stack {
	return &linkedListStack{}
}
