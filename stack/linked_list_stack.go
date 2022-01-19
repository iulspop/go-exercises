package stack

type listNode struct {
	Val int
  Next *listNode
}

type linkedListStack struct {
	head *listNode
	length int
}

func (s *linkedListStack) Push(val int)  {
	s.head = &listNode{ val, s.head }
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
