package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func Equals(t *testing.T, module string, given string, should string, result interface{}, expected interface{}) {
	message := createTestMessage(module, given, should, result, expected)

	if reflect.DeepEqual(result, expected) {
		fmt.Print(message)
	} else {
		t.Errorf(message)
	}
}

func createTestMessage(module string, given string, should string, result interface{}, expected interface{}) string {
	return fmt.Sprintf(
		`
%v
given:    %v
should:   %v
result:   %v
expected: %v
	`, module, given, should, result, expected)
}
