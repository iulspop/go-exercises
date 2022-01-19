package stack

type Stack interface {
	Push(val int)
	Pop() int
	Peak() int
	Size() int
}
