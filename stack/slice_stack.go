package stack

type sliceStack struct {
	array []int
}

func (s *sliceStack) Push(val int)  {
	s.array = append(s.array, val)
}

func (s *sliceStack) Pop() int {
	val := s.array[len(s.array) - 1]
	s.array = s.array[:len(s.array) - 1]
	return val
}

func (s *sliceStack) Peak() int {
	return s.array[len(s.array) - 1]
}

func (s *sliceStack) Size() int {
	return len(s.array)
}

func NewSliceStack() Stack {
	return &sliceStack{}
}