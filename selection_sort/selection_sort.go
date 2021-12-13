package selection_sort


func SelectionSort(numbers []int) []int {
	for i := 0; i < len(numbers) - 1; i++ {
		minimumNum := -1
		minimumNumIndex := -1

		for z := i; z < len(numbers); z++ {
			num := numbers[z]
			if minimumNumIndex == -1 || num < minimumNum {
				minimumNum = num
				minimumNumIndex = z
			}
		}

		numbers[minimumNumIndex] = numbers[i]
		numbers[i] = minimumNum
	}
	return numbers
}

/*

iterate from 0 to len - 2
	iterate from n to len - 1
		save index smallest num
	swap smallest num with index n

*/