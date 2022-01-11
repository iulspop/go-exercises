package linked_list

import (
	"go-exercises/assert"
	"testing"
)

func TestSliceStack(t *testing.T) {
	stack := NewSliceStack()

	stack.Push(1)
	stack.Push(2)

	result := stack.Read()
	expected := 2
	assert.Equals(t, "Stack.Read()", "Given two elements pushed", "Should read last element in stack", result, expected)

	stack.Push(3)
	result = stack.Pop()
	expected = 3
	assert.Equals(t, "Stack.Pop()", "Given three elements pushed & one popped", "Should return element popped", result, expected)

	result = stack.Read()
	expected = 2
	assert.Equals(t, "Stack.Read()", "Given three elements pushed & one popped", "Should read second element", result, expected)
}
