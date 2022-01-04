package reverse_string

import (
	"go-exercises/slice_stack"
)

func ReverseString(chars string) string {
	stack := slice_stack.NewSliceStack()

	for _, char := range chars {
		stack.Push(string(char))
	}

	reversedChars := ""
	for i := 0; i < len(chars); i++ {
		reversedChars += stack.Pop().(string)
	}

	return reversedChars
}