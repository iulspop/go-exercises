package insertion_sort


func InsertionSort(numbers []int) []int {
	newNumbers := make([]int, 0)
	newNumbers = append(newNumbers, numbers...)
	numbers = newNumbers

	for i := 1; i < len(numbers); i++ {
		number := numbers[i]
		for z := i - 1; z >= 0; z-- {
			comparedNumber := numbers[z]
			if comparedNumber > number {
				numbers[z + 1] = comparedNumber
				if z == 0 { numbers[z] = number }
			} else {
				numbers[z + 1] = number; break
			}
		}
	}

	return numbers
}

/*

iterate from 1 to len - 1
	save the num
	iterate from i to 0
		if current > num => insert next index over
		else insert savedNum current index and break

*/