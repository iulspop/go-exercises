package remove_nth_from_end

import (
	"go-exercises/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	result := RemoveNthFromEnd(createList([]int{1, 2, 3, 4, 5}), 2)
	expected := createList([]int{1, 2, 3, 5})
	assert.Equals(t, "RemoveNthFromEnd()", "", "", result, expected)

	result = RemoveNthFromEnd(createList([]int{1}), 1)
	expected = nil
	assert.Equals(t, "RemoveNthFromEnd()", "", "", result, expected)

	result = RemoveNthFromEnd(createList([]int{1, 2}), 2)
	expected = &ListNode{2, nil}
	assert.Equals(t, "RemoveNthFromEnd()", "", "", result, expected)

	result = RemoveNthFromEnd(createList([]int{1, 2}), 1)
	expected = &ListNode{1, nil}
	assert.Equals(t, "RemoveNthFromEnd()", "", "", result, expected)
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
