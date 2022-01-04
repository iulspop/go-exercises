package flatten

import (
	"go-exercises/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	result := Flatten([]interface{} { 1, 2, 3, []interface{} {4, 5, 6}, 7, []interface{} {8, []interface{} {9, 10, 11, []interface{} {12, 13, 14} } }, []interface{} {15, 16, 17, 18, 19, []interface{} {20, 21, 22, []interface{} {23, 24, 25, []interface{} {26, 27, 29} }, 30, 31 }, 32 }, 33 })
	expected := []interface{} { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 29, 30, 31, 32, 33 }
	assert.Equals(t, "Flatten()", "Given", "Should ", result, expected)
}
