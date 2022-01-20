package is_symmetric

import (
	"go-exercises/assert"
	"testing"
	"go-exercises/binary_tree/utils"
)

func TestIsSymmetric(t *testing.T) {
	result := IsSymmetric(utils.CreateTree([]int {1,2,2,3,4,4,3}))
	expected := true
	assert.Equals(t, "IsSymmetric()", "", "", result, expected)
}
