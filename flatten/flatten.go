package flatten

func Flatten(nested []interface{}) []interface{} {
	var numbers []interface{}

	for i := 0; i < len(nested); i++ {
		number, ok := nested[i].(int)

		if (ok) {
			numbers = append(numbers, number)
		} else {
			slice := nested[i].([]interface{})
			numbers = append(numbers, Flatten(slice)...)
		}
	}

	return numbers
}
