package slice_stack

type Stack interface {
	Push(val interface{})
	Pop() interface{}
	Read() interface{}
}

type stack struct {
	array []interface{}
}

func (s *stack) Push(val interface{})  {
	s.array = append(s.array, val)
}

func (s *stack) Pop() interface{} {
	val := s.array[len(s.array) - 1]
	s.array = s.array[:len(s.array) - 1]
	return val
}

func (s *stack) Read() interface{} {
	return s.array[len(s.array) - 1]
}

func NewSliceStack() Stack {
	return &stack{}
}