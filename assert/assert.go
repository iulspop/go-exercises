package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func Equals(t *testing.T, module string, result interface{}, expected interface{}) {
	message := CreateTestMessage(module, result, expected)

	if reflect.DeepEqual(result, expected) {
		fmt.Print(message)
	} else {
		t.Errorf(message)
	}
}

func CreateTestMessage(module string, result interface{}, expected interface{}) string {
	return fmt.Sprintf(
		`
%v
given:    a set of jobs with start and end
should:   return max number of non-overlapping jobs
result:   %v
expected: %v
	`, module, result, expected)
}
