package stack

import (
	"go-exercises/assert"
	"testing"
)

func TestSliceStack(t *testing.T) {
	stack := NewSliceStack()

	stack.Push(1)
	stack.Push(2)

	result := stack.Peak()
	expected := 2
	assert.Equals(t, "Stack.Peak()", "Given two elements pushed", "Should read last element in stack", result, expected)

	stack.Push(3)
	result = stack.Pop()
	expected = 3
	assert.Equals(t, "Stack.Pop()", "Given three elements pushed & one popped", "Should return element popped", result, expected)

	result = stack.Peak()
	expected = 2
	assert.Equals(t, "Stack.Peak()", "Given three elements pushed & one popped", "Should read second element", result, expected)

	stack.Push(1)
	stack.Push(2)

	result = stack.Size()
	expected = 4
	assert.Equals(t, "Stack.Peak()", "Given three elements pushed & one popped & two pushed", "Should return size 4", result, expected)
}

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack()

	stack.Push(1)
	stack.Push(2)

	result := stack.Peak()
	expected := 2
	assert.Equals(t, "Stack.Peak()", "Given two elements pushed", "Should read last element in stack", result, expected)

	stack.Push(3)
	result = stack.Pop()
	expected = 3
	assert.Equals(t, "Stack.Pop()", "Given three elements pushed & one popped", "Should return element popped", result, expected)

	result = stack.Peak()
	expected = 2
	assert.Equals(t, "Stack.Peak()", "Given three elements pushed & one popped", "Should read second element", result, expected)

	stack.Push(1)
	stack.Push(2)

	result = stack.Size()
	expected = 4
	assert.Equals(t, "Stack.Peak()", "Given three elements pushed & one popped & two pushed", "Should return size 4", result, expected)
}
