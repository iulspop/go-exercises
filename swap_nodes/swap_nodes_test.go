package swap_nodes

import (
	"go-exercises/assert"
	"testing"
)

func TestSwapNodes(t *testing.T) {
	result := listToSlice(SwapNodes(createList([]int {1,2,3,4,5}), 2))
	expected := []int {1,4,3,2,5}
	assert.Equals(t, "SwapNodes()", "", "", result, expected)
}

func createList(values []int) *ListNode {
	head := &ListNode{values[0], nil}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{values[i], nil}
		current = current.Next
	}

	return head
}

func listToSlice(head *ListNode) []int {
	slice := []int {}
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}